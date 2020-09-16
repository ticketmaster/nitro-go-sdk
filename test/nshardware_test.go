package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/tickemaster/nitro-go-sdk/client"
)

// TestGetNshwardware returns nshardware
func TestGetNshwardware(t *testing.T) {
	var err error

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	r, err := o.GetNshardware()
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
		if val == "" {
			err := fmt.Errorf("expected field did not contain a value: %v %v", k, v)
			t.Error(err)
			return
		}
	}

	return

}
