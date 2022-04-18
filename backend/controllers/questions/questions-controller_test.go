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

//Test-2
//To get the Questions topic wise from the database
func TestQuestionsTopicwise(t *testing.T) {
	models.Init()

	req, err := http.NewRequest("GET", "/topicwise/1/1", nil)
	req.Header.Set("Content_Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(GetTopicwiseQuestions)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	checkResponseCode(t, http.StatusOK, rr.Code)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	expectedOutput := `[{"ID":221,"Exam":"IAS","Type":"Prelims","Year":2018,"Question":"Increase in absolute and per capita real GNP do not connote a higher level of economic development, if","Answer":"C","Answered":"","Option1":"Industrial output fails to keep pace with agricultural output.","Option2":"Agricultural output fails to keep pace with industrial output.","Option3":"Poverty and unemployment increase.","Option4":"Imports grow faster than exports.","Explanation":"","Learning":"","Source":""},{"ID":531,"Exam":"IAS","Type":"Prelims","Year":2015,"Question":"With reference to India economy, consider the following statements:\\n\\n1. The rate of growth of real Gross Domestic Product has steadily increased in the last decade.\\n\\n2. The Gross Domestic Product at market prices (in rupees) has steadily increased in the last decade\\n\\nWhich of the statement given above is/are correct?","Answer":"B","Answered":"","Option1":"1 only","Option2":"2 only","Option3":"Both 1 and 2","Option4":"Neither 1 nor 2","Explanation":"","Learning":"","Source":""},{"ID":767,"Exam":"IAS","Type":"Prelims","Year":2013,"Question":"Economic growth in country X will necessarily have to occur if","Answer":"C","Answered":"","Option1":"There is technical progress in the world economy","Option2":"There is population growth in X","Option3":"There is capital formation in X","Option4":"The volume of trade grows in the world economy","Explanation":"","Learning":"","Source":""},{"ID":774,"Exam":"IAS","Type":"Prelims","Year":2013,"Question":"The national income of a country for a given period is equal to the","Answer":"D","Answered":"","Option1":"Total value of goods and services produced by the nationals","Option2":"Sum of total consumption and investment expenditure","Option3":"Sum of personal income of all individuals","Option4":"Money value of final goods and services produced.","Explanation":"","Learning":"","Source":""},{"ID":982,"Exam":"IAS","Type":"Prelims","Year":2011,"Question":"In the context of Indian economy, consider the following statements :\\n\\n1. The growth rate of GDP has steadily increased in the last five years.\\n\\n2. The growth rate in per capita income has steadily increased in the last five years.\\n\\nWhich of the statements given above is/ are correct?","Answer":"D","Answered":"","Option1":"1 only","Option2":"2 only","Option3":"Both 1 and 2","Option4":"Neither 1 nor 2","Explanation":"","Learning":"","Source":""}]`

	if rr.Body.String() != expectedOutput {
		t.Errorf("handler returned unexpected body : got %v want %v", rr.Body.String(), expectedOutput)
		fmt.Println(len(rr.Body.String()), "------", len(expectedOutput))
	}

}
