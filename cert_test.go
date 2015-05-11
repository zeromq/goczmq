package goczmq

import (
	"os"
	"testing"
)

func TestCert(t *testing.T) {
	cert := NewCert()
	defer cert.Destroy()

	cert.SetMeta("email", "taotetek@gmail.com")
	email := cert.Meta("email")
	if email != "taotetek@gmail.com" {
		t.Errorf("meta expected 'taotetek@gmail.com' got '%s'", email)
	}

	cert.SetMeta("name", "Brian Knox")
	name := cert.Meta("name")
	if name != "Brian Knox" {
		t.Errorf("Meta expected 'Brian Knox' got '%s'", name)
	}

	cert.SetMeta("organization", "ZeroMQ")
	organization := cert.Meta("organization")
	if organization != "ZeroMQ" {
		t.Errorf("Meta expected 'ZeroMQ' got '%s;", organization)
	}

	cert.SetMeta("version", "1")
	version := cert.Meta("version")
	if version != "1" {
		t.Errorf("Meta expected '1' got '%s'", version)
	}

	_ = cert.PublicText()

	// copy the cert
	dup := cert.Dup()
	defer dup.Destroy()

	name = dup.Meta("name")
	if name != "Brian Knox" {
		t.Errorf("Meta expected 'Brian Knox' got '%s'", name)
	}

	// test equality
	if !cert.Equal(dup) {
		t.Error("Duplicated cert should be equal, is not.")
	}

	cert.Print()

	cert.Save("./test_cert")
	loaded, err := NewCertFromFile("./test_cert")
	if err != nil {
		t.Error("NewCertFromFile failed")
	}
	defer loaded.Destroy()

	if !loaded.Equal(cert) {
		t.Error("Loaded cert is not equal to saved cert")
	}

	loaded.Print()

	os.Remove("./test_cert")
	os.Remove("./test_cert_secret")
}

func ExampleCert() {
	cert := NewCert()
	defer cert.Destroy()
	cert.SetMeta("email", "taotetek@gmail.com")
	cert.SetMeta("name", "Brian Knox")
	cert.SetMeta("organization", "ZeroMQ")
}
