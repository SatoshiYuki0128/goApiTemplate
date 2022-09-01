package controller

import (
	"goApiFramework/model"
	"goApiFramework/service/common"
	"letinstall-api/models"
	applistservice "letinstall-api/services/appList"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	flowData := model.FlowData{}
	if service.GetAndValidateRequest[models.PostAppReq](r, &flowData, "AL1") {
		applistservice.GetAppList(&flowData, "AL1")
	}
}
