package test

import (
	"testing"

	"github.com/tickemaster/nitro-go-sdk/client"
	"github.com/tickemaster/nitro-go-sdk/model"
)

// TestAddDeleteSSLCipherGroup creates a cipher group and deletes it.
func TestAddDeleteSslciphergroup(t *testing.T) {
	var err error
	ciphergroupname := "Foo"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure cipher group does not already exist.
	_, err = o.GetSslciphergroup(ciphergroupname)

	if err == nil {
		err = o.DeleteSslciphergroup(ciphergroupname)
		if err != nil {
			t.Error(err)
		}
	}

	// Generate request.
	req := model.SslcipherAdd{}
	req.Sslcipher.Ciphergroupname = ciphergroupname
	req.Sslcipher.Ciphgrpalias = ciphergroupname

	expected := model.Sslcipher{}
	expected.Ciphergroupname = ciphergroupname
	expected.Ciphername = ciphergroupname
	expected.Cipherpriority = "0"
	expected.Description = "User Defined Cipher Group"

	r, err := o.AddSslciphergroup(req)
	if err != nil {
		t.Error(err)
		return
	}

	if r != expected {
		t.Errorf("failed to add cipher group: %+v not equal to %+v", r, expected)
		return
	}

	// Cleanup
	err = o.DeleteSslciphergroup(ciphergroupname)
	if err != nil {
		t.Error(err)
	}
}
