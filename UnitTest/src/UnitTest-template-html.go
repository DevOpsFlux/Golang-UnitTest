package main

import (
	"html/template"
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {

	//-----------------------------------------------------
	//fmt.Print("1. Go-Template Html Test ")
	//t, _ := template.ParseFiles("UnitTest-template.html")
	//t.Execute(w, "Body Text, Hello world~!!!")

	//-----------------------------------------------------
	//fmt.Print("2. Go-Template Html Test ")
	t, err := template.ParseFiles("UnitTest-template.html")

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	t.Execute(w, data)
	check(err)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

// http://127.0.0.1:8080/process

// https://dksshddl.tistory.com/entry/Go-web-programming-%ED%85%9C%ED%94%8C%EB%A6%BF%EA%B3%BC-%ED%85%9C%ED%94%8C%EB%A6%BF-%EC%97%94%EC%A7%84
