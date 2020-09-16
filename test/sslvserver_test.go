package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/tickemaster/nitro-go-sdk/client"
	"github.com/tickemaster/nitro-go-sdk/model"
)

func TestAddDeleteSslvserver(t *testing.T) {
	var err error
	Name := "Foo"
	Ipaddress := "1.2.3.4"
	State := "Enabled"
	Port := 443
	Servicetype := "SSL_BRIDGE"
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

	// Cleanup
	err = o.DeleteLbvserver(Name)
	if err != nil {
		t.Error(err)
	}
}

func TestAddUpdateSslvserver(t *testing.T) {
	var err error
	Name := "Foo"
	Ipaddress := "1.2.3.4"
	State := "Enabled"
	Port := 443
	Servicetype := "SSL"
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
	req := model.SslvserverUpdate{}
	req.Sslvserver.Vservername = Name
	req.Sslvserver.TLS1 = "DISABLED"
	req.Sslvserver.TLS11 = "DISABLED"

	r := model.Sslvserver{}
	r, err = o.UpdateSslvserver(req)
	retry := 3

	for err != nil {
		time.Sleep(100 * time.Millisecond)
		retry = retry - 1

		if retry == 0 {
			t.Error(err)
			return
		}
	}

	if r.TLS11 != req.Sslvserver.TLS1 || r.TLS12 != req.Sslvserver.TLS12 {
		err = fmt.Errorf("sslvserver was not updated: [expected] %+v [received] %+v", req, r)
	}

	// Cleanup
	err = o.DeleteLbvserver(Name)
	if err != nil {
		t.Error(err)
	}
}
