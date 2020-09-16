package test

import (
	"fmt"
	"testing"

	"github.com/tickemaster/nitro-go-sdk/client"
	"github.com/tickemaster/nitro-go-sdk/model"
)

// TestAddDeleteServicegroupServicegroupmemberBinding creates a resource and deletes it.
func TestAddDeleteServicegroupServicegroupmemberBinding(t *testing.T) {
	var err error
	Ipaddress := "1.2.3.4"
	Serveripaddress := "1.2.3.5"
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
	reqServeradd := model.ServerAdd{}
	reqServeradd.Server.State = State
	reqServeradd.Server.Ipaddress = Serveripaddress
	reqServeradd.Server.Name = Serveripaddress
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
	_, err = createTestServer(reqServeradd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	////////////////////////////////////////////////////////
	_, err = o.GetServicegroupServicegroupmemberBindingByIP(Servicegroupname, Serveripaddress, Port)

	if err == nil {
		err = o.DeleteServicegroupServicegroupmemberBindingByIP(Servicegroupname, Serveripaddress, Port)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	req := model.ServicegroupServicegroupmemberBindingAdd{}
	req.ServicegroupServicegroupmemberBinding.Servicegroupname = Servicegroupname
	req.ServicegroupServicegroupmemberBinding.IP = Serveripaddress
	req.ServicegroupServicegroupmemberBinding.Port = Port
	req.ServicegroupServicegroupmemberBinding.State = State

	expected := model.ServicegroupServicegroupmemberBinding{}
	expected.Servicegroupname = Servicegroupname
	expected.IP = Serveripaddress
	expected.Port = Port
	expected.Servername = Serveripaddress

	r, err := o.AddServicegroupServicegroupmemberBinding(req)
	if err != nil {
		t.Error(err)
		return
	}

	if expected.Servicegroupname != r.Servicegroupname || expected.IP != r.IP || expected.Port != expected.Port {
		t.Errorf("failed to add resource: %+v not equal to %+v", r, expected)
		return
	}

	// Cleanup
	err = o.DeleteServicegroupServicegroupmemberBindingByIP(Servicegroupname, Serveripaddress, Port)
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
	err = o.DeleteServer(Serveripaddress)
	if err != nil {
		t.Error(err)
	}

}

// TestDisableServicegroupServicegroupmemberBinding creates a resource and disables it.
func TestDisableServicegroupServicegroupmemberBinding(t *testing.T) {
	var err error
	Ipaddress := "1.2.3.4"
	Serveripaddress := "1.2.3.5"
	Name := "Foo"
	Port := 80
	Servicetype := "HTTP"
	Servicegroupname := "Foo_Pool"
	State := "ENABLED"
	Graceful := "Yes"
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
	reqServeradd := model.ServerAdd{}
	reqServeradd.Server.State = State
	reqServeradd.Server.Ipaddress = Serveripaddress
	reqServeradd.Server.Name = Serveripaddress
	////////////////////////////////////////////////////////
	reqServicegroupServicegroupmemberBinding := model.ServicegroupServicegroupmemberBindingAdd{}
	reqServicegroupServicegroupmemberBinding.ServicegroupServicegroupmemberBinding.Servicegroupname = Servicegroupname
	reqServicegroupServicegroupmemberBinding.ServicegroupServicegroupmemberBinding.IP = Serveripaddress
	reqServicegroupServicegroupmemberBinding.ServicegroupServicegroupmemberBinding.Port = Port
	reqServicegroupServicegroupmemberBinding.ServicegroupServicegroupmemberBinding.State = State
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
	_, err = createTestServer(reqServeradd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	_, err = createTestServicegroupServicegroupmemberBinding(reqServicegroupServicegroupmemberBinding, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	// Disable request.
	////////////////////////////////////////////////////////
	req := model.ServicegroupDisable{}
	req.Servicegroup.Servicegroupname = Servicegroupname
	req.Servicegroup.Servername = Serveripaddress
	req.Servicegroup.Graceful = Graceful
	req.Servicegroup.Port = Port

	err = o.DisableServicegroupServicegroupmemberBinding(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get response.
	resp, err := o.GetServicegroupServicegroupmemberBindingByIP(Servicegroupname, Serveripaddress, Port)
	if err != nil {
		t.Error(err)
		return
	}

	if resp.State != "DISABLED" {
		err = fmt.Errorf("unable to disable resource %+v", resp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteServicegroupServicegroupmemberBindingByIP(Servicegroupname, Serveripaddress, Port)
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
	err = o.DeleteServer(Serveripaddress)
	if err != nil {
		t.Error(err)
	}
}

func TestEnableServicegroupServicegroupmemberBinding(t *testing.T) {
	var err error
	Ipaddress := "1.2.3.4"
	Serveripaddress := "1.2.3.5"
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
	reqServeradd := model.ServerAdd{}
	reqServeradd.Server.State = State
	reqServeradd.Server.Ipaddress = Serveripaddress
	reqServeradd.Server.Name = Serveripaddress
	////////////////////////////////////////////////////////
	reqServicegroupServicegroupmemberBinding := model.ServicegroupServicegroupmemberBindingAdd{}
	reqServicegroupServicegroupmemberBinding.ServicegroupServicegroupmemberBinding.Servicegroupname = Servicegroupname
	reqServicegroupServicegroupmemberBinding.ServicegroupServicegroupmemberBinding.IP = Serveripaddress
	reqServicegroupServicegroupmemberBinding.ServicegroupServicegroupmemberBinding.Port = Port
	reqServicegroupServicegroupmemberBinding.ServicegroupServicegroupmemberBinding.State = State
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
	_, err = createTestServer(reqServeradd, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	_, err = createTestServicegroupServicegroupmemberBinding(reqServicegroupServicegroupmemberBinding, o)
	if err != nil {
		t.Error(err)
		return
	}
	////////////////////////////////////////////////////////
	// Enable request.
	////////////////////////////////////////////////////////
	req := model.ServicegroupEnable{}
	req.Servicegroup.Servicegroupname = Servicegroupname
	req.Servicegroup.Servername = Serveripaddress
	req.Servicegroup.Port = Port

	err = o.EnableServicegroupServicegroupmemberBinding(req)
	if err != nil {
		t.Error(err)
		return
	}

	// Get response.
	resp, err := o.GetServicegroupServicegroupmemberBindingByIP(Servicegroupname, Serveripaddress, Port)
	if err != nil {
		t.Error(err)
		return
	}

	if resp.State == "DISABLED" {
		err = fmt.Errorf("unable to disable resource %+v", resp)
		t.Error(err)
		return
	}

	// Cleanup
	err = o.DeleteServicegroupServicegroupmemberBindingByIP(Servicegroupname, Serveripaddress, Port)
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
	err = o.DeleteServer(Serveripaddress)
	if err != nil {
		t.Error(err)
	}
}
