package sendazureiot

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("connectionString", "HostName=HomeAutoHub.azure-devices.net;DeviceId=raspi;SharedAccessKey=IHx8ac6Bad4vHbv4I0HiJkhgeCNZhuzQdnllCAMSR+o=")
	tc.SetInput("message", "Testing")
	tc.SetInpu("deviceID", "raspi")

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	status := tc.GetOutput("status")
	assert.Equal(t, result, "Successful Response")
	assert.Equal(t, result, "200 Ok")

}
