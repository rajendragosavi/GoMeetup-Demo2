package main

import (
	"database/sql"
	"fmt"
	"github.com/rajendragosavi/GoMeetup-Demo2/models"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func main() {
	log.Debug().Msg("Main Started....")
	// Initialise the connection pool.
	db, err := models.InitDB()
	if err != nil {
		log.Fatal().Msgf("Error in creating DB connection - %v ", err)
	}

	// Create an instance of Env containing the connection pool.
	env := &Env{db: db}

	http.HandleFunc("/employee", env.GetEmployees)
	err = http.ListenAndServe(":8585", nil)
	if err != nil {
		log.Fatal().Msgf("%+v \n", err)
	}
}

func (e *Env) GetEmployees(w http.ResponseWriter, req *http.Request) {
	employeelist, err := models.GetAllEmployee()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	for _, i := range employeelist {
		fmt.Fprintf(w, " %s, %s, %s,%s,%s \n", i.FirstName, i.LastName, i.Department, i.ProfessionalBand, i.Location)
	}
}
