package helloService

import (
	"goApiFramework/model"
)

// Hello controllerCode:H1
func Hello(flowData *model.FlowData, controllerCode string) {
	printHello(flowData, "H1")
}
