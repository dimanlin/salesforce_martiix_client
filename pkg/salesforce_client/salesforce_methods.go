package salesforce_client

import (
	"github.com/dimanlin/salesforce_martiix_client/internal/sql"
	"github.com/simpleforce/simpleforce"
)

type Smc struct {
	Client simpleforce.Client
}

func (sfc *Smc) GetClientByEmail(email string) (*simpleforce.QueryResult, error) {
	query := sql.Select + sql.ClientFields + sql.FromClient + " WHERE Contact_Email__c = " + "'" + email + "'"
	var results, err = sfc.Client.Query(query)
	return results, err
}

func (sfc *Smc) GetClientAll() (*simpleforce.QueryResult, error) {
	query := sql.Select + sql.ClientFields + sql.FromClient
	var results, err = sfc.Client.Query(query)
	return results, err
}

func (sfc *Smc) GetAllMarttis() (*simpleforce.QueryResult, error) {
	query := sql.Select + sql.MarttiFields + sql.FromMartti
	var results, err = sfc.Client.Query(query)
	return results, err
}

func (sfc *Smc) GetClientsByMarttiId(martiixId string) (*simpleforce.QueryResult, error) {
	query := "SELECT Id, Client__r.Id, Client__r.Name__c, Client__r.Micro_Call_Center__c, Client__r.Physical_Address_State__c FROM MARTTIX__c WHERE Id = '" + martiixId + "'"
	var results, err = sfc.Client.Query(query)
	return results, err
}

//func (sfc *Smc) GetClientsByMarttiId(martiixId string) (*simpleforce.QueryResult, error) {
//	query := "SELECT MARTTIX__c.Id, Id, Name FROM ClientX__c WHERE MARTTIX__c.Id = '" + martiixId + "'"
//	var results, err = sfc.Client.Query(query)
//	return results, err
//}

//func (sfc *Smc) GetMarttixByClientId(clientId string) (*simpleforce.QueryResult, error) {
//	query := "SELECT Id, Client__r.Id, Client__r.Name__c, Client__r.Micro_Call_Center__c, Client__r.Physical_Address_State__c FROM MARTTIX__c WHERE Id = '" + martiixId + "'"
//	var results, err = sfc.Client.Query(query)
//	return results, err
//}


