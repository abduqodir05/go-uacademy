package main

import "fmt"

type User struct {
	Username string
	Password string
}

type Auth interface {
	Register(user User) error
	Login(username, password string) bool
}

func (i *InMemoryAuth) Register(req User) {
	for _, v := range i.users {
		if v.Username == req.Username {
			fmt.Println("this Username already exist")
			return
		} else {
			i.users = append(i.users, req)
			fmt.Println("added")
		}
	}
}
func (i *InMemoryAuth) Login(req User) bool {
	for _, v := range i.users {
		if v.Username == req.Username && v.Password == req.Password {
			return true
		}

	}
	return false
}

type InMemoryAuth struct {
	users []User
}

func main() {

	Db := InMemoryAuth{}
	for {
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. tamonlash")
		var input int
		fmt.Scan(&input)

		if input == 1 {
			var register_name string
			fmt.Println("username kiriting")
			fmt.Scan(&register_name)
			var register_pas string
			fmt.Println("password kiriting")
			fmt.Scan(&register_pas)
			Db.Register(User{Username: register_name, Password: register_pas})
			fmt.Println("registratsiyadan o'tdingiz")
			return

		} else if input == 2 {
			var login_name string
			fmt.Println("Username kiriting")
			fmt.Scan(&login_name)
			var login_pas string
			fmt.Println("password kiriting")
			fmt.Scan(&login_pas)
			err := Db.Login(User{Username: login_name, Password: login_pas})
			if !err {
				fmt.Println("error")
			}
		} else if input == 0 {
			return
		}
	}
}
