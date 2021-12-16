package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;column:id;type:int unsigned auto_increment"`
	Name     string `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Password string `json:"password" gorm:"column:password;type:varchar(255);not null"`
}

// TableName returns the table name of the User model
func (u *User) TableName() string {
	return "user"
}
