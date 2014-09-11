package goczmq

import "C"

import (
	"errors"
	"fmt"
	"strings"
)

type picturePage struct {
	intVal  int
	strVal  string
	byteVal []byte
}
type picture struct {
	picture string
	pages   []*picturePage
}

func newPicture(parts ...interface{}) (*picture, error) {
	numParts := len(parts)

	p := &picture{
		pages: make([]*picturePage, numParts),
	}

	pictureType := make([]string, numParts)

	for i, val := range parts {
		switch v := val.(type) {
		case int:
			pictureType[i] = "i"
			p.pages[i] = &picturePage{intVal: val.(int)}
		case string:
			pictureType[i] = "s"
			p.pages[i] = &picturePage{strVal: val.(string)}
		case []byte:
			if len(v) > 0 {
				pictureType[i] = "b"
			} else {
				pictureType[i] = "n"
			}
			p.pages[i] = &picturePage{byteVal: make([]byte, len(val.([]byte)))}
			p.pages[i].byteVal = val.([]byte)
		default:
			return nil, errors.New(fmt.Sprintf("unsupported type at index %d", i))
		}
	}

	p.picture = strings.Join(pictureType, "")

	return p, nil
}
