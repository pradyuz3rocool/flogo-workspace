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

	ovURL   = "url"
	ovToken = "SAS Token"
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
	log.Debug("sharedAccessKeyName", sharedAccessKeyName)
	log.Debug("sharedAccessKey ", sharedAccessKey)
	url := fmt.Sprintf("https://%s/devices/%s/messages/deviceBound?api-version=2016-11-14", hostName, deviceID)
	SaS := createSharedAccessToken(url, sharedAccessKeyName)

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

func createSharedAccessToken(uri string, saKey string) string {

	if len(uri) == 0 || len(saKey) == 0 {
		return "Missing required parameter"
	}

	encoded := template.URLQueryEscaper(uri)
	now := time.Now().Unix()
	week := 60 * 60 * 24 * 7
	ts := now + int64(week)
	signature := encoded + "\n" + strconv.Itoa(int(ts))
	key := []byte(saKey)
	hmac := hmac.New(sha256.New, key)
	hmac.Write([]byte(signature))
	hmacString := template.URLQueryEscaper(base64.StdEncoding.EncodeToString(hmac.Sum(nil)))
	result := "SharedAccessSignature sr=" + encoded + "&sig=" + hmacString + "&se=" + strconv.Itoa(int(ts))
	return result
}
