package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/tickemaster/nitro-go-sdk/client"
	"github.com/tickemaster/nitro-go-sdk/model"
)

func TestGetLbVserver(t *testing.T) {
	var err error
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}
	defer o.Session.Logout()
	r, err := o.GetLbvserver("Foo")
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("%+v", r)
}

// TestAddDeleteLbvserver creates a resource and deletes it.
func TestAddDeleteLbvserver(t *testing.T) {
	var err error
	Name := "Foo"
	Ipaddress := "1.2.3.4"
	State := "DISABLED"
	Port := 80
	Servicetype := "HTTP"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	_, err = o.GetLbvserver(Name)

	if err == nil {
		err = o.DeleteLbvserver(Name)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	req := model.LbvserverAdd{}
	req.Lbvserver.Name = Name
	req.Lbvserver.Ipv46 = Ipaddress
	req.Lbvserver.State = State
	req.Lbvserver.Port = Port
	req.Lbvserver.Servicetype = Servicetype

	r, err := o.AddLbvserver(req)
	if err != nil {
		t.Error(err)
		return
	}

	if r.Name != req.Lbvserver.Name || r.Ipv46 != req.Lbvserver.Ipv46 || r.Curstate != "OUT OF SERVICE" {
		t.Errorf("failed to add resource: %+v not equal to %+v", r, req)
		return
	}

	// Cleanup
	err = o.DeleteLbvserver(Name)
	if err != nil {
		t.Error(err)
	}
}

// TestDisableLbvserver creates a resource and disables it.
func TestDisableLbvserver(t *testing.T) {
	////////////////////////////////////////////////////////
	var err error
	Name := "Foo"
	Ipaddress := "1.2.3.4"
	State := "ENABLED"
	Port := 80
	Servicetype := "HTTP"
	////////////////////////////////////////////////////////
	reqLbvserveradd := model.LbvserverAdd{}
	reqLbvserveradd.Lbvserver.Name = Name
	reqLbvserveradd.Lbvserver.Ipv46 = Ipaddress
	reqLbvserveradd.Lbvserver.State = State
	reqLbvserveradd.Lbvserver.Port = Port
	reqLbvserveradd.Lbvserver.Servicetype = Servicetype
	////////////////////////////////////////////////////////
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()
	////////////////////////////////////////////////////////
	_, err = createTestLbvserver(reqLbvserveradd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	// Disable request.
	////////////////////////////////////////////////////////
	req := model.LbvserverDisable{}
	req.Lbvserver.Name = Name

	err = o.DisableLbvserver(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	LbvserverResp, err := o.GetLbvserver(Name)
	if err != nil {
		t.Error(err)
		return
	}

	if LbvserverResp.Curstate != "OUT OF SERVICE" {
		err = fmt.Errorf("unable to disable resource %+v", LbvserverResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteLbvserver(Name)
	if err != nil {
		t.Error(err)
	}
}

// TestEnableLbvserver creates a resource and enables it.
func TestEnableLbvserver(t *testing.T) {
	var err error
	Name := "Foo"
	Ipaddress := "1.2.3.4"
	State := "DISABLED"
	Port := 80
	Servicetype := "HTTP"
	////////////////////////////////////////////////////////
	reqLbvserveradd := model.LbvserverAdd{}
	reqLbvserveradd.Lbvserver.Name = Name
	reqLbvserveradd.Lbvserver.Ipv46 = Ipaddress
	reqLbvserveradd.Lbvserver.State = State
	reqLbvserveradd.Lbvserver.Port = Port
	reqLbvserveradd.Lbvserver.Servicetype = Servicetype
	////////////////////////////////////////////////////////
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()
	////////////////////////////////////////////////////////
	_, err = createTestLbvserver(reqLbvserveradd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	// Enable request.
	////////////////////////////////////////////////////////
	req := model.LbvserverEnable{}
	req.Lbvserver.Name = Name

	err = o.EnableLbvserver(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	LbvserverResp, err := o.GetLbvserver(Name)
	if err != nil {
		t.Error(err)
		return
	}

	if LbvserverResp.Curstate == "OUT OF SERVICE" {
		err = fmt.Errorf("unable to enable resource %+v", LbvserverResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteLbvserver(Name)
	if err != nil {
		t.Error(err)
	}
}

// TestRenameLbvserver creates a resource and renames it.
func TestRenameLbvserver(t *testing.T) {
	var err error
	Name := "Foo"
	Ipaddress := "1.2.3.4"
	State := "Enabled"
	Port := 80
	Servicetype := "HTTP"
	NewName := "Bar"
	////////////////////////////////////////////////////////
	reqLbvserveradd := model.LbvserverAdd{}
	reqLbvserveradd.Lbvserver.Name = Name
	reqLbvserveradd.Lbvserver.Ipv46 = Ipaddress
	reqLbvserveradd.Lbvserver.State = State
	reqLbvserveradd.Lbvserver.Port = Port
	reqLbvserveradd.Lbvserver.Servicetype = Servicetype
	////////////////////////////////////////////////////////
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()
	////////////////////////////////////////////////////////
	_, err = createTestLbvserver(reqLbvserveradd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	// Rename request.
	////////////////////////////////////////////////////////
	req := model.LbvserverRename{}
	req.Lbvserver.Name = Name
	req.Lbvserver.Newname = NewName

	err = o.RenameLbvserver(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	LbvserverResp, err := o.GetLbvserver(NewName)
	if err != nil {
		t.Error(err)
		return
	}

	if LbvserverResp.Name != NewName {
		err = fmt.Errorf("unable to rename resource %+v", LbvserverResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteLbvserver(NewName)
	if err != nil {
		t.Error(err)
	}
}

// TestUpdateLbvserver creates a resource and updates it.
func TestUpdateLbvserver(t *testing.T) {
	var err error
	Name := "Foo"
	Ipaddress := "1.2.3.4"
	State := "Enabled"
	Port := 80
	Servicetype := "HTTP"
	////////////////////////////////////////////////////////
	reqLbvserveradd := model.LbvserverAdd{}
	reqLbvserveradd.Lbvserver.Name = Name
	reqLbvserveradd.Lbvserver.Ipv46 = Ipaddress
	reqLbvserveradd.Lbvserver.State = State
	reqLbvserveradd.Lbvserver.Port = Port
	reqLbvserveradd.Lbvserver.Servicetype = Servicetype
	////////////////////////////////////////////////////////
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()
	////////////////////////////////////////////////////////
	_, err = createTestLbvserver(reqLbvserveradd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	// Update request.
	////////////////////////////////////////////////////////
	req := model.LbvserverUpdate{}
	req.Lbvserver.Name = Name
	req.Lbvserver.Ipv46 = "1.2.3.5"
	req.Lbvserver.Comment = "Test Lbvserver"

	r, err := o.UpdateLbvserver(req)
	if err != nil {
		t.Error(err)
		return
	}

	if r.Ipv46 != req.Lbvserver.Ipv46 || r.Comment != req.Lbvserver.Comment {
		t.Errorf("failed to update Lbvserver: %+v not equal to %+v", r, req.Lbvserver)
		return
	}

	// Cleanup
	err = o.DeleteLbvserver(Name)
	if err != nil {
		t.Error(err)
	}
}
