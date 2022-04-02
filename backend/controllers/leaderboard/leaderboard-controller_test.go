package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	models "prepboost.com/web/models"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

//Test-1
//All time Leaderboard test cases for the users
func TestAlltimeLeaderboard(t *testing.T) {
	models.Init()

	req, err := http.NewRequest("GET", "/leaderboard/alltime", nil)
	req.Header.Set("Content_Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(LeaderboardAlltime)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	checkResponseCode(t, http.StatusOK, rr.Code)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	expectedOutput := `[{"ID":1,"Username":"bhargav bvs","Email":"bhargav@gmail.com","Mobile":"+13528707176","Answered":"2"},{"ID":2,"Username":"chaitanya","Email":"","Mobile":"+13528707156","Answered":"2"},{"ID":3,"Username":"arhul","Email":"","Mobile":"+3544888888","Answered":"2"},{"ID":4,"Username":"bhargav lolla b","Email":"ramddeshf@gmailcom","Mobile":"+35290dd7448sds","Answered":"2"},{"ID":5,"Username":"bhargav bffvs","Email":"rameshjhj@gmailcom","Mobile":"+352989000990","Answered":"1"}]`

	if rr.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body : got %v want %v", rr.Body.String(), expectedOutput)
		fmt.Println(len(rr.Body.String()), "------", len(expectedOutput))
	}
}

//Test-2
//All time Leaderboard test cases for the users
func TestDailyLeaderboard(t *testing.T) {
	models.Init()

	req, err := http.NewRequest("GET", "/leaderboard/daily", nil)
	req.Header.Set("Content_Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(LeaderboardDaily)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	checkResponseCode(t, http.StatusOK, rr.Code)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	expectedOutput := `[{"ID":1,"Username":"bhargav bvs","Email":"bhargav@gmail.com","Mobile":"+13528707176","Answered":"2"},{"ID":2,"Username":"chaitanya","Email":"","Mobile":"+13528707156","Answered":"2"},{"ID":3,"Username":"arhul","Email":"","Mobile":"+3544888888","Answered":"2"},{"ID":4,"Username":"bhargav lolla b","Email":"ramddeshf@gmailcom","Mobile":"+35290dd7448sds","Answered":"2"},{"ID":5,"Username":"bhargav bffvs","Email":"rameshjhj@gmailcom","Mobile":"+352989000990","Answered":"1"}]`

	if rr.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body : got %v want %v", rr.Body.String(), expectedOutput)
		fmt.Println(len(rr.Body.String()), "------", len(expectedOutput))
	}
}
