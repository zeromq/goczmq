package goczmq

import (
	"testing"
)

func TestNewPicture(t *testing.T) {
	picture, err := newPicture("the", []byte("answer"), []byte(""), 42)
	if err != nil {
		t.Error("newPicture: %s", err)
	}

	if picture.picture != "sbni" {
		t.Errorf("picture.picture should be 'sbni', is %s", picture.picture)
	}

	if picture.pages[0] != "the" {
		t.Errorf("picture.pages[0] should be 'the', is %s", picture.pages[0])
	}

	if picture.pages[1] != "answer" {
		t.Errorf("picture.pages[1] should be 'answer', is %s", picture.pages[1])
	}

	if picture.pages[2] != "" {
		t.Errorf("picture.pages[2] should be '', is %s", picture.pages[2])
	}

	if picture.pages[3] != "42" {
		t.Errorf("picture.pages[3] should be '42', is %s", picture.pages[3])
	}
}

func BenchmarkNewPicture(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := newPicture("the", []byte("answer"), []byte(""), 42)
		if err != nil {
			b.Fatalf("newPicture: %s", err)
		}
	}
}
