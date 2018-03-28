package azureiot

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-azureiot")

const (
	ivconnectionString = "connectionString"

	ovResult = "result"
	ovStatus = "status"
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
	client, err := NewIotHubHTTPClientFromConnectionString(connectionString)
	if err != nil {
		log.Error("Error creating http client from connection string", err)
	}
	resp, status := client.GetDeviceID(client.deviceID)
	//result := fmt.Sprintf("%s is the response, %s is the status", resp, status)
	context.SetOutput(ovResult, resp)
	context.SetOutput(ovStatus, status)
	return true, nil
}
