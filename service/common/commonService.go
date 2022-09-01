package common

import (
	"goApiFramework/model"
	"net/http"
)

// GetAndValidateRequest serviceCode:GAVR
func GetAndValidateRequest[T any](httpRequest *http.Request, flowData *model.FlowData, controllerCode string) (result bool) {
	initFlowData(flowData, controllerCode, "C1")
	data, isOK := getRequestBody(httpRequest, flowData, controllerCode, "C1")
	if isOK {
		checkModel[T](data, httpRequest, flowData, controllerCode, "C1")

		if flowData.Err == nil {
			result = true
		}
	}

	return
}
