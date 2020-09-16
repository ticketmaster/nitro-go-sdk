package test

import (
	"log"
	"testing"

	"github.com/ticketmaster/nitro-go-sdk/client"
	"github.com/ticketmaster/nitro-go-sdk/model"
)

func TestGetLbmonitors(t *testing.T) {
	var err error
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()

	r, err := o.GetLbmonitors()
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("%+v", r)
}

// TestAddDeleteLbmonitor creates a resource and deletes it.
func TestAddDeleteLbmonitor(t *testing.T) {
	var err error
	Name := "foomon"
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	resp, err := o.GetLbmonitor(Name)

	if err == nil {
		err = o.DeleteLbmonitor(Name, resp.Type)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	req := model.LbmonitorAdd{}
	req.Lbmonitor.Monitorname = Name
	req.Lbmonitor.Respcode = []string{"200"}
	req.Lbmonitor.Resptimeout = 10
	req.Lbmonitor.Retries = 1
	req.Lbmonitor.Type = "HTTP"
	req.Lbmonitor.Interval = 13
	req.Lbmonitor.Httprequest = "GET /"

	r, err := o.AddLbmonitor(req)
	if err != nil {
		t.Error(err)
		return
	}

	log.Printf("%+v", r)

	// Cleanup
	err = o.DeleteLbmonitor(r.Monitorname, r.Type)
	if err != nil {
		t.Error(err)
	}
}
