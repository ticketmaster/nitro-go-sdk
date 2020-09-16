package test

import (
	"log"
	"testing"

	"github.com/tickemaster/nitro-go-sdk/client"
)

func TestGetLbmonbindingsbinding(t *testing.T) {
	var err error
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()

	r, err := o.GetLbmonbindingsBindings()
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("%+v", r)
}
