package model

type User struct {
	Id             int    `json:"id" `
	PassportNumber string `json:"passportNumber"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}

type UpdateUserInput struct {
	PassportNumber string `json:"passportNumber" sql:"passport_number"`
	Surname        string `json:"surname" sql:"surname"`
	Name           string `json:"name" sql:"name"`
	Patronymic     string `json:"patronymic" sql:"patronymic"`
	Address        string `json:"address" sql:"address"`
}

type UserIDResponse struct {
	ID int `json:"id" example:"1"`
}
