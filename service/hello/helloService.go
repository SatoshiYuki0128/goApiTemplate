package hello

import "goApiFramework/model"

func Hello(flowData *model.FlowData, controllerCode string) {
	flowData.Response = "Hello"
}
