# Service Account OAuth2 for GSpreadsheets in Golang

A simple example program that illustrates the use of oauth2 authentication, for server-to-server interaction. In this case the application proves its own identity to the API, with the use of an authentication key, without the need for user consent. Delegation to domain-wide authority must be granted to the service account. More information can be found on [Google OAuth2 Service Account][goauth2].

### Usage
Change the following three constants:
```
sheetId       string = "the spreadsheet ID"
key           string = "the service account key"
delegateEmail string = "the delagation email"
```
* [sheetId] - Information on what is a sheet Id and where to find it can be found [here][concepts]. 
* [key] - [Here][serviceaccount] you may find information on how to acquire a service account key. Use full path to the location of the key, relative to the program.
* [delegateEmail] - The email used to apply the changes to the sheet.

From command line execute the following:
```
go run update_gsheet.go
```

### Extra
More information on how to use the sheets API is available [here][api].

[goauth2]: <https://developers.google.com/identity/protocols/OAuth2ServiceAccount>
[concepts]: <https://developers.google.com/sheets/api/guides/concepts>
[serviceaccount]: <https://developers.google.com/identity/protocols/OAuth2ServiceAccount#creatinganaccount>
[api]: <https://developers.google.com/sheets/api/>
