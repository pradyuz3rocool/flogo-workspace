package tester

import (
	"os"
	"strings"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/definition"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/instance"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/model"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/support"
	"github.com/TIBCOSoftware/flogo-contrib/model/simple"
	"github.com/TIBCOSoftware/flogo-lib/util"
)

const (
	ENV_ENABLED         = "TESTER_ENABLED"
	ENV_SETTING_PORT    = "TESTER_PORT"
	ENV_SETTING_SR_HOST = "TESTER_SR_SERVER"
)

//ExtensionProvider is the extension provider for the flow action
type TesterProvider struct {
	flowProvider  definition.Provider
	flowModel     *model.FlowModel
	stateRecorder instance.StateRecorder
	flowTester    *RestEngineTester
}

func NewExtensionProvider() *TesterProvider {
	return &TesterProvider{}
}

func (fp *TesterProvider) GetFlowProvider() definition.Provider {
	if fp.flowProvider == nil {
		fp.flowProvider = &support.BasicRemoteFlowProvider{}
	}

	return fp.flowProvider
}

func (fp *TesterProvider) GetDefaultFlowModel() *model.FlowModel {
	if fp.flowModel == nil {
		fp.flowModel = simple.New()
	}

	return fp.flowModel
}

func (fp *TesterProvider) GetStateRecorder() instance.StateRecorder {

	if fp.stateRecorder == nil {
		config := &util.ServiceConfig{Enabled: true}

		server := os.Getenv(ENV_SETTING_SR_HOST)

		if server != "" {
			parts := strings.Split(server, ":")

			host := parts[0]
			port := "9090"

			if len(parts) > 1 {
				port = parts[1]
			}

			settings := map[string]string{
				"host": host,
				"port": port,
			}
			config.Settings = settings
		} else {
			config.Enabled = false
		}

		fp.stateRecorder = instance.NewRemoteStateRecorder(config)
	}

	return fp.stateRecorder
}

func (fp *TesterProvider) GetMapperFactory() definition.MapperFactory {
	return nil
}

func (fp *TesterProvider) GetLinkExprManagerFactory() definition.LinkExprManagerFactory {
	return nil
}

func (fp *TesterProvider) GetFlowTester() *RestEngineTester {

	config := &util.ServiceConfig{Enabled: true}

	settings := map[string]string{
		"port": os.Getenv(ENV_SETTING_PORT),
	}
	config.Settings = settings
	return NewRestEngineTester(config)
}
