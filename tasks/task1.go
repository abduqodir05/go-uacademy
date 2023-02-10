package main

import (
"go-uacademy/tasks"
)
var Users []tasks.User

func CreateUser(data tasks.User) {
  Users = append(Users, data)
}

func Getbyid(id int) tasks.User {
  for _, val := range Users {
    if val.Id == id {
      return val
    }
  }
  return tasks.User{}
}

func Update(data tasks.User) tasks.User {
  for idx, val := range Users {
    if val.Id == data.Id {
      Users[idx].Name = data.Name
      Users[idx].Surname = data.Surname
      return Users[idx]
    }
  }
  return tasks.User{}
}

func Delete(id int) tasks.User {
  for idx, val := range Users {
    if val.Id == id {
      Users = append(Users[:idx], Users[idx+1:]...)
      return val
    }
  }
  return tasks.User{}
}

func GetListUser() []tasks.User {
  return Users
}