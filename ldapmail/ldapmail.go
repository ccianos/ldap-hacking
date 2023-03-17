package main

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap"
)

const (
	adminDN  = "cn=read-only-admin,dc=example,dc=com"
	password = "password"
	ldapURL  = "ldap://ldap.forumsys.com:389"
	base     = "dc=example,dc=com"
	nt       = "(uid=tesla)"
)

const (
	ScopeWholeSubtree = ldap.ScopeWholeSubtree
	NeverDerefAliases = ldap.NeverDerefAliases
)

var mail = []string{"mail"}

func main() {
	conn, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = conn.Bind(adminDN, password)
	if err != nil {
		log.Fatal(err)
	}

	searchRqst := ldap.NewSearchRequest(
		base,                                              // Base DN to search
		ScopeWholeSubtree, NeverDerefAliases, 0, 0, false, // scope, aliases & limits
		nt, mail, nil, // filter, attributes & controls
	)

	sr, err := conn.Search(searchRqst)
	if err != nil {
		log.Fatal(err)
	}

	entry := sr.Entries
	res := entry[0].GetAttributeValue("mail")
	fmt.Printf("\nEmail Address: %s\n", res)
}
