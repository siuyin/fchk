package fchk

import (
	"testing"
)

func TestCheck(t *testing.T) {
	a := Args{"junk.txt", "/home/siuyin,/home/kitsiew"}
	f := FChk{}
	var r []FChk
	err := f.Check(&a, &r)
	if err != nil {
		t.Error("bad")
	}
	if len(r) != 2 {
		t.Errorf("bad reply len: %d", len(r))
	}
	if r[0].Path != "/home/siuyin/junk.txt" {
		t.Errorf("bad path: %s", r[0].Path)
	}
	if !r[0].Found {
		t.Error("should be found")
	}
	if r[1].Found != false {
		t.Error("should not be found")
	}
}
