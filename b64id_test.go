package b64id_test

import (
	"testing"

	"github.com/sdqri/b64id"
)

func TestGenerateB64IDLength(t *testing.T) {
	lengths := []b64id.IDLength{b64id.IDL2, b64id.IDL3, b64id.IDL4, b64id.IDL8, b64id.IDL10, b64id.IDL11, b64id.IDL12, b64id.IDL14, b64id.IDL15, b64id.IDL16}
	for _, l := range lengths {
		res := b64id.GenerateB64ID(l)
		if len(res) != int(l) {
			t.Errorf("Expected generateB64ID(%v)==%v, Got %v", l, l, len(res))
		}
	}
}

func TestGenerateB64IDUnique(t *testing.T) {
	id1 := b64id.GenerateB64ID(b64id.IDL8)
	id2 := b64id.GenerateB64ID(b64id.IDL8)
	if id1 == id2 {
		t.Errorf("Expected unique value when calling generateB64ID(%v) twice but got %v, %v", b64id.IDL8, id1, id2)
	}
}
