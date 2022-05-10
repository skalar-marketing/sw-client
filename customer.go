package sw

const (
	loginCustomerEndpoint       = "/account/login"
	logoutCustomerEndpoint      = "/account/logout"
	confirmRegistrationEndpoint = "/account/register-confirm"
	registerCustomerEndpoint    = "/account/register"
)

// CustomerClient handles all requests regarding customer login/registration.
type CustomerClient struct {
	client *Client
}

// Login returns a successful login.
func (client *CustomerClient) Login(login Login) (*LoginResult, error) {
	result := new(LoginResult)

	if err := client.client.performPost(loginCustomerEndpoint, login, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Logout returns a successful logout
func (client *CustomerClient) Logout() (*LogoutResult, error) {
	result := new(LogoutResult)

	if err := client.client.performPost(logoutCustomerEndpoint, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

//Confirmation confirms a customer registration.
func (client *CustomerClient) Confirmation(confirmRegistration ConfirmRegistration) (*LoginResult, error) {
	result := new(LoginResult) /* Aber vlt. auch nur Login? */

	if err := client.client.performPost(confirmRegistrationEndpoint, confirmRegistration, result); err != nil {
		return result, nil
	}
	return result, nil
}

// Registration returns a successful user.
func (client *CustomerClient) Registration(register Register) (*RegisterResult, error) {
	result := new(RegisterResult)

	if err := client.client.performPost(registerCustomerEndpoint, register, result); err != nil {

		return result, nil
	}

	return result, nil

}
