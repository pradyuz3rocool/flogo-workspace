package azureiot

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-azureiot")

const (
	ivdeviceID        = "deviceID"
	ivhostName        = "hostName"
	ivsharedAccessKey = "sharedAccessKey"

	ovresult = "result"
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

	sharedAccessKey := context.GetInput(ivsharedAccessKey).(string)
	hostName := context.GetInput(ivhostName).(string)
	deviceID := context.GetInput(ivdeviceID).(string)

	log.Debug("The hostname is [%s]", hostName)
	log.Debug("The device is [%s]", deviceID)
	log.Debug("The shared access key is [%s]", sharedAccessKey)

	context.SetOutput("result", "Trying to connect to device "+deviceID+" using hostname: "+hostName+"and sharedAccesskey as "+sharedAccessKey)

	return true, nil
}
