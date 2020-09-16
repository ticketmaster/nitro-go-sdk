package test

import (
	"log"
	"testing"

	"github.com/ticketmaster/nitro-go-sdk/client"
)

func TestGetNsip(t *testing.T) {
	var err error

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	r, err := o.GetNsips()
	if err != nil {
		t.Error(err)
		return
	}

	log.Printf("%+v", r)
}
