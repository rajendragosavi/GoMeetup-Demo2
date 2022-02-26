package main

import (
	"github.com/rajendragosavi/GoMeetup-Demo2/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockEmployeeModel struct {
}

func (mock *mockEmployeeModel) GetAllEmployee() ([]models.Employee, error) {
	var emp []models.Employee
	emp = append(emp, models.Employee{FirstName: "Rajiv", LastName: "Gupta", Department: "Engineering", ProfessionalBand: "B1", Location: "India"})
	return emp, nil
}

func TestGetAllEmployees(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/getbooks", nil)

	env := Env{employee: &mockEmployeeModel{}}

	http.HandlerFunc(env.GetEmployees).ServeHTTP(rec, req)
	expected := "Rajiv,Gupta,Engineering,B1,India\n"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}

}
