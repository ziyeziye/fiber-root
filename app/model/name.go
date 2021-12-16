package model

type Name struct {
	ID       string `json:"id" gorm:"column:id;type:varchar(255)"`
	Name     string `json:"name" gorm:"column:name;type:varchar(255);comment:姓名"`       // 姓名
	Idnum    string `json:"idnum" gorm:"column:idnum;type:varchar(255);comment:身份id"`   // 身份id
	Address  string `json:"address" gorm:"column:address;type:varchar(255);comment:地址"` // 地址
	Phonenum int32  `json:"phoneNum" gorm:"column:phoneNum;type:tinyint;comment:电话号码"`  // 电话号码
}

// TableName returns the table name of the Name model
func (n *Name) TableName() string {
	return "name"
}
