package controller

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"online-store/model"
	"testing"
)

func TestAuthController_Register(t *testing.T)  {
	user := model.RegisterRequest{
		Name:     "a",
		Email:    "a@gmail.com",
		Password: "secret123",
		Address:  "bandung",
	}

	requestBody, _ := json.Marshal(user)
	request := httptest.NewRequest("POST","/auth/register",bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	response, _ := app.Test(request)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.HttpResponse{}
	_ = json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	jsonData, _ := json.Marshal(webResponse.Data)
	getUserResponse := model.GetUserResponse{}
	_ = json.Unmarshal(jsonData, &getUserResponse)
}

func TestAuthController_Login(t *testing.T)  {
	user := model.LoginRequest{
		Email:    "a@gmail.com",
		Password: "secret123",
	}

	requestBody, _ := json.Marshal(user)
	request := httptest.NewRequest("POST","/auth/login",bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	response, _ := app.Test(request)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.HttpResponse{}
	_ = json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	jsonData, _ := json.Marshal(webResponse.Data)
	getUserResponse := model.GetUserResponse{}
	_ = json.Unmarshal(jsonData, &getUserResponse)
}
