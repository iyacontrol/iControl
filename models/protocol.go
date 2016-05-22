package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Protocol struct {
	Id        int
	Com       int
	BaudRate  int
	DataBit   int
	ParityBit int
	StopBit   int
}

func (p *Protocol) TableName() string {
	return TableName("protocol")
}

func ProtocolGetList() ([]*Protocol, error) {
	list := make([]*Protocol, 0)

	query := orm.NewOrm().QueryTable(TableName("protocol"))
	query.OrderBy("-id").All(&list)

	return list, nil
}

func ProtocolGetById(id int) (*Protocol, error) {
	obj := &Protocol{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func ProtocolAdd(obj *Protocol) (int64, error) {
	if obj.Com == 0 {
		return 0, fmt.Errorf("Com口不能为0")
	}
	if obj.BaudRate == 0 {
		return 0, fmt.Errorf("波特率不能为0")
	}
	if obj.DataBit == 0 {
		return 0, fmt.Errorf("数据位不能为0")
	}
	return orm.NewOrm().Insert(obj)
}

func ProtocolDelete(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("protocol")).Filter("id", id).Delete()
	return err
}

func (p *Protocol) Update(fields ...string) error {
	if p.Com == 0 {
		return fmt.Errorf("Com口不能为0")
	}
	if p.BaudRate == 0 {
		return fmt.Errorf("波特率不能为0")
	}
	if p.DataBit == 0 {
		return fmt.Errorf("数据位不能为0")
	}
	if _, err := orm.NewOrm().Update(p, fields...); err != nil {
		return err
	}
	return nil
}
