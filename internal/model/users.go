package model

type User struct {
	Id             uint   `json:"id" `
	PassportNumber string `json:"passport_number" gorm:"type:varchar(50)"`
	Surname        string `json:"surname" gorm:"type:varchar(100)"`
	Name           string `json:"name" gorm:"type:varchar(100)"`
	Patronymic     string `json:"patronymic" gorm:"type:varchar(100)"`
	Address        string `json:"address" gorm:"type:varchar(255)"`
}
