package models

type User struct {
	Id int			`json:"id"`
	Name string		`json:"name"`
	Passwd string	`json:"passwd"`
}
