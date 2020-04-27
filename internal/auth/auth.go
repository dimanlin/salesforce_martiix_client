package auth

import (
	"fmt"
	"github.com/simpleforce/simpleforce"
)

func CreteClient(host string, user string, password string, APIVersion string) *simpleforce.Client {
	sfToken := ""

	client := simpleforce.NewClient(host, simpleforce.DefaultClientID, APIVersion)

	err := client.LoginPassword(user, password, sfToken)
	if err != nil {
		fmt.Println("Login error:", err)
		return nil
	}

	// Do some other stuff with the client instance if needed.

	return client
}

