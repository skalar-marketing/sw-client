package sw

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//func TestCustomerClient_Login(t *testing.T) {
//	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
//	assert.NoError(t, err)
//	result, err := client.Customer.Login(Login{})
//
//	assert.NoError(t, err)
//	assert.NotNil(t, result)
//}

//func TestCustomerClient_Logout(t *testing.T) {
//	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
//	assert.NoError(t, err)
//	result, err := client.Customer.Logout(Logout{})
//
//	assert.NoError(t, err)
//	assert.NotNil(t, result)
//}

func TestCustomerClient_Registration(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Customer.Registration(Register{})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestCustomerClient_Confirmation(t *testing.T) {
	client, err := NewClient(os.Getenv("SW_HOST"), os.Getenv("SW_ACCESS_KEY"))
	assert.NoError(t, err)
	result, err := client.Customer.Confirmation(ConfirmRegistration{ //Bin mir nicht sicher

	})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
