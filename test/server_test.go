package test

import (
	"fmt"
	"testing"

	"github.com/tickemaster/nitro-go-sdk/client"
	"github.com/tickemaster/nitro-go-sdk/model"
)

// TestAddDeleteServer creates a resource and deletes it.
func TestAddDeleteServer(t *testing.T) {
	var err error
	Name := "Foo"
	Ipaddress := "99.99.99.99"
	State := "ENABLED"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	_, err = o.GetServer(Name)

	if err == nil {
		err = o.DeleteServer(Name)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	req := model.ServerAdd{}
	req.Server.Name = Name
	req.Server.Ipaddress = Ipaddress
	req.Server.State = State

	expected := model.Server{}
	expected.Name = Name
	expected.Ipaddress = Ipaddress
	expected.State = State

	r, err := o.AddServer(req)
	if err != nil {
		t.Error(err)
		return
	}

	if r.Name != req.Server.Name || r.Ipaddress != req.Server.Ipaddress || r.State != req.Server.State {
		t.Errorf("failed to add resource: %+v not equal to %+v", r, req)
		return
	}

	// Cleanup
	err = o.DeleteServer(Name)
	if err != nil {
		t.Error(err)
	}
}

// TestDisableServer creates a resource and disables it.
func TestDisableServer(t *testing.T) {
	var err error
	Name := "Foo"
	Ipaddress := "99.99.99.99"
	State := "ENABLED"
	Delay := 0
	Grafeful := "Yes"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	_, err = o.GetServer(Name)
	if err == nil {
		err = o.DeleteServer(Name)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServerAdd := model.ServerAdd{}
	reqServerAdd.Server.Name = Name
	reqServerAdd.Server.Ipaddress = Ipaddress
	reqServerAdd.Server.State = State

	_, err = o.AddServer(reqServerAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Disable request.
	req := model.ServerDisable{}
	req.Server.Name = Name
	req.Server.Delay = Delay
	req.Server.Graceful = Grafeful

	err = o.DisableServer(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	serverResp, err := o.GetServer(Name)
	if err != nil {
		t.Error(err)
		return
	}

	if serverResp.State != "DISABLED" {
		err = fmt.Errorf("unable to disable resource %+v", serverResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteServer(Name)
	if err != nil {
		t.Error(err)
	}
}

// TestEnableServer creates a resource and enables it.
func TestEnableServer(t *testing.T) {
	var err error
	Name := "Foo"
	Ipaddress := "99.99.99.99"
	State := "DISABLED"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	_, err = o.GetServer(Name)
	if err == nil {
		err = o.DeleteServer(Name)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServerAdd := model.ServerAdd{}
	reqServerAdd.Server.Name = Name
	reqServerAdd.Server.Ipaddress = Ipaddress
	reqServerAdd.Server.State = State

	_, err = o.AddServer(reqServerAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Disable request.
	req := model.ServerEnable{}
	req.Server.Name = Name

	err = o.EnableServer(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	serverResp, err := o.GetServer(Name)
	if err != nil {
		t.Error(err)
		return
	}

	if serverResp.State != "ENABLED" {
		err = fmt.Errorf("unable to enable resource %+v", serverResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteServer(Name)
	if err != nil {
		t.Error(err)
	}
}

// TestRenameServer creates a resource and renames it.
func TestRenameServer(t *testing.T) {
	var err error
	Name := "Foo"
	NewName := "Bar"
	Ipaddress := "99.99.99.99"
	State := "ENABLED"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	_, err = o.GetServer(Name)
	if err == nil {
		err = o.DeleteServer(Name)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServerAdd := model.ServerAdd{}
	reqServerAdd.Server.Name = Name
	reqServerAdd.Server.Ipaddress = Ipaddress
	reqServerAdd.Server.State = State

	_, err = o.AddServer(reqServerAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Disable request.
	req := model.ServerRename{}
	req.Server.Name = Name
	req.Server.Newname = NewName

	err = o.RenameServer(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	serverResp, err := o.GetServer(NewName)
	if err != nil {
		t.Error(err)
		return
	}

	if serverResp.Name != NewName {
		err = fmt.Errorf("unable to rename resource %+v", serverResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteServer(NewName)
	if err != nil {
		t.Error(err)
	}
}

// TestUpdateServer creates a resource and updates it.
func TestUpdateServer(t *testing.T) {
	var err error
	Name := "Foo"
	Ipaddress := "99.99.99.99"
	State := "ENABLED"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	_, err = o.GetServer(Name)
	if err == nil {
		err = o.DeleteServer(Name)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServerAdd := model.ServerAdd{}
	reqServerAdd.Server.Name = Name
	reqServerAdd.Server.Ipaddress = Ipaddress
	reqServerAdd.Server.State = State

	_, err = o.AddServer(reqServerAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Update request.
	req := model.ServerUpdate{}
	req.Server.Name = Name
	req.Server.Ipaddress = "99.99.99.98"
	req.Server.Comment = "Test Server"

	r, err := o.UpdateServer(req)
	if err != nil {
		t.Error(err)
		return
	}

	if r.Ipaddress != req.Server.Ipaddress || r.Comment != req.Server.Comment {
		t.Errorf("failed to update server: %+v not equal to %+v", r, req.Server)
		return
	}

	// Cleanup
	err = o.DeleteServer(Name)
	if err != nil {
		t.Error(err)
	}
}
