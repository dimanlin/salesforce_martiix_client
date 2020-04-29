package main

import (
	"fmt"
	"github.com/dimanlin/salesforce_martiix_client/internal/auth"
	"github.com/dimanlin/salesforce_martiix_client/pkg/salesforce_client"
)

func CreateMClient(host string, email string, password string, apiVersion string) salesforce_client.Smc {
	client := auth.CreteClient(host, email, password, apiVersion)
	smc := &salesforce_client.Smc{Client: *client}
	return *smc
}



type Client struct {
	AccessCodeValidationRequired bool
	CallCenterPlatform string
	ContactEmail string
	ContactName string
	ContactPhoneNumber string
	CreatedDate string
	Id string
	MicroCallCenterRouterWorkspaceId string
	MicroCallCenter bool
	Name string
	NameC string
	PEMEmail string
	PEMNameC string
	PEMPhoneNumber string
	PhotoDocumentID string
	Photo string
	PhysicalAddressLine1 string
	PhysicalAddressLine2 string
	PhysicalAddressState string
	PhysicalAddressZip uint
	PhysicalLocationCity string
	Status string
	TechnicalSupportNumber string
}

//var clients []Client
//clients = append(clients, Client{name: "Diman"})
//fmt.Println(clients[0])

func main() {
	client := CreateMClient("https://test.salesforce.com", "skrn@carenection.com.cvgdev", "Polycom!23dZRpwQlbm5WcRTqks0j8jbOS5", "41.0")


	results, err := client.GetClientAll()
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println((*results)[0])
	for i := 0; i < len(*results); i++  {
		fmt.Println((*results)[i].Id)
	}

}

