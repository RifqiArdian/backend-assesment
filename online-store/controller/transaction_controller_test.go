package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"online-store/model"
	"testing"
)

func TestTransactionController_Create(t *testing.T)  {
	transaction := model.InsertTransactionRequest{
		ProductId:  "6e1fac78-e31d-4507-888b-882c2ec309c0",
		Quantity:   2,
		Address:  	"bandung",
	}

	requestBody, _ := json.Marshal(transaction)
	request := httptest.NewRequest("POST","/transaction",bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.AddCookie(&http.Cookie{Name: "token",Value: "YzgwMWIyMTgtYzA3My00NjI3LTg4ZjgtYzBjNGRkZmRhMjg5"})
	request.AddCookie(&http.Cookie{Name: "user_id",Value: "31e60518-bbd5-49a9-bc7f-7ed9747e831e"})
	response, _ := app.Test(request)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.HttpResponse{}
	_ = json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	fmt.Println(webResponse)
}