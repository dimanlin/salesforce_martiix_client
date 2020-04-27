package salesforce_marttix_client

import (
	"github.com/dimanlin/salesforce_martiix_client/internal/auth"
	"github.com/dimanlin/salesforce_martiix_client/pkg/salesforce_client"
)

func CreateMClient(host string, email string, password string, api_version string) salesforce_client.Smc {
	client := auth.CreteClient(host, email, password, api_version)
	smc := &salesforce_client.Smc{Client: *client}
	return *smc
}
