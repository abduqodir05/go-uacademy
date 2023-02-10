package main

import (
	"fmt"
)

type Record struct {
	id    string
	name  string
	email string
}

type Database interface {
	Insert(record Record)
	Update(record Record)
	Delete(record Record)
	GetAll(limit, offset int) []Record
}

func (i *InMemoryDB) Insert(req Record) {
	for _, v := range i.records {
		if v.id == req.id {
			fmt.Println("ERROR")
		} else if v.id != req.id {
			fmt.Println("not exist this User3")
			return
		}
	}
	i.records = append(i.records, req)
	fmt.Println(req)
}
func (i *InMemoryDB) Update(req Record) {
	for idx, v := range i.records {
		if v.id == req.id {
			i.records[idx].name = req.name
			i.records[idx].email = req.email
		} else if v.id != req.id {
			fmt.Println("not exist this User")
			return
		}
	}
	fmt.Println(req)
}
func (i *InMemoryDB) Delete(req Record) {
	for ind, v := range i.records {
		if v.id == req.id {
			i.records = append(i.records[ind:], i.records[:ind]...)
			fmt.Println(req)
		} else if v.id != req.id {
			fmt.Println("not exist this User2")
			return
		}
	}
}

func (i *InMemoryDB) GetAll() []Record {
	return i.records
}

type InMemoryDB struct {
	records []Record
}

func main() {	
	obj1 := Record{
		id:    "1",
		name:  "John",
		email: "John@gmail.com",
	}
	obj2 := Record{
		id:    "2",
		name:  "mike",
		email: "mike@gmail.com",
	}
	obj3 := Record{
		id:    "3",
		name:  "Peter",
		email: "Peter@gmail.com",
	}
	indb := InMemoryDB{}
	indb.Insert(obj2)
	indb.Delete(obj3)
	indb.Update(obj1)
	indb.Update(obj2)
	indb.Update(obj3)
	res := indb.GetAll()
	fmt.Println(res)
}
