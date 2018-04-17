package azureiot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
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

	ovResult = "result"
	ovStatus = "status"
	ovURL    = "url"
	ovToken  = "token"
)

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

	log.Debug("The connection string to device is [%s]", connectionString)
	hostName, sharedAccessKey, sharedAccessKeyName, deviceID, err := parseConnectionString(connectionString)
	url := fmt.Sprintf("https://%s/devices/%s/messages/deviceBound?api-version=2016-11-14", hostName, deviceID)
	SaS := createSharedAccessToken(url, sharedAccessKeyName, sharedAccessKey)

	context.SetOutput(ovResult, hostName+sharedAccessKey)
	context.SetOutput(ovStatus, sharedAccessKeyName+deviceID)
	context.SetOutput(ovURL, url)
	context.SetOutput(ovToken, SaS)
	return true, nil
}

func parseConnectionString(connString string) (string, string, string, string, error) {
	url, err := url.ParseQuery(connString)
	if err != nil {
		return "", "", "", "", err
	}

	h := tryGetKeyByName(url, "HostName")
	kn := tryGetKeyByName(url, "SharedAccessKeyName")
	k := tryGetKeyByName(url, "SharedAccessKey")
	d := tryGetKeyByName(url, "DeviceId")

	hostName := h
	sharedAccessKeyName := kn
	sharedAccessKey := k
	deviceID := d
	return hostName, sharedAccessKeyName, sharedAccessKey, deviceID, nil
}

func tryGetKeyByName(v url.Values, key string) string {
	if len(v[key]) == 0 {
		return ""
	}

	return strings.Replace(v[key][0], " ", "+", -1)
}

func createSharedAccessToken(uri string, saName string, saKey string) string {

	timestamp := time.Now().Unix() + int64(3600)
	encodedURI := template.URLQueryEscaper(uri)

	toSign := encodedURI + "\n" + strconv.FormatInt(timestamp, 10)

	binKey, _ := base64.StdEncoding.DecodeString(saKey)
	mac := hmac.New(sha256.New, []byte(binKey))
	mac.Write([]byte(toSign))

	encodedSignature := template.URLQueryEscaper(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	return fmt.Sprintf("SharedAccessSignature sig=%s&se=%d&skn=%s&sr=%s", encodedSignature, timestamp, saName, encodedURI)
}
