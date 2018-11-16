package cursor

import (
	"testing"
)

func TestCursor(t *testing.T) {
	var b = []byte("select * from stu")

	// spawn a new Cursor
	var c = NewCursor(b,0)

	c.Next()
	if c.Index != 0 {
		t.Errorf("Next() skipped Index 0")
	}

	c.Next()
	if c.Index != 1 {
		t.Errorf("Next() failed to Increment")
	}

	c.SetByte(byte('!'))
	if c.GetByte() != byte('!') {
		t.Errorf("SetByte() failed to set the provided byte")
	}

	// these will create out of range errors if they dont perform correctly
	var c2 = NewCursor(b,0)
	for c2.Next() {

	}
	for c2.Prev() {

	}

}
