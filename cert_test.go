package goczmq

import (
	"os"
	"testing"
)

func TestCert(t *testing.T) {
	cert := NewCert()
	defer cert.Destroy()

	cert.SetMeta("email", "taotetek@gmail.com")
	if want, got := "taotetek@gmail.com", cert.Meta("email"); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	cert.SetMeta("name", "Brian Knox")
	if want, got := "Brian Knox", cert.Meta("name"); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	cert.SetMeta("organization", "ZeroMQ")
	if want, got := "ZeroMQ", cert.Meta("organization"); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	cert.SetMeta("version", "1")
	if want, got := "1", cert.Meta("version"); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	_ = cert.PublicText()

	dup := cert.Dup()
	defer dup.Destroy()

	if want, got := "Brian Knox", dup.Meta("name"); want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	if want, got := true, cert.Equal(dup); want != got {
		t.Errorf("want '%v', got '%v'", want, got)
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

	if want, got := true, loaded.Equal(cert); want != got {
		t.Errorf("want '%v', got '%v'", want, got)
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
