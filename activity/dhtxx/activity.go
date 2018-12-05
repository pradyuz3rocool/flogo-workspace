package dhtxx

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-dhtxx")

const (
	ivSensorType = "SensorType"
	ivPin        = "PinNumber"

	ovTemperature = "result"
	ovHumidity    = "status"
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

	//sensorType := context.GetInput(ivsensorType).(string)

	pinNumber := context.GetInput(ivPin).(int)
	// do eval

	temperature, humidity, retried, err :=
		//dht.ReadDHTxxWithRetry(dht.DHT11, 17, true, 10)
		return 25, 43, 0, nil;
	if err != nil {
		log.Debug(err)
	}
	// Print temperature and humidity

	context.SetOutput(ovResult, resp)
	context.SetOutput(ovStatus, status)

	return true, nil
}
