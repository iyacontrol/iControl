package controllers

import (
	"encoding/json"
	"iControl/models"
	"log"
	"strconv"
	"strings"
)

type IpSetController struct {
	BaseController
}

//渲染页面
func (this *IpSetController) List() {
	this.display()
}

//供实时查询页面ajax异步获取数据
func (this *IpSetController) GetData() {
	result, _ := models.IpGetList()

	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["Id"] = v.Id
		row["Address"] = v.Address
		row["CenterName"] = v.CenterName
		row["CenterPosition"] = v.CenterPosition

		list[k] = row
	}
	this.jsonResult(list)
}

func (this *IpSetController) AddData() {
	var response *models.Ip

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &response)
	if err != nil {
		log.Println(err)
		this.Data["json"] = "{\"data\":0}" //1表示修改成功，0表示失败
		this.ServeJSON()
		return
	}
	_, err = models.IpAdd(response)
	if err != nil {
		log.Println(err)
		this.Data["json"] = "{\"data\":0}" //1表示修改成功，0表示失败
		this.ServeJSON()
		return
	}
	this.Data["json"] = "{\"data\":1}" //1表示修改成功，0表示失败
	this.ServeJSON()
}

func (this *IpSetController) UpdateData() {
	var ip models.Ip
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ip)
	if err != nil {
		log.Println(err)
		this.Data["json"] = "{\"data\":0}" //1表示修改成功，0表示失败
		this.ServeJSON()
		return
	}

	obj, _ := models.IpGetById(ip.Id)
	if this.isPost() {
		obj.Address = ip.Address
		obj.CenterName = ip.CenterName
		obj.CenterPosition = ip.CenterPosition
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

type myDelStruct struct {
	Id string
}

func (this *IpSetController) DelData() {
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
		err = models.IpDelete(i)
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
