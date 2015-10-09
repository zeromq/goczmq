package goczmq

import (
	"os"
	"testing"
)

func TestCert(t *testing.T) {
	cert := NewCert()
	defer cert.Destroy()

	cert.SetMeta("email", "taotetek@gmail.com")
	if want, have := "taotetek@gmail.com", cert.Meta("email"); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	cert.SetMeta("name", "Brian Knox")
	if want, have := "Brian Knox", cert.Meta("name"); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	cert.SetMeta("organization", "ZeroMQ")
	if want, have := "ZeroMQ", cert.Meta("organization"); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	cert.SetMeta("version", "1")
	if want, have := "1", cert.Meta("version"); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	_ = cert.PublicText()

	dup := cert.Dup()
	defer dup.Destroy()

	if want, have := "Brian Knox", dup.Meta("name"); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if want, have := true, cert.Equal(dup); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if testing.Verbose() {
		cert.Print()
	}

	cert.Save("./test_cert")
	loaded, err := NewCertFromFile("./test_cert")
	if err != nil {
		t.Error(err)
	}
	defer loaded.Destroy()

	if want, have := true, loaded.Equal(cert); want != have {
		t.Errorf("want %#v, have %#v", want, have)
	}

	if testing.Verbose() {
		loaded.Print()
	}

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
