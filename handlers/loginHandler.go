package handlers

import (
	"encoding/json"
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

	// Logica de autenticación

	// Placeholder de autenticación
	if user.Username == "asdasd" {
		// Crea el coockie token con un hash identificador
		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: "5fd924625f6ab16a19cc9807c7c506ae1813490e4ba675f843d5a10e0baacdb8",
		})

		// Se dispara un redireccionamiento a la base de la app
		w.Header().Set("HX-Redirect", "/")
	}

	// En caso de no atenticar se devuelve error
	http.Error(w, "Error al autenticar", http.StatusBadRequest)
}
