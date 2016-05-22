package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Ip struct {
	Id             int
	Address        string
	CenterName     string
	CenterPosition string
}

func (i *Ip) TableName() string {
	return TableName("ip")
}

func IpGetList() ([]*Ip, error) {
	list := make([]*Ip, 0)

	query := orm.NewOrm().QueryTable(TableName("ip"))
	query.OrderBy("-id").All(&list)

	return list, nil
}

func IpGetById(id int) (*Ip, error) {
	obj := &Ip{
		Id: id,
	}

	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func IpAdd(obj *Ip) (int64, error) {
	if obj.Address == "" {
		return 0, fmt.Errorf("IP 不能为空")
	}
	if obj.CenterName == "" {
		return 0, fmt.Errorf("数据中心名不能为空")
	}
	if obj.CenterPosition == "" {
		return 0, fmt.Errorf("数据中心地址不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

func IpDelete(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("ip")).Filter("id", id).Delete()
	return err
}

func (i *Ip) Update(fields ...string) error {
	if i.Address == "" {
		return fmt.Errorf("IP 不能为空")
	}
	if i.CenterName == "" {
		return fmt.Errorf("数据中心名不能为空")
	}
	if i.CenterPosition == "" {
		return fmt.Errorf("数据中心地址不能为空")
	}
	if _, err := orm.NewOrm().Update(i, fields...); err != nil {
		return err
	}
	return nil
}
