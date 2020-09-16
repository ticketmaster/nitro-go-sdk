package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tickemaster/nitro-go-sdk/client"
)

// TestGetHanode returns Hanode
func TestGetHanode(t *testing.T) {
	var err error

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	r, err := o.GetHanode(testNetscaler)
	if err != nil {
		t.Error(err)
		return
	}

	mR := make(map[string]interface{})

	iR, err := json.Marshal(r)
	if err != nil {
		t.Error(err)
		return
	}

	err = json.Unmarshal(iR, &mR)
	if err != nil {
		t.Error(err)
		return
	}

	for k, v := range mR {
		val := fmt.Sprint(v)
		// ifaces is optional.
		// haheartbeatifaces is optional.
		// pfifaces is optional.
		// hamonifaces is optional.
		if val == "" && k != "ifaces" && k != "haheartbeatifaces" && k != "pfifaces" && k != "hamonifaces" {
			err := fmt.Errorf("expected field did not contain a value: %v %v", k, v)
			t.Error(err)
			return
		}
	}

	return

}
