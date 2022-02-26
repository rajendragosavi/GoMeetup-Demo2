package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Employee struct {
	FirstName        string
	LastName         string
	Department       string
	ProfessionalBand string
	Location         string
}

type EmployeeModel struct {
	DB *sql.DB
}

type EmpInterface interface {
	GetAllEmployee() ([]Employee, error)
}

func InitDB() (*sql.DB, error) {
	log.Debug().Msg("Init DB function is running... \n")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "password", "shopdb")
	Db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal().Msgf("error - %v \n", err)
	}
	// defer Db.Close()
	return Db, nil
}

func (e *EmployeeModel) GetAllEmployee() ([]Employee, error) {
	log.Debug().Msg("GetAllEmployee function is running..")
	//psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "password", "shopdb")
	//DB, err := sql.Open("postgres", psqlconn)
	//if err != nil {
	//	log.Fatal().Msgf("error - %v \n", err)
	//}
	//defer DB.Close()

	rows, err := e.DB.Query("SELECT * FROM employee")
	if err != nil {
		log.Fatal().Msgf("Error %v - ", err)
	}
	defer rows.Close()
	var emp []Employee

	for rows.Next() {
		var e Employee
		err := rows.Scan(&e.FirstName, &e.LastName, &e.Department, &e.ProfessionalBand, &e.Location)
		if err != nil {
			log.Fatal().Msgf("Error - %v ", err)
		}
		emp = append(emp, e)
	}

	if err = rows.Err(); err != nil {
		log.Fatal().Msgf("error - %v", err)
	}
	return emp, nil
}
