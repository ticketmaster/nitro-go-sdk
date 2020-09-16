package test

import (
	"testing"

	"github.com/ticketmaster/nitro-go-sdk/client"
	"github.com/ticketmaster/nitro-go-sdk/model"
)

// TestAddDeleteSslcipherSslciphersuiteBinding creates a ssl group binding and then deletes it.
func TestAddDeleteSslcipherSslciphersuiteBinding(t *testing.T) {
	var err error
	ciphergroupname := "Foo"
	ciphername := "TLS1-AES-256-CBC-SHA"
	description := "SSLv3 Kx=RSA      Au=RSA  Enc=AES(256)  Mac=SHA1   HexCode=0x0035"

	o := new(client.Netscaler)
	o.Session, err = connectToTestNetscaler()
	if err != nil {
		t.Error(err)
		return
	}

	defer o.Session.Logout()

	// Ensure cipher group does not already exist.
	_, err = o.GetSslciphergroup(ciphergroupname)
	if err != nil {
		// Create cipher group.
		ciphergroupReq := model.SslcipherAdd{}
		ciphergroupReq.Sslcipher.Ciphergroupname = ciphergroupname
		ciphergroupReq.Sslcipher.Ciphgrpalias = ciphergroupname
		_, err = o.AddSslciphergroup(ciphergroupReq)
		if err != nil {
			t.Error(err)
			return
		}
	}

	// Add cipher suite to cipher group.
	resp := model.SslcipherSslciphersuiteBinding{}
	req := model.SslcipherSslciphersuiteBindingAdd{}
	req.SSLCipherSSLCipherSuiteBinding.Ciphergroupname = ciphergroupname
	req.SSLCipherSSLCipherSuiteBinding.Ciphername = ciphername
	req.SSLCipherSSLCipherSuiteBinding.Cipherpriority = "1"

	expected := model.SslcipherSslciphersuiteBinding{}
	expected.Ciphergroupname = ciphergroupname
	expected.Stateflag = "512"
	expected.Ciphername = ciphername
	expected.Cipherpriority = "1"
	expected.Description = description
	expected.Peflags = "4"

	resp, err = o.AddSslcipherSslciphersuiteBinding(req)

	if err != nil {
		t.Error(err)
		return
	}

	if resp != expected {
		t.Errorf("failed to add cipher group binding: [expected] %+v [returned] %+v", expected, resp)
	}

	// Cleanup
	err = o.DeleteSslcipherSslciphersuiteBinding(ciphergroupname, ciphername)
	if err != nil {
		t.Error(err)
		return
	}

	err = o.DeleteSslciphergroup(ciphergroupname)
	if err != nil {
		t.Error(err)
	}
}
