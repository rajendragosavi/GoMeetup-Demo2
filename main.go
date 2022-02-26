package main

import (
	"fmt"
	"github.com/rajendragosavi/GoMeetup-Demo2/models"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Env struct {
	employee models.EmpInterface
}

func main() {
	fmt.Println("Main Started....")
	// Initialise the connection pool.
	Db, err := models.InitDB()
	if err != nil {
		log.Fatal().Msgf("Error in creating DB connection - %v ", err)
	}

	defer Db.Close()
	// Create an instance of Env containing the connection pool.
	env := &Env{employee: &models.EmployeeModel{
		DB: Db,
	}}

	// Use env.booksIndex as the handler function for the /books route.
	http.HandleFunc("/employee", env.GetEmployees)
	http.ListenAndServe(":8585", nil)
}

func (e *Env) GetEmployees(w http.ResponseWriter, req *http.Request) {
	employeelist, err := e.employee.GetAllEmployee()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	for _, i := range employeelist {
		fmt.Fprintf(w, "%s,%s,%s,%s,%s\n", i.FirstName, i.LastName, i.Department, i.ProfessionalBand, i.Location)
	}
}
