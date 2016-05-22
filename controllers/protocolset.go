package controllers

import (
	"encoding/json"
	"iControl/models"
	"log"
	"strconv"
	"strings"
)

type ProtocolSetController struct {
	BaseController
}

//渲染页面
func (this *ProtocolSetController) List() {
	this.display()
}

//供实时查询页面ajax异步获取数据
func (this *ProtocolSetController) GetData() {
	result, _ := models.ProtocolGetList()

	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["Id"] = v.Id
		row["Com"] = v.Com
		row["BaudRate"] = v.BaudRate
		row["DataBit"] = v.DataBit
		row["ParityBit"] = v.ParityBit
		row["StopBit"] = v.StopBit

		list[k] = row
	}
	this.jsonResult(list)
}

func (this *ProtocolSetController) AddData() {
	var response *models.Protocol

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &response)
	if err != nil {
		log.Println(err)
		this.Data["json"] = "{\"data\":0}" //1表示修改成功，0表示失败
		this.ServeJSON()
		return
	}
	_, err = models.ProtocolAdd(response)
	if err != nil {
		log.Println(err)
		this.Data["json"] = "{\"data\":0}" //1表示修改成功，0表示失败
		this.ServeJSON()
		return
	}
	this.Data["json"] = "{\"data\":1}" //1表示修改成功，0表示失败
	this.ServeJSON()
}

func (this *ProtocolSetController) UpdateData() {
	var protocol models.Protocol
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &protocol)
	if err != nil {
		log.Println(err)
		this.Data["json"] = "{\"data\":0}" //1表示修改成功，0表示失败
		this.ServeJSON()
		return
	}

	obj, _ := models.ProtocolGetById(protocol.Id)
	if this.isPost() {
		obj.Com = protocol.Com
		obj.BaudRate = protocol.BaudRate
		obj.DataBit = protocol.DataBit
		obj.ParityBit = protocol.ParityBit
		obj.StopBit = protocol.StopBit
		err = obj.Update()
		if err != nil {
			log.Println(err)
			this.Data["json"] = "{\"data\":0}" //1表示修改成功，0表示失败
			this.ServeJSON()
			return
		}
	}

	this.Data["json"] = "{\"data\":1}"
	this.ServeJSON()

}

func (this *ProtocolSetController) DelData() {
	var response myDelStruct

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &response)
	if err != nil {
		log.Println(err)
		this.Data["json"] = "{\"data\":0}" //1表示修改成功，0表示失败
		this.ServeJSON()
		return
	}
	sli := strings.Split(response.Id, ",")
	for _, value := range sli {
		i, _ := strconv.Atoi(value)
		err = models.ProtocolDelete(i)
		if err != nil {
			log.Println(err)
			this.Data["json"] = "{\"data\":0}" //1表示修改成功，0表示失败
			this.ServeJSON()
			return
		}
	}

	this.Data["json"] = "{\"data\":1}" //1表示修改成功，0表示失败
	this.ServeJSON()
}
