package models

import "ctools"

type Tag struct {
	Id    int
	Type  string
	Num   int
	Value float64
}

func TagGetList() ([]Tag, error) {

	list := make([]Tag, 0)

	var t1 Tag
	t1.Id = ctools.SetId()
	t1.Type = ctools.SetType()
	t1.Num = ctools.SetNum()
	t1.Value = ctools.SetValue()
	list = append(list, t1)

	return list, nil
}
