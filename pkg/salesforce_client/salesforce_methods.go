package salesforce_client

import (
	"bitbucket.org/dimanlin/salesforce_marttix_client/internal/sql"
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
