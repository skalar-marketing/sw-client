package sw

const (
	loginCustomerEndpoint       = "/account/login"
	logoutCustomerEndpoint      = "/account/logout"
	confirmRegistrationEndpoint = "/account/register-confirm"
	registerCustomerEndpoint    = "/account/register"
)

// UserClient handles all requests regarding user login/registration.
type UserClient struct {
	client *Client
}

// Login returns a successful login.
func (client *UserClient) Login(login Login) (*LoginResult, error) {
	result := new(LoginResult)

	if err := client.client.performPost(loginCustomerEndpoint, login, result); err != nil {
		return nil, err
	}

	return result, nil
}

// Logout returns a successful logout
func (client *UserClient) Logout() {

}

// Registration returns a successful user.
func (client *UserClient) Registration() {

}
