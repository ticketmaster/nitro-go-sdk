package test

import (
	"fmt"
	"testing"

	"github.com/ticketmaster/nitro-go-sdk/client"
	"github.com/ticketmaster/nitro-go-sdk/model"
)

// NOTE: You can test memberport (shared port across all backends), but you will need to also set autoscale.

// TestAddDeleteServicegroup creates a resource and deletes it.
func TestAddDeleteServicegroup(t *testing.T) {
	var err error
	Servicegroupname := "Foo"
	State := "DISABLED"
	Servicetype := "HTTP"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	_, err = o.GetServicegroup(Servicegroupname)

	if err == nil {
		err = o.DeleteServicegroup(Servicegroupname)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	req := model.ServicegroupAdd{}
	req.Servicegroup.Servicegroupname = Servicegroupname
	req.Servicegroup.State = State
	req.Servicegroup.Servicetype = Servicetype

	r, err := o.AddServicegroup(req)
	if err != nil {
		t.Error(err)
		return
	}

	if r.Servicegroupname != req.Servicegroup.Servicegroupname || r.State != "DISABLED" {
		t.Errorf("failed to add resource: %+v not equal to %+v", r, req)
		return
	}

	// Cleanup
	err = o.DeleteServicegroup(Servicegroupname)
	if err != nil {
		t.Error(err)
	}
}

// TestDisableServicegroup creates a resource and disables it.
func TestDisableServicegroup(t *testing.T) {
	var err error
	Servicegroupname := "Foo"
	State := "DISABLED"
	Servicetype := "HTTP"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	_, err = o.GetServicegroup(Servicegroupname)
	if err == nil {
		err = o.DeleteServicegroup(Servicegroupname)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServicegroupAdd := model.ServicegroupAdd{}
	reqServicegroupAdd.Servicegroup.Servicegroupname = Servicegroupname
	reqServicegroupAdd.Servicegroup.State = State
	reqServicegroupAdd.Servicegroup.Servicetype = Servicetype

	_, err = o.AddServicegroup(reqServicegroupAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Disable request.
	req := model.ServicegroupDisable{}
	req.Servicegroup.Servicegroupname = Servicegroupname

	err = o.DisableServicegroup(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	ServicegroupResp, err := o.GetServicegroup(Servicegroupname)
	if err != nil {
		t.Error(err)
		return
	}

	if ServicegroupResp.State != "DISABLED" {
		err = fmt.Errorf("unable to disable resource %+v", ServicegroupResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteServicegroup(Servicegroupname)
	if err != nil {
		t.Error(err)
	}
}

// TestEnableServicegroup creates a resource and enables it.
func TestEnableServicegroup(t *testing.T) {
	var err error
	Servicegroupname := "Foo"
	State := "DISABLED"
	Servicetype := "HTTP"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	_, err = o.GetServicegroup(Servicegroupname)
	if err == nil {
		err = o.DeleteServicegroup(Servicegroupname)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServicegroupAdd := model.ServicegroupAdd{}
	reqServicegroupAdd.Servicegroup.Servicegroupname = Servicegroupname
	reqServicegroupAdd.Servicegroup.State = State
	reqServicegroupAdd.Servicegroup.Servicetype = Servicetype

	_, err = o.AddServicegroup(reqServicegroupAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Disable request.
	req := model.ServicegroupEnable{}
	req.Servicegroup.Servicegroupname = Servicegroupname

	err = o.EnableServicegroup(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	ServicegroupResp, err := o.GetServicegroup(Servicegroupname)
	if err != nil {
		t.Error(err)
		return
	}

	if ServicegroupResp.State == "DISABLED" {
		err = fmt.Errorf("unable to enable resource %+v", ServicegroupResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteServicegroup(Servicegroupname)
	if err != nil {
		t.Error(err)
	}
}

// TestRenameServicegroup creates a resource and renames it.
func TestRenameServicegroup(t *testing.T) {
	var err error
	Servicegroupname := "Foo"
	State := "DISABLED"
	Servicetype := "HTTP"
	NewServicegroupname := "Bar"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	_, err = o.GetServicegroup(Servicegroupname)
	if err == nil {
		err = o.DeleteServicegroup(Servicegroupname)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServicegroupAdd := model.ServicegroupAdd{}
	reqServicegroupAdd.Servicegroup.Servicegroupname = Servicegroupname
	reqServicegroupAdd.Servicegroup.State = State
	reqServicegroupAdd.Servicegroup.Servicetype = Servicetype

	_, err = o.AddServicegroup(reqServicegroupAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Disable request.
	req := model.ServicegroupRename{}
	req.Servicegroup.Servicegroupname = Servicegroupname
	req.Servicegroup.Newname = NewServicegroupname

	err = o.RenameServicegroup(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get request.
	ServicegroupResp, err := o.GetServicegroup(NewServicegroupname)
	if err != nil {
		t.Error(err)
		return
	}

	if ServicegroupResp.Servicegroupname != NewServicegroupname {
		err = fmt.Errorf("unable to rename resource %+v", ServicegroupResp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteServicegroup(NewServicegroupname)
	if err != nil {
		t.Error(err)
	}
}

// TestUpdateServicegroup creates a resource and updates it.
func TestUpdateServicegroup(t *testing.T) {
	var err error
	Servicegroupname := "Foo"
	State := "DISABLED"
	Servicetype := "HTTP"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure resource does not already exist.
	_, err = o.GetServicegroup(Servicegroupname)
	if err == nil {
		err = o.DeleteServicegroup(Servicegroupname)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	reqServicegroupAdd := model.ServicegroupAdd{}
	reqServicegroupAdd.Servicegroup.Servicegroupname = Servicegroupname
	reqServicegroupAdd.Servicegroup.State = State
	reqServicegroupAdd.Servicegroup.Servicetype = Servicetype

	_, err = o.AddServicegroup(reqServicegroupAdd)
	if err != nil {
		t.Error(err)
		return
	}

	// Update request.
	req := model.ServicegroupUpdate{}
	req.Servicegroup.Servicegroupname = Servicegroupname
	req.Servicegroup.Comment = "Test Servicegroup"

	r, err := o.UpdateServicegroup(req)
	if err != nil {
		t.Error(err)
		return
	}

	if r.Comment != req.Servicegroup.Comment {
		t.Errorf("failed to update Servicegroup: %+v not equal to %+v", r, req.Servicegroup)
		return
	}

	// Cleanup
	err = o.DeleteServicegroup(Servicegroupname)
	if err != nil {
		t.Error(err)
	}
}
