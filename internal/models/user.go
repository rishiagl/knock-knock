package models

type User struct {
	Id []byte			`json:"id"`
	Usrname string		`json:"name"`
	Passwd string		`json:"passwd"`
}

func GetUser(n string, p string) {

}
