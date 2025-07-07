package models
type User struct{
	Id int `json :id gorm:"primarykey"`
	Name string `json:name`
}