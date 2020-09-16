package test

import (
	"log"
	"testing"

	"github.com/ticketmaster/nitro-go-sdk/client"
	"github.com/ticketmaster/nitro-go-sdk/model"
)

// TestAddNsacls retrieves resource.
func TestAddNsacls(t *testing.T) {
	///////////////////////////////////////////////////////////////////////
	var err error
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}
	defer o.Session.Logout()
	///////////////////////////////////////////////////////////////////////
	_, err = o.GetNsacl("DENY_10.73.46.22:80_10.0.0.0-10.255.255.255")
	if err == nil {
		err = o.DeleteNsacl("DENY_10.73.46.22:80_10.0.0.0-10.255.255.255")
		if err != nil {
			t.Error(err)
			return
		}
	}
	///////////////////////////////////////////////////////////////////////
	nsacl := new(model.NsaclAdd)
	nsacl.Nsacl.Aclname = "DENY_10.73.46.22:80_10.0.0.0-10.255.255.255"
	nsacl.Nsacl.Aclaction = "DENY"
	//nsacl.Nsacl.Srcportval = "45-1024"
	nsacl.Nsacl.Srcport = false
	nsacl.Nsacl.Destipval = "192.168.1.1"
	nsacl.Nsacl.Destip = true
	nsacl.Nsacl.State = "DISABLED"
	nsacl.Nsacl.Srcipval = "192.168.1.2"
	nsacl.Nsacl.Srcip = true
	nsacl.Nsacl.Protocol = "TCP"

	r, err := o.AddNsacl(*nsacl)
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("%s", toPrettyJSON(r))
	return
}

// TestGetNsacls retrieves resource.
func TestGetNsacls(t *testing.T) {
	///////////////////////////////////////////////////////////////////////
	TestAddNsacls(t)
	///////////////////////////////////////////////////////////////////////
	var err error
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}
	defer o.Session.Logout()
	///////////////////////////////////////////////////////////////////////
	r, err := o.GetNsacls()
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("%s", toPrettyJSON(r))
	return
}

// TestDeleteNsacl deletes resource.
func TestDeleteNsacl(t *testing.T) {
	///////////////////////////////////////////////////////////////////////
	TestAddNsacls(t)
	///////////////////////////////////////////////////////////////////////
	var err error
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}
	defer o.Session.Logout()
	///////////////////////////////////////////////////////////////////////
	err = o.DeleteNsacl("restrict")
	if err != nil {
		t.Error(err)
		return
	}
	return
}

// TestDisableNsacl creates a resource and disables it.
func TestDisableNsacl(t *testing.T) {
	///////////////////////////////////////////////////////////////////////
	TestAddNsacls(t)
	///////////////////////////////////////////////////////////////////////
	var err error
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}
	defer o.Session.Logout()
	///////////////////////////////////////////////////////////////////////
	nsacl := new(model.NsaclDisable)
	nsacl.Nsacl.Aclname = "restrict"
	o.DisableNsacl(*nsacl)
	return
}

// TestEnableNsacl creates a resource and enables it.
func TestEnableNsacl(t *testing.T) {
	///////////////////////////////////////////////////////////////////////
	TestAddNsacls(t)
	///////////////////////////////////////////////////////////////////////
	var err error
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}
	defer o.Session.Logout()
	///////////////////////////////////////////////////////////////////////
	nsacl := new(model.NsaclEnable)
	nsacl.Nsacl.Aclname = "restrict"
	o.EnableNsacl(*nsacl)
	return
}

// TestUpdateNsacl creates a resource and updates it.
func TestUpdateNsacl(t *testing.T) {
	///////////////////////////////////////////////////////////////////////
	TestAddNsacls(t)
	///////////////////////////////////////////////////////////////////////
	var err error
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}
	defer o.Session.Logout()
	///////////////////////////////////////////////////////////////////////
	nsacl := new(model.NsaclUpdate)
	nsacl.Nsacl.Aclname = "restrict"
	nsacl.Nsacl.Aclaction = "DENY"
	nsacl.Nsacl.Srcportval = "45-100"
	nsacl.Nsacl.Srcport = true
	nsacl.Nsacl.Destipval = "192.168.1.1"
	nsacl.Nsacl.Destip = true
	nsacl.Nsacl.Protocol = "TCP"

	r, err := o.UpdateNsacl(*nsacl)
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("%s", toPrettyJSON(r))
	return
}
