package goczmq

import (
	"os"
	"testing"
)

func TestCertStore(t *testing.T) {
	testDir := ".test_zcertstore"
	os.Mkdir(testDir, 0777)

	certstore := NewCertStore(testDir)
	defer certstore.Destroy()

	client0 := NewCert()
	client0Key := client0.PublicText()
	client0.Destroy()
	client0Loaded := certstore.Lookup(client0Key)
	if client0Loaded != nil {
		t.Errorf("A Lookup of a cert that doesn't exist should return nil.")
	}

	client1 := NewCert()
	client1.SetMeta("name", "Brian")
	client1Key := client1.PublicText()
	client1.Save(testDir + "/mycert.txt")
	client1.Destroy()

	client1Loaded := certstore.Lookup(client1Key)
	if client1Loaded.Meta("name") != "Brian" {
		t.Errorf("Loaded cert should have name %s but has %s",
			"Brian", client1Loaded.Meta("name"))
	}

	if client1Key != client1Loaded.PublicText() {
		t.Errorf("client key should be %s is %s", client1Key, client1.PublicText())
	}

	client2 := NewCert()
	client2.SetMeta("name", "Luna")
	client2Key := client2.PublicText()

	certstore.Insert(client2)

	client2Loaded := certstore.Lookup(client2Key)
	if client2Loaded.Meta("name") != "Luna" {
		t.Errorf("Loaded cert should have name %s but has %s",
			"Luna", client2Loaded.Meta("name"))
	}

	if client2Key != client2Loaded.PublicText() {
		t.Errorf("Cert key should be %s is %s", client2Key, client2Loaded.PublicText())
	}

	err := os.RemoveAll(testDir)
	if err != nil {
		t.Fatal(err)
	}
}
