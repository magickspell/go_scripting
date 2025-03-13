package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOkText = `1
2
3
4
5`

var testOkTextResult = `1
2
3
4
5
`

func TestOK(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOkText))
	out := new(bytes.Buffer)
	err := unique(in, out)
	if err != nil {
		t.Errorf("test failed")
	}
	if out.String() != testOkTextResult {
		t.Errorf("test failed out\nin: %v\nout: %v", testOkText, out.String())
	}
}

var testBadText = `1
2
1`

func TestBad(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testBadText))
	out := new(bytes.Buffer)
	err := unique(in, out)
	if err == nil {
		t.Errorf("test failed")
	}
}
