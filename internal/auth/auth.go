package auth

import "github.com/simpleforce/simpleforce"

func CreteClient(host string, user string, password string, APIVersion string) *simpleforce.Client {
	sfToken := ""

	client := simpleforce.NewClient(host, simpleforce.DefaultClientID, APIVersion)
	if client == nil {
		// handle the error

		return nil
	}

	err := client.LoginPassword(user, password, sfToken)
	if err != nil {
		// handle the error

		return nil
	}

	// Do some other stuff with the client instance if needed.

	return client
}

