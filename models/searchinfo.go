package models

type Searchinfo struct {
	Id int64 `gorm:"primary_key"`
	Type int64
	Value string
	TypeId int64 `gorm:"column:typeId"`
}


func (c *Searchinfo)TableName() string {
	return "searchinfo"
}
