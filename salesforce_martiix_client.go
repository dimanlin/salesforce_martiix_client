package main

import (
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


	results, _ := client.GetClientAll()
	//results, _ := client.GetAllMarttis()

	var clients []Client
	//clientS := clientSlice{prt: *clientO}

	for _, record := range results.Records {
		var client Client
		if record["Access_Code_Validation_Required__c"] != nil {
			client.AccessCodeValidationRequired = record["Access_Code_Validation_Required__c"].(bool)
		}

		if record["Call_Center_Platform__c"] != nil {
			client.CallCenterPlatform = record["Call_Center_Platform__c"].(string)
		}

		if record["ContactEmail"] != nil {
			client.ContactEmail = record["Contact_Email__c"].(string)
		}

		if record["ContactName"] != nil {
			client.ContactName = record["Contact_Name__c"].(string)
		}

		if record["ContactPhoneNumber"] != nil {
			client.ContactPhoneNumber = record["Contact_Phone_Number__c"].(string)
		}

		if record["CreatedDate"] != nil {
			client.CreatedDate = record["CreatedDate"].(string)
		}

		if record["Id"] != nil {
			client.Id = record["Id"].(string)
		}

		if record["Micro_Call_Center_Router_Workspace_ID__c"] != nil {
			client.MicroCallCenterRouterWorkspaceId = record["Micro_Call_Center_Router_Workspace_ID__c"].(string)
		}

		if record["Micro_Call_Center__c"] != nil {
			client.MicroCallCenter = record["Micro_Call_Center__c"].(bool)
		}

		if record["name"] != nil {
			client.Name = record["name"].(string)
		}

		if record["Name__c"] != nil {
			client.NameC = record["Name__c"].(string)
		}

		if record["PEM_Name__c"] != nil {
			client.PEMEmail = record["PEM_Name__c"].(string)
		}

		if record["PEM_Phone_Number__c"] != nil {
			client.PEMNameC = record["PEM_Phone_Number__c"].(string)
		}

		if record["PEM_Phone_Number__c"] != nil {
			client.PEMPhoneNumber = record["PEM_Phone_Number__c"].(string)
		}

		if record["Photo_Document_ID__c"] != nil {
			client.PhotoDocumentID = record["Photo_Document_ID__c"].(string)
		}

		if record["Photo__c"] != nil {
			client.Photo = record["Photo__c"].(string)
		}

		if record["Photo__c"] != nil {
			client.Photo = record["Photo__c"].(string)
		}

		if record["Physical_Address_Line_1__c"] != nil {
			client.PhysicalAddressLine1 = record["Physical_Address_Line_1__c"].(string)
		}

		if record["Physical_Address_Line_2__c"] != nil {
			client.PhysicalAddressLine2 = record["Physical_Address_Line_2__c"].(string)
		}

		if record["Physical_Address_State__c"] != nil {
			client.PhysicalAddressState = record["Physical_Address_State__c"].(string)
		}

		if record[" Physical_Address_Zip__c"] != nil {
			client.PhysicalAddressZip = record[" Physical_Address_Zip__c"].(uint)
		}

		if record[" Physical_Location_City__c"] != nil {
			client.PhysicalLocationCity = record[" Physical_Location_City__c"].(string)
		}

		if record[" Status__c"] != nil {
			client.Status = record[" Status__c"].(string)
		}

		if record[" Technical_Support_Number__c"] != nil {
			client.TechnicalSupportNumber = record[" Technical_Support_Number__c"].(string)
		}

		clients = append(clients, client)
	}
}

