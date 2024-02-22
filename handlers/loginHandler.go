package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var Store *sessions.FilesystemStore

func init() {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}
	Store = sessions.NewFilesystemStore("./sessions", []byte(os.Getenv("SESSION_SECRET")))

	Store.Options = &sessions.Options{
		MaxAge:   604800, // una semana en segundos
		HttpOnly: true,
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/login.html"))
	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	token := r.PostFormValue("token")

	session, _ := Store.Get(r, "session-name")

	// Se realizan validaciones pertinentes

	if session.Values["TOKEN"] != nil {
		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}

	// Se guardan los datos de la sesion del usuario
	session.Values["TOKEN"] = token
	fmt.Println(session.Values["TOKEN"])
	session.Save(r, w)

	// Se hace un redirect a la base del proyecto
	w.Header().Set("HX-Redirect", "/")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session-name")
	session.Options.MaxAge = -1
	session.Save(r, w)
	w.Header().Set("HX-Redirect", "/")
}
