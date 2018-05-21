package sendazureiot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-azureiot")

const (
	ivconnectionString = "connectionString"
	ivMessage          = "message"
	ivDeviceID         = "Device ID"
	ivaction           = "Action"

	ovResult = "result"
	ovStatus = "status"

	maxIdleConnections int    = 100
	requestTimeout     int    = 10
	tokenValidSecs     int    = 3600
	apiVersion         string = "2016-11-14"
)

type sharedAccessKey = string
type sharedAccessKeyName = string
type hostName = string
type deviceID = string

// IotHubHTTPClient is a simple client to connect to Azure IoT Hub
type IotHubHTTPClient struct {
	sharedAccessKeyName sharedAccessKeyName
	sharedAccessKey     sharedAccessKey
	hostName            hostName
	deviceID            deviceID
	client              *http.Client
}

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval

	connectionString := context.GetInput(ivconnectionString).(string)
	message := context.GetInput(ivMessage).(string)
	//action := context.GetInput(ivaction).(string)
	deviceID := context.GetInput(ivDeviceID).(string)

	log.Debug("The connection string to device is [%s]", connectionString)
	log.Debug("The Method type selected is [%s]", message)
	log.Debug("The Devic ID is [%s]", deviceID)

	client, err := NewIotHubHTTPClientFromConnectionString(connectionString)
	if err != nil {
		log.Error("Error creating http client from connection string", err)
	}
	resp, status := client.SendMessage(message)
	context.SetOutput(ovResult, resp)
	context.SetOutput(ovStatus, status)

	return true, nil
}

func parseConnectionString(connString string) (hostName, sharedAccessKey, sharedAccessKeyName, deviceID, error) {
	url, err := url.ParseQuery(connString)
	if err != nil {
		return "", "", "", "", err
	}

	h := tryGetKeyByName(url, "HostName")
	kn := tryGetKeyByName(url, "SharedAccessKeyName")
	k := tryGetKeyByName(url, "SharedAccessKey")
	d := tryGetKeyByName(url, "DeviceId")

	return hostName(h), sharedAccessKey(k), sharedAccessKeyName(kn), deviceID(d), nil
}

func tryGetKeyByName(v url.Values, key string) string {
	if len(v[key]) == 0 {
		return ""
	}

	return strings.Replace(v[key][0], " ", "+", -1)
}

// NewIotHubHTTPClient is a constructor of IutHubClient
func NewIotHubHTTPClient(hostName string, sharedAccessKeyName string, sharedAccessKey string, deviceID string) *IotHubHTTPClient {
	return &IotHubHTTPClient{
		sharedAccessKeyName: sharedAccessKeyName,
		sharedAccessKey:     sharedAccessKey,
		hostName:            hostName,
		deviceID:            deviceID,
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: maxIdleConnections,
			},
			Timeout: time.Duration(requestTimeout) * time.Second,
		},
	}
}

// NewIotHubHTTPClientFromConnectionString creates new client from connection string
func NewIotHubHTTPClientFromConnectionString(connectionString string) (*IotHubHTTPClient, error) {
	h, k, kn, d, err := parseConnectionString(connectionString)
	if err != nil {
		return nil, err
	}

	return NewIotHubHTTPClient(h, kn, k, d), nil
}

// IsDevice tell either device id was specified when client created.
// If device id was specified in connection string this will enabled device scoped requests.
func (c *IotHubHTTPClient) IsDevice() bool {
	return c.deviceID != ""
}

// Device API

// SendMessage from a logged in device
func (c *IotHubHTTPClient) SendMessage(message string) (string, string) {
	url := fmt.Sprintf("%s/devices/%s/messages/events?api-version=%s", c.hostName, c.deviceID, apiVersion)
	return c.performRequest("POST", url, message)
}

func (c *IotHubHTTPClient) buildSasToken(uri string) string {
	timestamp := time.Now().Unix() + int64(3600)
	encodedURI := template.URLQueryEscaper(uri)

	toSign := encodedURI + "\n" + strconv.FormatInt(timestamp, 10)

	binKey, _ := base64.StdEncoding.DecodeString(c.sharedAccessKey)
	mac := hmac.New(sha256.New, []byte(binKey))
	mac.Write([]byte(toSign))

	encodedSignature := template.URLQueryEscaper(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	if c.sharedAccessKeyName != "" {
		return fmt.Sprintf("SharedAccessSignature sig=%s&se=%d&skn=%s&sr=%s", encodedSignature, timestamp, c.sharedAccessKeyName, encodedURI)
	}

	return fmt.Sprintf("SharedAccessSignature sig=%s&se=%d&sr=%s", encodedSignature, timestamp, encodedURI)
}

func (c *IotHubHTTPClient) performRequest(method string, uri string, data string) (string, string) {
	token := c.buildSasToken(uri)
	//log.("%s https://%s\n", method, uri)
	log.Debug(token)
	req, _ := http.NewRequest(method, "https://"+uri, bytes.NewBufferString(data))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "golang-iot-client")
	req.Header.Set("Authorization", token)

	//log.Println("Authorization:", token)

	if method == "DELETE" {
		req.Header.Set("If-Match", "*")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		log.Error(err)
	}

	// read the entire reply to ensure connection re-use
	text, _ := ioutil.ReadAll(resp.Body)

	io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()

	return string(text), resp.Status
}
