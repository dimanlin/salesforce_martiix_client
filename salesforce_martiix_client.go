package salesforce_marttix_client

import (
	"bitbucket.org/dimanlin/salesforce_marttix_client/internal/auth"
	"bitbucket.org/dimanlin/salesforce_marttix_client/pkg/salesforce_client"
)

func CreateMClient(host string, email string, password string, api_version string) salesforce_client.Smc {
	client := auth.CreteClient(host, email, password, api_version)
	smc := &salesforce_client.Smc{Client: *client}
	return *smc
}
