package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type userData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/login.html"))
	tmpl.Execute(w, nil)
}

func Auth(w http.ResponseWriter, r *http.Request) {

	var user userData

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Fprintf(w, "Nombre enviado: %s", user.Username)
}
