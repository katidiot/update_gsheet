// A simple example program that illustrates the use of oauth2 authentication,
// for server-to-server interaction. In this case the application proves its
// own identity to the API, with the use of an authentication key, without the
// need for user consent. Delegation to domain-wide authority must be granted
// to the service account. More information can be found on the following link:
// https://developers.google.com/identity/protocols/OAuth2ServiceAccount
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

const (
	// Where to find the Id: https://developers.google.com/sheets/api/guides/concepts
	sheetId       string = "the spreadsheet ID"

	// How to get key: https://developers.google.com/identity/protocols/OAuth2ServiceAccount#creatinganaccount
	key           string = "the service account key"

	// The email used to apply the changes to the sheet.
	delegateEmail string = "the delagation email"
)

// AuthGSpreadsheet authenticates the provided key to the API. The impersonation
// of the user account also happens through this function when a delegated domain-wide
// access to the service account has been granted. The function returns a *sheets.Service
// object that will be used to update the sheet or an error.
func AuthGSpreadsheet(key, delegate string) (srv *sheets.Service, err error) {
	jsonKey, err := ioutil.ReadFile(key)
	if err != nil {
		return nil, err
	}
	// Use a JSON key file to read the credentials that authorize and authenticate the requests
	conf, err := google.JWTConfigFromJSON(jsonKey, sheets.SpreadsheetsScope)
	if err != nil {
		return nil, err
	}
	// Impersonate user
	conf.Subject = delegate
	srv, err = sheets.New(conf.Client(context.Background()))
	if err != nil {
		return nil, err
	}
	return srv, nil
}

// updateGSpreadsheet is an example function that will update spreadsheet (sheetId) 
// with myval values. More info on how to use the sheets API can be found here:
// https://developers.google.com/sheets/api/
func updateGSpreadsheet(srv *sheets.Service) error {
	var vr sheets.ValueRange
	writeRange := "A1"
	myval := []interface{}{"One", "Two", "Three"}
	vr.Values = append(vr.Values, myval)
	// Update sheetId's writeRange cells with myval values.
	_, err := srv.Spreadsheets.Values.Update(sheetId, writeRange, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		return fmt.Errorf("Unable to update data in GSheet. \n%v", err)
	}
	return nil
}

func main() {
	srv, err := AuthGSpreadsheet(key, delegateEmail)
	if err != nil {
		log.Fatal(err)
	}
	err = updateGSpreadsheet(srv)
	if err != nil {
		log.Fatal(err)
	}
}
