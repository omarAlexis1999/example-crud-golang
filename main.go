package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// conectionBD create the connection to the database
// return a pointer with the connection to our database
func conectionBD() *sql.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	driver := os.Getenv("DB_DRIVER")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	conection, err := sql.Open(driver, user+":"+password+"@tcp("+host+":"+port+")/"+name)
	if err != nil {
		panic(err.Error())
	}
	return conection
}

var templateG = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", Start)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insert", Insert)

	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)

	fmt.Println("Server Running...")
	http.ListenAndServe(":8080", nil)
}

type Employee struct {
	Id    int
	Name  string
	Email string
}

// Start receives a w (Response Writer) and r (the Request http)
func Start(w http.ResponseWriter, r *http.Request) {
	conection := conectionBD()
	getRegisters, err := conection.Query("SELECT * from employee")

	if err != nil {
		panic(err.Error())
	}
	employe := Employee{}
	arrayEmployees := []Employee{}
	for getRegisters.Next() {
		var id int
		var name, email string
		err = getRegisters.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		employe.Id = id
		employe.Name = name
		employe.Email = email
		arrayEmployees = append(arrayEmployees, employe)
	}
	templateG.ExecuteTemplate(w, "home", arrayEmployees)
}

func Create(w http.ResponseWriter, r *http.Request) {
	templateG.ExecuteTemplate(w, "create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")

		conection := conectionBD()
		insertRegister, err := conection.Prepare("Insert into employee (name, email) values(?,?)")

		if err != nil {
			panic(err.Error())
		}
		insertRegister.Exec(name, email)

		http.Redirect(w, r, "/", 301)

	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")

	conection := conectionBD()
	deleteRegister, err := conection.Prepare("DELETE from employee where id=?")

	if err != nil {
		panic(err.Error())
	}
	deleteRegister.Exec(idEmployee)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idEmployee := r.URL.Query().Get("id")

	conection := conectionBD()
	getEmploye, err := conection.Query("SELECT * from employee where id=?", idEmployee)

	if err != nil {
		panic(err.Error())
	}

	employe := Employee{}
	for getEmploye.Next() {
		var id int
		var name, email string
		err = getEmploye.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		employe.Id = id
		employe.Name = name
		employe.Email = email
	}

	templateG.ExecuteTemplate(w, "edit", employe)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		email := r.FormValue("email")

		conection := conectionBD()
		editRegister, err := conection.Prepare("UPDATE employee SET name=?,email=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}
		editRegister.Exec(name, email, id)
		http.Redirect(w, r, "/", 301)
	}
}
