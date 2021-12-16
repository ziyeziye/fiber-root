package model

type API struct {
	ID          int    `json:"id" gorm:"primaryKey;column:id;type:int auto_increment;comment:接口id"`                             // 接口id
	Path        string `json:"path" gorm:"column:path;type:varchar(255);uniqueIndex:path_method;comment:接口路径"`                  // 接口路径
	Description string `json:"description" gorm:"column:description;type:varchar(255);comment:接口描述"`                            // 接口描述
	Group       string `json:"group" gorm:"column:group;type:varchar(255);index:group;comment:接口属组"`                            // 接口属组
	Method      string `json:"method" gorm:"column:method;type:varchar(255);uniqueIndex:path_method;default:POST;comment:接口方法"` // 接口方法
	CreateTime  int64  `json:"create_time" gorm:"column:create_time;type:bigint;comment:创建时间"`                                  // 创建时间
	UpdateTime  int64  `json:"update_time" gorm:"column:update_time;type:bigint;comment:更新时间"`                                  // 更新时间
}

// TableName returns the table name of the API model
func (a *API) TableName() string {
	return "api"
}
