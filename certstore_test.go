package goczmq

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCertStore(t *testing.T) {
	testDir := ".test_zcertstore"
	err := os.Mkdir(testDir, 0777)
	require.NoError(t, err)

	certstore := NewCertStore(testDir)
	defer certstore.Destroy()

	client0 := NewCert()
	client0Key := client0.PublicText()
	client0.Destroy()
	client0Loaded := certstore.Lookup(client0Key)

	if client0Loaded != nil {
		t.Errorf("want %#v, have %#v", nil, client0Loaded)
	}

	client1 := NewCert()
	client1.SetMeta("name", "Brian")
	client1Key := client1.PublicText()

	err = client1.Save(testDir + "/mycert.txt")
	require.NoError(t, err)

	client1.Destroy()

	client1Loaded := certstore.Lookup(client1Key)

	if want, have := "Brian", client1Loaded.Meta("name"); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := client1Key, client1Loaded.PublicText(); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	client2 := NewCert()
	client2.SetMeta("name", "Luna")
	client2Key := client2.PublicText()

	certstore.Insert(client2)
	client2Loaded := certstore.Lookup(client2Key)

	if want, have := "Luna", client2Loaded.Meta("name"); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := client2Key, client2Loaded.PublicText(); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	err = os.RemoveAll(testDir)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCertStoreMem(t *testing.T) {
	certstore := NewCertStoreInMemory()
	defer certstore.Destroy()

	client0 := NewCert()
	client0Key := client0.PublicText()
	client0.Destroy()
	client0Loaded := certstore.Lookup(client0Key)

	if client0Loaded != nil {
		t.Errorf("want %#v, have %#v", nil, client0Loaded)

	}
}
