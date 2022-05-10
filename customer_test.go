package sw

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//func TestCustomerClient_Login(t *testing.T) {}

//func TestCustomerClient_Logout(t *testing.T) {}

func TestCustomerClient_Confirmation(t *testing.T) {
	data := new(ConfirmRegistration)
	var req *http.Request
	ts, client := shopwareMock(t, new(LoginResult), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Customer.Confirmation(ConfirmRegistration{})
	checkResponse(t, req, "/store-api/account/register-confirm")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCustomerClient_Registration(t *testing.T) {
	data := new(Register)
	var req *http.Request
	ts, client := shopwareMock(t, new(RegisterResult), data, func(r *http.Request) {
		req = r
	})
	defer ts.Close()
	result, err := client.Customer.Registration(Register{})
	checkResponse(t, req, "/store-api/account/register")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
