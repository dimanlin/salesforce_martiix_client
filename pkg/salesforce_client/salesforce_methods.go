package salesforce_client

import (
	"fmt"
	"github.com/dimanlin/salesforce_martiix_client/internal/sql"
	"github.com/simpleforce/simpleforce"
)

type Smc struct {
	Client simpleforce.Client
}

type Marttix struct {
	Id string
	CallCenterPlatform string
	Name string
	Status string
	Version string
}

func (marti *Marttix) IdExist() bool {
	if len(marti.Id) > 0 {
		return true
	} else {
		return false
	}
}

type Client struct {
	Id string
	AccessCodeValidationRequired bool
	CallCenterPlatform string
	ContactEmail string
	ContactName string
	ContactPhoneNumber string
	CreatedDate string
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

func (sfc *Smc) GetClientByEmail(email string) (Client, error) {
	query := sql.Select + sql.ClientFields + sql.FromClient + " WHERE Contact_Email__c = " + "'" + email + "'"
	var results, err = sfc.Client.Query(query)
	if err != nil {
		fmt.Println("Error: ", err)
		client := Client{}
		return client, err
	}
	clients := parseClientRecord(results)
	return (*clients)[0], err
}

func parseMatriixRecord(results *simpleforce.QueryResult) *[]Marttix {
	var martiixs []Marttix
	for _, record := range results.Records {
		martiix := Marttix{}

		if record["Call_Center_Platform__c"] != nil {
			martiix.CallCenterPlatform = record["Call_Center_Platform__c"].(string)
		}

		if record["Id"] != nil {
			martiix.Id = record["Id"].(string)
		}

		if record["Name"] != nil {
			martiix.Name = record["Name"].(string)
		}

		if record["Status__c"] != nil {
			martiix.Status = record["Status__c"].(string)
		}

		if record["Version__c"] != nil {
			martiix.Version = record["Version__c"].(string)
		}

		martiixs = append(martiixs, martiix)
	}

	return &martiixs
}

func parseClientRecord(results *simpleforce.QueryResult) *[]Client {
	var clients []Client
	for _, record := range results.Records {
		var client Client
		if record["Access_Code_Validation_Required__c"] != nil {
			client.AccessCodeValidationRequired = record["Access_Code_Validation_Required__c"].(bool)
		}

		if record["Call_Center_Platform__c"] != nil {
			client.CallCenterPlatform = record["Call_Center_Platform__c"].(string)
		}

		if record["Contact_Email__c"] != nil {
			client.ContactEmail = record["Contact_Email__c"].(string)
		}

		if record["Contact_Name__c"] != nil {
			client.ContactName = record["Contact_Name__c"].(string)
		}

		if record["Contact_Phone_Number__c"] != nil {
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
	return &clients
}

func (sfc *Smc) GetClientAll() (*[]Client, error) {
	query := sql.Select + sql.ClientFields + sql.FromClient
	var results, err = sfc.Client.Query(query)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	clients := parseClientRecord(results)
	return clients, err
}

func (sfc *Smc) GetAllMarttix() *[]Marttix {
	query := sql.Select + sql.MarttiFields + sql.FromMartti
	var results, err = sfc.Client.Query(query)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	martiix := parseMatriixRecord(results)
	return martiix
}

//func (sfc *Smc) GetClientsByMarttixId(martiixId string) (*[]Client, error) {
//	query := "SELECT Id, Client__r.Id, Client__r.Name__c, Client__r.Micro_Call_Center__c, Client__r.Physical_Address_State__c FROM MARTTIX__c WHERE Id = '" + martiixId + "'"
//	var results, err = sfc.Client.Query(query)
//	//fmt.Println(results)
//	clients := parseClientRecord(results)
//	return clients, err
//}

//func (sfc *Smc) GetClientsByMarttixId(martiixId string) {
//	//query := "SELECT Id, Client__r.Id, Client__r.Name__c, Client__r.Micro_Call_Center__c, Client__r.Physical_Address_State__c FROM MARTTIX__c WHERE Id = '" + martiixId + "'"
//	query := "SELECT MARTTIX__c.Id FROM ClientX__c"
//	var results, _ = sfc.Client.Query(query)
//	//fmt.Println(results)
//	for _, record := range results.Records {
//		// access the record as SObjects.
//		fmt.Println(record)
//	}
//}
