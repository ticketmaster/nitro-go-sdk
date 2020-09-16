package test

import (
	"fmt"
	"testing"

	"github.com/tickemaster/nitro-go-sdk/client"
	"github.com/tickemaster/nitro-go-sdk/model"
)

// NOTE: You can test memberport (shared port across all backends), but you will need to also set autoscale.

// TestAddDeleteService creates a resource and deletes it.
func TestAddDeleteService(t *testing.T) {
	var err error
	Name := "Foo"
	State := "DISABLED"
	Servicetype := "HTTP"
	Port := 80
	IP := "1.2.2.3"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()
	///////////////////////////////////////////////////////////
	// Add server.
	///////////////////////////////////////////////////////////
	reqServeradd := model.ServerAdd{}
	reqServeradd.Server.State = State
	reqServeradd.Server.Ipaddress = IP
	reqServeradd.Server.Name = IP
	_, err = createTestServer(reqServeradd, o)
	if err != nil {
		t.Error(err)
	}
	///////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	_, err = o.GetService(Name)
	if err == nil {
		err = o.DeleteService(Name)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	req := model.ServiceAdd{}
	req.Service.Name = Name
	req.Service.State = State
	req.Service.Servicetype = Servicetype
	req.Service.Port = Port
	req.Service.Ipaddress = IP
	req.Service.IP = IP

	r, err := o.AddService(req)
	if err != nil {
		t.Error(err)
		return
	}

	if r.Name != req.Service.Name || r.Svrstate != "OUT OF SERVICE" {
		t.Errorf("failed to add resource: %+v not equal to %+v", r, req)
		return
	}

	// Cleanup
	err = o.DeleteService(Name)
	if err != nil {
		t.Error(err)
	}
	err = o.DeleteServer(IP)
	if err != nil {
		t.Error(err)
	}
}

// TestDisableService creates a resource and disables it.
func TestDisableService(t *testing.T) {
	var err error
	Name := "Foo"
	State := "DISABLED"
	Servicetype := "HTTP"
	Port := 80
	IP := "1.2.2.3"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()
	///////////////////////////////////////////////////////////
	// Add server.
	///////////////////////////////////////////////////////////
	reqServeradd := model.ServerAdd{}
	reqServeradd.Server.State = State
	reqServeradd.Server.Ipaddress = IP
	reqServeradd.Server.Name = IP
	_, err = createTestServer(reqServeradd, o)
	if err != nil {
		t.Error(err)

	}
	///////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	_, err = o.GetService(Name)
	if err == nil {
		err = o.DeleteService(Name)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServiceAdd := model.ServiceAdd{}
	reqServiceAdd.Service.Name = Name
	reqServiceAdd.Service.State = State
	reqServiceAdd.Service.Servicetype = Servicetype
	reqServiceAdd.Service.Port = Port
	reqServiceAdd.Service.IP = IP
	reqServiceAdd.Service.Ipaddress = IP

	_, err = o.AddService(reqServiceAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Disable request.
	req := model.ServiceDisable{}
	req.Service.Name = Name

	err = o.DisableService(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	ServiceResp, err := o.GetService(Name)
	if err != nil {
		t.Error(err)
		return
	}

	if ServiceResp.Svrstate != "OUT OF SERVICE" {
		err = fmt.Errorf("unable to disable resource %+v", ServiceResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteService(Name)
	if err != nil {
		t.Error(err)
	}

	err = o.DeleteServer(IP)
	if err != nil {
		t.Error(err)
	}
}

// TestEnableService creates a resource and enables it.
func TestEnableService(t *testing.T) {
	var err error
	Name := "Foo"
	State := "DISABLED"
	Servicetype := "HTTP"
	Port := 80
	IP := "1.2.2.3"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()
	///////////////////////////////////////////////////////////
	// Add server.
	///////////////////////////////////////////////////////////
	reqServeradd := model.ServerAdd{}
	reqServeradd.Server.State = State
	reqServeradd.Server.Ipaddress = IP
	reqServeradd.Server.Name = IP
	_, err = createTestServer(reqServeradd, o)
	if err != nil {
		t.Error(err)

	}
	///////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	_, err = o.GetService(Name)
	if err == nil {
		err = o.DeleteService(Name)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServiceAdd := model.ServiceAdd{}
	reqServiceAdd.Service.Name = Name
	reqServiceAdd.Service.State = State
	reqServiceAdd.Service.Servicetype = Servicetype
	reqServiceAdd.Service.Port = Port
	reqServiceAdd.Service.IP = IP
	reqServiceAdd.Service.Ipaddress = IP

	_, err = o.AddService(reqServiceAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Disable request.
	req := model.ServiceEnable{}
	req.Service.Name = Name

	err = o.EnableService(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	ServiceResp, err := o.GetService(Name)
	if err != nil {
		t.Error(err)
		return
	}

	if ServiceResp.Svrstate == "OUT OF SERVICE" {
		err = fmt.Errorf("unable to enable resource %+v", ServiceResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteService(Name)
	if err != nil {
		t.Error(err)
	}
	err = o.DeleteServer(IP)
	if err != nil {
		t.Error(err)
	}
}

// TestRenameService creates a resource and renames it.
func TestRenameService(t *testing.T) {
	var err error
	Name := "Foo"
	State := "DISABLED"
	Servicetype := "HTTP"
	NewName := "Bar"
	Port := 80
	IP := "1.2.2.3"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()
	///////////////////////////////////////////////////////////
	// Add server.
	///////////////////////////////////////////////////////////
	reqServeradd := model.ServerAdd{}
	reqServeradd.Server.State = State
	reqServeradd.Server.Ipaddress = IP
	reqServeradd.Server.Name = IP
	_, err = createTestServer(reqServeradd, o)
	if err != nil {
		t.Error(err)

	}
	///////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	_, err = o.GetService(Name)
	if err == nil {
		err = o.DeleteService(Name)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServiceAdd := model.ServiceAdd{}
	reqServiceAdd.Service.Name = Name
	reqServiceAdd.Service.State = State
	reqServiceAdd.Service.Servicetype = Servicetype
	reqServiceAdd.Service.Port = Port
	reqServiceAdd.Service.IP = IP
	reqServiceAdd.Service.Ipaddress = IP

	_, err = o.AddService(reqServiceAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Disable request.
	req := model.ServiceRename{}
	req.Service.Name = Name
	req.Service.Newname = NewName

	err = o.RenameService(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	ServiceResp, err := o.GetService(NewName)
	if err != nil {
		t.Error(err)
		return
	}

	if ServiceResp.Name != NewName {
		err = fmt.Errorf("unable to rename resource %+v", ServiceResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteService(NewName)
	if err != nil {
		t.Error(err)
	}
	err = o.DeleteServer(IP)
	if err != nil {
		t.Error(err)
	}
}

// TestUpdateService creates a resource and updates it.
func TestUpdateService(t *testing.T) {
	var err error
	Name := "Foo"
	State := "DISABLED"
	Servicetype := "HTTP"
	Port := 80
	IP := "1.2.2.3"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()
	///////////////////////////////////////////////////////////
	// Add server.
	///////////////////////////////////////////////////////////
	reqServeradd := model.ServerAdd{}
	reqServeradd.Server.State = State
	reqServeradd.Server.Ipaddress = IP
	reqServeradd.Server.Name = IP
	_, err = createTestServer(reqServeradd, o)
	if err != nil {
		t.Error(err)

	}
	///////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	_, err = o.GetService(Name)
	if err == nil {
		err = o.DeleteService(Name)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServiceAdd := model.ServiceAdd{}
	reqServiceAdd.Service.Name = Name
	reqServiceAdd.Service.State = State
	reqServiceAdd.Service.Servicetype = Servicetype
	reqServiceAdd.Service.Port = Port
	reqServiceAdd.Service.IP = IP
	reqServiceAdd.Service.Ipaddress = IP

	_, err = o.AddService(reqServiceAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Update request.
	req := model.ServiceUpdate{}
	req.Service.Name = Name
	req.Service.Comment = "Test Service"
	req.Service.Ipaddress = IP

	r, err := o.UpdateService(req)
	if err != nil {
		t.Error(err)
		return
	}

	if r.Comment != req.Service.Comment {
		t.Errorf("failed to update Service: %+v not equal to %+v", r, req.Service)
		return
	}

	// Cleanup
	err = o.DeleteService(Name)
	if err != nil {
		t.Error(err)
	}
	err = o.DeleteServer(IP)
	if err != nil {
		t.Error(err)
	}
}
