package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

type userData struct {
	GhToken string `json:"ghToken"`
}

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
	defer r.Body.Close()
	var user userData

	session, _ := Store.Get(r, "session-name")

	// Se parsea el body a nuestra struc
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Se realizan validaciones pertinentes

	// Se guardan los datos de la sesion del usuario
	session.Values["TOKEN"] = user.GhToken // Simulación de un ID de usuario
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
