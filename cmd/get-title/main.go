package main

import (
	"fmt"

	"github.com/simonbuckner/sharepoint"
)

func main() {

	tenantId := "00000000-1111-2222-3333-444444444444"
	clientId := "55555555-6666-7777-8888-999999999999"
	certPath := "./cert.pfx"
	certPass := "0123456789QWERTY"

	sp := sharepoint.NewSharePoint(tenantId, clientId, certPath, certPass)

	site, err := sp.ConnectToSite("https://mytenant.sharepoint.com/sites/MySite")
	if err != nil {
		panic(err)
	}

	title, err := site.GetTitle()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to site %s", title)
}
