package main

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}
type UserMeta struct {
	Visits int
}

func main() {
	log.Println("experimantal main")
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "john smith",
		Age:  123,
		Meta: UserMeta{
			Visits: 6,
		},
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
