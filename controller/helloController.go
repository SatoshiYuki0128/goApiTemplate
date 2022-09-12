package controller

import (
	"goApiFramework/model"
	commonService "goApiFramework/service/common"
	helloService "goApiFramework/service/hello"
	"net/http"
)

// Hello controllerCode:H1
func Hello(w http.ResponseWriter, r *http.Request) {
	flowData := model.FlowData{}
	if commonService.GetAndValidateRequest[model.HelloReq](r, &flowData, "H1") {
		helloService.Hello(&flowData, "H1")
	}
	commonService.ServeResponse(w, &flowData)
}
