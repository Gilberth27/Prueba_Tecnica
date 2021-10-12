package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func conec() (conexion *sql.DB) {

	Driver := "mysql"
	Usuario := "root"
	Contrasena := ""
	Nombre := "tickets"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasena+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("conextado")
	return conexion

}

var modelo = template.Must(template.ParseGlob("Templates/*"))

type tickets struct {
	ID      int    `json:"ID"`
	USUARIO string `json:"USUARIO"`
	FCH_CRE string `json:"FCH_CRE"`
	FCH_UPD string `json:"FCH_UPD"`
	ESTADO  string `json:"ESTADO"`
}

type allTask []tickets

var Tickets = allTask{
	{
		ID:      1,
		USUARIO: "Gilberth",
		FCH_CRE: "11-10-2021",
		FCH_UPD: "12-10-2021",
		ESTADO:  "activo",
	},
}

func Mainlink(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Esta es Mi prueba tecnicaaaaa")
}

func ConsultaTickets(w http.ResponseWriter, r *http.Request) {
	modelo.ExecuteTemplate(w, "consultar", nil)

}
func CrearTickets(w http.ResponseWriter, r *http.Request) {
	modelo.ExecuteTemplate(w, "Crear", nil)
}

func ActualizarTickets(w http.ResponseWriter, r *http.Request) {
	modelo.ExecuteTemplate(w, "consultar", nil)
}

func RegistrarTickets(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		usuario := r.FormValue("usuario")
		creacion := r.FormValue("creacion")
		actualizar := r.FormValue("actualizar")
		estado := r.FormValue("estado")
		Correcta := conec()
		insertarCasos, err := Correcta.Prepare("INSERT INTO casos (USUARIO,FECHA_CRE ,FECHA_UPD,ESTADO) VALUES (?,?,?,?)")

		if err != nil {
			panic((err.Error()))
		}
		insertarCasos.Exec(usuario, creacion, actualizar, estado)
		http.Redirect(w, r, "/", 301)
	}
}

func main() {

	link := mux.NewRouter().StrictSlash(true)
	link.HandleFunc("/", Mainlink)
	link.HandleFunc("/consultar", ConsultaTickets)
	link.HandleFunc("/crear", CrearTickets)
	link.HandleFunc("/registrar", RegistrarTickets)
	log.Println("Conexion ejecutandose")

	http.ListenAndServe(":5000", link)

}
