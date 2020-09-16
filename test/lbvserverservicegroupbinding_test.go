package test

import (
	"fmt"
	"testing"

	"github.com/tickemaster/nitro-go-sdk/client"
	"github.com/tickemaster/nitro-go-sdk/model"
)

// TestAddDeleteLbvserverServicegroupBinding creates a resource and deletes it.
func TestAddDeleteLbvserverServicegroupBinding(t *testing.T) {
	var err error
	Ipaddress := "1.2.3.4"
	Name := "Foo"
	Port := 80
	Servicetype := "HTTP"
	Servicegroupname := "Foo_Pool"
	State := "ENABLED"
	////////////////////////////////////////////////////////
	reqLbvserveradd := model.LbvserverAdd{}
	reqLbvserveradd.Lbvserver.Name = Name
	reqLbvserveradd.Lbvserver.Ipv46 = Ipaddress
	reqLbvserveradd.Lbvserver.State = State
	reqLbvserveradd.Lbvserver.Port = Port
	reqLbvserveradd.Lbvserver.Servicetype = Servicetype
	////////////////////////////////////////////////////////
	reqServicegroupadd := model.ServicegroupAdd{}
	reqServicegroupadd.Servicegroup.Servicegroupname = Servicegroupname
	reqServicegroupadd.Servicegroup.State = State
	reqServicegroupadd.Servicegroup.Servicetype = Servicetype
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
	_, err = createTestServicegroup(reqServicegroupadd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	////////////////////////////////////////////////////////
	_, err = o.GetLbvserverServicegroupBinding(Name, Servicegroupname)

	if err == nil {
		err = o.DeleteLbvserverServicegroupBinding(Name, Servicegroupname)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	req := model.LbvserverServicegroupBindingAdd{}
	req.LbvserverServicegroupBinding.Name = Name
	req.LbvserverServicegroupBinding.Servicegroupname = Servicegroupname

	expected := model.LbvserverServicegroupBinding{}
	expected.Name = Name
	expected.Servicegroupname = Servicegroupname
	expected.Servicename = Servicegroupname

	r, err := o.AddLbvserverServicegroupBinding(req)
	if err != nil {
		t.Error(err)
		return
	}

	if r != expected {
		t.Errorf("failed to add resource: %+v not equal to %+v", r, expected)
		return
	}

	// Cleanup
	err = o.DeleteLbvserverServicegroupBinding(Name, Servicegroupname)
	if err != nil {
		t.Error(err)
	}
	err = o.DeleteLbvserver(Name)
	if err != nil {
		t.Error(err)
	}
	err = o.DeleteServicegroup(Servicegroupname)
	if err != nil {
		t.Error(err)
	}
}

// TestDisableLbvserverServicegroupBinding creates a resource and disables it.
func TestDisableLbvserverServicegroupBinding(t *testing.T) {
	var err error
	Ipaddress := "1.2.3.4"
	Name := "Foo"
	Port := 80
	Servicetype := "HTTP"
	Servicegroupname := "Foo_Pool"
	State := "ENABLED"
	Graceful := "YES"
	////////////////////////////////////////////////////////
	reqLbvserveradd := model.LbvserverAdd{}
	reqLbvserveradd.Lbvserver.Name = Name
	reqLbvserveradd.Lbvserver.Ipv46 = Ipaddress
	reqLbvserveradd.Lbvserver.State = State
	reqLbvserveradd.Lbvserver.Port = Port
	reqLbvserveradd.Lbvserver.Servicetype = Servicetype
	////////////////////////////////////////////////////////
	reqServicegroupadd := model.ServicegroupAdd{}
	reqServicegroupadd.Servicegroup.Servicegroupname = Servicegroupname
	reqServicegroupadd.Servicegroup.State = State
	reqServicegroupadd.Servicegroup.Servicetype = Servicetype
	////////////////////////////////////////////////////////
	reqLbvserverServicegroupBindingadd := model.LbvserverServicegroupBindingAdd{}
	reqLbvserverServicegroupBindingadd.LbvserverServicegroupBinding.Name = Name
	reqLbvserverServicegroupBindingadd.LbvserverServicegroupBinding.Servicegroupname = Servicegroupname
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
	_, err = createTestServicegroup(reqServicegroupadd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	_, err = createTestLbvserverServicegroupBinding(reqLbvserverServicegroupBindingadd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	// Disable request.
	////////////////////////////////////////////////////////
	req := model.ServicegroupDisable{}
	req.Servicegroup.Servicegroupname = Servicegroupname
	req.Servicegroup.Graceful = Graceful

	err = o.DisableLbvserverServicegroupBinding(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get response.
	resp, err := o.GetServicegroup(Servicegroupname)
	if err != nil {
		t.Error(err)
		return
	}

	if resp.State != "DISABLED" {
		err = fmt.Errorf("unable to disable resource %+v", Servicegroupname)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteLbvserverServicegroupBinding(Name, Servicegroupname)
	if err != nil {
		t.Error(err)
	}
	err = o.DeleteLbvserver(Name)
	if err != nil {
		t.Error(err)
	}
	err = o.DeleteServicegroup(Servicegroupname)
	if err != nil {
		t.Error(err)
	}
}

// TestDisableLbvserverServicegroupBinding creates a resource and disables it.
func TestEnableLbvserverServicegroupBinding(t *testing.T) {
	var err error
	Ipaddress := "1.2.3.4"
	Name := "Foo"
	Port := 80
	Servicetype := "HTTP"
	Servicegroupname := "Foo_Pool"
	State := "DISABLED"
	////////////////////////////////////////////////////////
	reqLbvserveradd := model.LbvserverAdd{}
	reqLbvserveradd.Lbvserver.Name = Name
	reqLbvserveradd.Lbvserver.Ipv46 = Ipaddress
	reqLbvserveradd.Lbvserver.State = State
	reqLbvserveradd.Lbvserver.Port = Port
	reqLbvserveradd.Lbvserver.Servicetype = Servicetype
	////////////////////////////////////////////////////////
	reqServicegroupadd := model.ServicegroupAdd{}
	reqServicegroupadd.Servicegroup.Servicegroupname = Servicegroupname
	reqServicegroupadd.Servicegroup.State = State
	reqServicegroupadd.Servicegroup.Servicetype = Servicetype
	////////////////////////////////////////////////////////
	reqLbvserverServicegroupBindingadd := model.LbvserverServicegroupBindingAdd{}
	reqLbvserverServicegroupBindingadd.LbvserverServicegroupBinding.Name = Name
	reqLbvserverServicegroupBindingadd.LbvserverServicegroupBinding.Servicegroupname = Servicegroupname
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
	_, err = createTestServicegroup(reqServicegroupadd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	_, err = createTestLbvserverServicegroupBinding(reqLbvserverServicegroupBindingadd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	// Enable request.
	////////////////////////////////////////////////////////
	req := model.ServicegroupEnable{}
	req.Servicegroup.Servicegroupname = Servicegroupname

	err = o.EnableLbvserverServicegroupBinding(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get response.
	resp, err := o.GetServicegroup(Servicegroupname)
	if err != nil {
		t.Error(err)
		return
	}

	if resp.State == "DISABLED" {
		err = fmt.Errorf("unable to disable resource %+v", Servicegroupname)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteLbvserverServicegroupBinding(Name, Servicegroupname)
	if err != nil {
		t.Error(err)
	}
	err = o.DeleteLbvserver(Name)
	if err != nil {
		t.Error(err)
	}
	err = o.DeleteServicegroup(Servicegroupname)
	if err != nil {
		t.Error(err)
	}
}
