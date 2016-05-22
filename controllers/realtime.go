package controllers

import (
	"iControl/models"
)

type RealTimeController struct {
	BaseController
}

//渲染页面
func (this *RealTimeController) List() {
	this.display()
}

//供实时查询页面ajax异步获取数据
func (this *RealTimeController) GetData() {
	result, _ := models.TagGetList()

	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["Id"] = v.Id
		row["Type"] = v.Type
		row["Num"] = v.Num
		row["Value"] = v.Value

		list[k] = row
	}
	this.jsonResult(list)
}
