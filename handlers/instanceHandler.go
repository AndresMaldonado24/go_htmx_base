package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Instance(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/instance.html"))
	profile := GetProfileData(r)
	tmpl.Execute(w, profile)
}

func Change(w http.ResponseWriter, r *http.Request) {
	opt := r.PostFormValue("env")
	cpn := r.PostFormValue("cpn")
	motive := r.PostFormValue("motive")

	fmt.Println(opt)
	fmt.Println(cpn)
	fmt.Println(motive)
}
