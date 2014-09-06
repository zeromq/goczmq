// A Go interface to CZMQ
package goczmq

import "C"

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type picture struct {
	picture string
	pages   []string
}

func newPicture(parts ...interface{}) (*picture, error) {
	numParts := len(parts)

	p := &picture{
		pages: make([]string, numParts),
	}

	picturePages := make([]string, numParts)

	for i, val := range parts {
		switch v := val.(type) {
		case int:
			picturePages[i] = "i"
			p.pages[i] = strconv.Itoa(val.(int))
		case string:
			picturePages[i] = "s"
			p.pages[i] = val.(string)
		case []byte:
			if len(v) > 0 {
				picturePages[i] = "b"
			} else {
				picturePages[i] = "n"
			}
			p.pages[i] = string(val.([]byte))
		default:
			return nil, errors.New(fmt.Sprintf("unsupported type at index %d", i))
		}
	}

	p.picture = strings.Join(picturePages, "")

	return p, nil
}
