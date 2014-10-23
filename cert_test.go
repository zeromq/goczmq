package goczmq

import (
	"testing"
)

func TestCert(t *testing.T) {
	cert := NewCert()

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

	cert.SetMeta("organization", "SPEE")
	organization := cert.Meta("organization")
	if organization != "SPEE" {
		t.Errorf("Meta expected 'SPEE' got '%s;", organization)
	}

	cert.SetMeta("version", "1")
	version := cert.Meta("version")
	if version != "1" {
		t.Errorf("Meta expected '1' got '%s'", version)
	}

	_ = cert.PublicText()

	cert.Destroy()
}
