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
		t.Error("picture.picture should be 'sbni', is %s", picture.picture)
	}

	if picture.pages[0].strVal != "the" {
		t.Error("picture.pages[0] should be 'the'")
	}

	if string(picture.pages[1].byteVal) != "answer" {
		t.Error("picture.pages[1] should be 'answer'")
	}

	if string(picture.pages[2].byteVal) != "" {
		t.Error("picture.pages[2] should be ''")
	}

	if picture.pages[3].intVal != 42 {
		t.Error("picture.pages[3] should be '42'")
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
