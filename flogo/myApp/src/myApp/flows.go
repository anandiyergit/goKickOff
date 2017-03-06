package main

import (
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/engine"
	"github.com/TIBCOSoftware/flogo-lib/flow/flowinst"
	"github.com/TIBCOSoftware/flogo-lib/flow/service"
	"github.com/TIBCOSoftware/flogo-lib/flow/service/flowprovider"
	"github.com/TIBCOSoftware/flogo-lib/flow/service/staterecorder"
	"github.com/TIBCOSoftware/flogo-lib/flow/service/tester"
	"github.com/TIBCOSoftware/flogo-lib/flow/support"
)

var embeddedJSONFlows map[string]string

func init() {

	embeddedJSONFlows = make(map[string]string)

	embeddedJSONFlows["embedded://myflow"] = "H4sIAAAJbogC/7yTMW+DMBCFd37FyXM6pCNbl0qVOlVsVQaDL8iK8aX2QRVF+e8V2ClGWBWqqrLx7t3dh3m+FgAAQjI7XfeMXpTwftgFtSOFRpQgWNcNPXjdnQ2KWLSyw7F2NPR51xwRV9KfRAlhcBjesB40X6rLeeqI7qmmlShhnwhG21MKsdg1zk+7Wfpg/pbG57p4yyKEDzLUJtNm8+IwVvX8ihVrh97LFjMLZvxI49lp2/7kHKTpJ+sbNqgHVODQMzj86NGzyDbedr9FH3/piz3SFvaayKC0m+DZ9fjXrFKpip7nDP4H7Eo9ZFI0JfsxU7iTv1ILTzGXuRhG6H2R353ej4U1nuUs3oovAAAA//8BAAD//7jgquTmAwAA"

}

// EnableFlowServices enables flow services and action for engine
func EnableFlowServices(engine *engine.Engine, engineConfig *engine.Config) {

	log.Debug("Flow Services and Actions enabled")

	embeddedFlowMgr := support.NewEmbeddedFlowManager(true, embeddedJSONFlows)

	fpConfig := engineConfig.Services[service.ServiceFlowProvider]
	flowProvider := flowprovider.NewRemoteFlowProvider(fpConfig, embeddedFlowMgr)
	engine.RegisterService(flowProvider)

	srConfig := engineConfig.Services[service.ServiceStateRecorder]
	stateRecorder := staterecorder.NewRemoteStateRecorder(srConfig)
	engine.RegisterService(stateRecorder)

	etConfig := engineConfig.Services[service.ServiceEngineTester]
	engineTester := tester.NewRestEngineTester(etConfig)
	engine.RegisterService(engineTester)

	options := &flowinst.ActionOptions{Record: stateRecorder.Enabled()}

	flowAction := flowinst.NewFlowAction(flowProvider, stateRecorder, options)
	action.Register(flowinst.ActionType, flowAction)
}
