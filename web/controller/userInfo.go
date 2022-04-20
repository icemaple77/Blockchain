package controller

import "education/service"

type Application struct {
	Setup *service.ServiceSetup
}

type User struct {
	LoginName string
	Password  string
	IsAdmin   string
}

var users []User

func init() {

	admin := User{LoginName: "N10629297", Password: "123456", IsAdmin: "T"}
	Doris := User{LoginName: "Doris", Password: "123456", IsAdmin: "F"}
	bob := User{LoginName: "alice", Password: "123456", IsAdmin: "F"}
	jack := User{LoginName: "bob", Password: "123456", IsAdmin: "F"}

	users = append(users, admin)
	users = append(users, Doris)
	users = append(users, bob)
	users = append(users, jack)

}
