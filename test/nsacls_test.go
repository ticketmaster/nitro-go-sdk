package test

import (
	"testing"

	"github.com/tickemaster/nitro-go-sdk/client"
)

func TestClearNsacls(t *testing.T) {
	///////////////////////////////////////////////////////////////////////
	var err error
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}
	defer o.Session.Logout()
	///////////////////////////////////////////////////////////////////////
	err = o.ClearNsacls()
	if err != nil {
		t.Error(err)
		return
	}
	return
}

func TestApplyNsacls(t *testing.T) {
	///////////////////////////////////////////////////////////////////////
	var err error
	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}
	defer o.Session.Logout()
	///////////////////////////////////////////////////////////////////////
	err = o.ApplyNsacls()
	if err != nil {
		t.Error(err)
		return
	}
	return
}
