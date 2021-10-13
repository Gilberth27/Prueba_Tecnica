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
	fmt.Println("conectado")
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

func Mainlink(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Esta es Mi prueba tecnicaaaaa")
}

type Caso struct {
	ID        int
	USUARIO   string
	FECHA_CRE string
	FECHA_UPD string
	ESTADO    string
}

func CrearTickets(w http.ResponseWriter, r *http.Request) {
	modelo.ExecuteTemplate(w, "Crear", nil)
}

func UpdateTickets(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		usuario := r.FormValue("usuario")
		creacion := r.FormValue("creacion")
		actualizar := r.FormValue("actualizar")
		estado := r.FormValue("estado")
		Correcta := conec()
		UPDATECasos, err := Correcta.Prepare("UPDATE casos SET USUARIO=?,FECHA_CRE=?,FECHA_UPD=?,ESTADO=? WHERE ID=?")

		if err != nil {
			panic((err.Error()))
		}
		UPDATECasos.Exec(usuario, creacion, actualizar, estado, id)
		http.Redirect(w, r, "/consultar", 301)
	}
}

func ModificarTickets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("ID")
	fmt.Println(ID)
	Correcta := conec()
	EditarCasos, err := Correcta.Query("SELECT * FROM casos where ID=?", ID)
	caso := Caso{}
	for EditarCasos.Next() {
		var ID int
		var USUARIO, FECHA_CRE, FECHA_UPD, ESTADO string

		err = EditarCasos.Scan(&ID, &USUARIO, &FECHA_CRE, &FECHA_UPD, &ESTADO)
		if err != nil {
			panic(err.Error())
		}
		caso.ID = ID
		caso.USUARIO = USUARIO
		caso.FECHA_CRE = FECHA_CRE
		caso.FECHA_UPD = FECHA_UPD
		caso.ESTADO = ESTADO
	}
	fmt.Println(caso)
	modelo.ExecuteTemplate(w, "editar", caso)
}
func EliminarTickets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("ID")
	fmt.Println(ID)

	Correcta := conec()
	EliminarCasos, err := Correcta.Prepare("DELETE FROM casos where ID=?")
	if err != nil {
		panic((err.Error()))
	}
	EliminarCasos.Exec(ID)
	http.Redirect(w, r, "/consultar", 301)
}

func ConsultaTickets(w http.ResponseWriter, r *http.Request) {

	Correcta := conec()
	ConsultarCasos, err := Correcta.Query("SELECT * FROM casos")

	if err != nil {
		panic((err.Error()))
	}
	caso := Caso{}
	Acaso := []Caso{}

	for ConsultarCasos.Next() {
		var ID int
		var USUARIO, FECHA_CRE, FECHA_UPD, ESTADO string

		err = ConsultarCasos.Scan(&ID, &USUARIO, &FECHA_CRE, &FECHA_UPD, &ESTADO)
		if err != nil {
			panic(err.Error())
		}
		caso.ID = ID
		caso.USUARIO = USUARIO
		caso.FECHA_CRE = FECHA_CRE
		caso.FECHA_UPD = FECHA_UPD
		caso.ESTADO = ESTADO

		Acaso = append(Acaso, caso)

	}

	modelo.ExecuteTemplate(w, "consultar", Acaso)
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
		http.Redirect(w, r, "/consultar", 301)
	}
}

func main() {

	link := mux.NewRouter().StrictSlash(true)
	link.HandleFunc("/", Mainlink)
	link.HandleFunc("/consultar", ConsultaTickets)
	link.HandleFunc("/crear", CrearTickets)
	link.HandleFunc("/registrar", RegistrarTickets)
	link.HandleFunc("/eliminar", EliminarTickets)
	link.HandleFunc("/editar", ModificarTickets)
	link.HandleFunc("/actualizar", UpdateTickets)
	log.Println("Conexion ejecutandose")

	http.ListenAndServe(":5000", link)

}
