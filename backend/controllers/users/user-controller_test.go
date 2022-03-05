package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	config "prepboost.com/web/config"
)

func TestGetUser(t *testing.T) {
	// models.Init()
	config.DatabaseConnection()

	req, err := http.NewRequest("GET", "/users/5", nil)
	req.Header.Set("Content_Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	expectedUser := `[{"ID":5,"Username":"bhargav bffvs","Mobile":"+352989000990","Email":"rameshjhj@gmailcom","Password":"mypass","Paid":0,"Search_left":1,"Session_id":"akdhdfajdddfddsdfsddddsdf","Created_at":"0001-01-01T00:00:00Z","Updated_at":null}]`

	if rr.Body.String() != expectedUser {
		t.Errorf("handler returned unexpected body : got %v want %v", rr.Body.String(), expectedUser)
		fmt.Println(len(rr.Body.String()), "------", len(expectedUser))
	}

}

// Test - 2
// func Test_addUser(t *testing.T) {
// 	samplePerson := .Person{Id: 1, UserName: "username", Password: "password"}
// 	bytePerson, _ := json.Marshal(samplePerson)

// 	req, err := http.NewRequest("POST", "/signUp", bytes.NewReader(bytePerson))

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	//rr.Body = bytes.NewBuffer(bytePerson)
// 	handler := http.HandlerFunc(addUser)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
// 	}

// 	req.Header.Set("Token", samplePerson.Token)

// }
