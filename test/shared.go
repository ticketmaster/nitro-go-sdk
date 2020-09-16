package test

import (
	"encoding/json"
	"log"

	"github.com/ticketmaster/nitro-go-sdk/client"
	"github.com/ticketmaster/nitro-go-sdk/model"
	"github.com/tidwall/pretty"
)

// connectToTestNetscaler connects to LB target.
func connectToTestNetscaler() (s *client.Session, err error) {
	s = new(client.Session)
	s, err = client.New(testNetscaler, testUser, testPassword)
	return
}

func createTestLbvserver(req model.LbvserverAdd, o *client.Netscaler) (r model.Lbvserver, err error) {
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	////////////////////////////////////////////////////////
	_, err = o.GetLbvserver(req.Lbvserver.Name)
	if err == nil {
		err = o.DeleteLbvserver(req.Lbvserver.Name)
		if err != nil {
			return
		}
	}
	////////////////////////////////////////////////////////
	// Generate request.
	////////////////////////////////////////////////////////
	return o.AddLbvserver(req)
}

func createTestServer(req model.ServerAdd, o *client.Netscaler) (r model.Server, err error) {
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	////////////////////////////////////////////////////////
	_, err = o.GetServer(req.Server.Name)
	if err == nil {
		err = o.DeleteServer(req.Server.Name)
		if err != nil {
		}
	}
	////////////////////////////////////////////////////////
	// Generate request.
	////////////////////////////////////////////////////////
	return o.AddServer(req)
}

func createTestServicegroup(req model.ServicegroupAdd, o *client.Netscaler) (r model.Servicegroup, err error) {
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	////////////////////////////////////////////////////////
	_, err = o.GetServicegroup(req.Servicegroup.Servicegroupname)
	if err == nil {
		err = o.DeleteServicegroup(req.Servicegroup.Servicegroupname)
		if err != nil {
		}
	}
	////////////////////////////////////////////////////////
	// Generate request.
	////////////////////////////////////////////////////////
	return o.AddServicegroup(req)
}

func createTestLbvserverServicegroupBinding(req model.LbvserverServicegroupBindingAdd, o *client.Netscaler) (r model.LbvserverServicegroupBinding, err error) {
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	////////////////////////////////////////////////////////
	_, err = o.GetLbvserverServicegroupBinding(req.LbvserverServicegroupBinding.Name, req.LbvserverServicegroupBinding.Servicegroupname)
	if err == nil {
		err = o.DeleteLbvserverServicegroupBinding(req.LbvserverServicegroupBinding.Name, req.LbvserverServicegroupBinding.Servicegroupname)
		if err != nil {
		}
	}
	////////////////////////////////////////////////////////
	// Generate request.
	////////////////////////////////////////////////////////
	return o.AddLbvserverServicegroupBinding(req)
}

func createTestServicegroupServicegroupmemberBinding(req model.ServicegroupServicegroupmemberBindingAdd, o *client.Netscaler) (r model.ServicegroupServicegroupmemberBinding, err error) {
	if err != nil {
		return
	}
	////////////////////////////////////////////////////////
	// Ensure resource does not already exist.
	////////////////////////////////////////////////////////
	_, err = o.GetServicegroupServicegroupmemberBindingByIP(req.ServicegroupServicegroupmemberBinding.Servicegroupname, req.ServicegroupServicegroupmemberBinding.IP, req.ServicegroupServicegroupmemberBinding.Port)
	if err == nil {
		err = o.DeleteServicegroupServicegroupmemberBindingByIP(req.ServicegroupServicegroupmemberBinding.Servicegroupname, req.ServicegroupServicegroupmemberBinding.IP, req.ServicegroupServicegroupmemberBinding.Port)
		if err != nil {
		}
	}
	////////////////////////////////////////////////////////
	// Generate request.
	////////////////////////////////////////////////////////
	return o.AddServicegroupServicegroupmemberBinding(req)
}

func toPrettyJSON(p interface{}) []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		log.Println(err.Error())
	}
	return pretty.Pretty(bytes)
}
