package cursor

/*
  The Cursor package makes moving up and down a bit slices easier.
	Cursors are intended to have a short life span, used to perform an operation
	within a function.
*/

// NewCursor is a Cursor constructor
func NewCursor(b []byte, startIndex int) (c Cursor) {
	c.Index = startIndex
	c.bytes = b
	return c
}

// Cursor
type Cursor struct {
	Index     int
	bytes     []byte

}

// GetByte return the document byte at the Cursors current Index
func (c *Cursor) GetByte() byte {
	return (*c).bytes[(*c).Index]
}

// SetByte updates the document byte at the Cursors current Index
func (c *Cursor) SetByte(b byte) {
	(*c).bytes[(*c).Index] = b
}

// Forward moves the Cursor forward n bytes. Will return an false if Index is
// larger then document byte array
func (c *Cursor) Forward(n int) {
	(*c).Index += n
}

// Backward moves the Cursor Backward n bytes and return true if the new Cursor
// Index would be is less then zero.
func (c *Cursor) Backward(n int) {
	(*c).Index -= n
}

// Next Increments the Cursor by 1. Next is intended to be used in a loop.
func (c *Cursor) Next()(b bool) {
	if (*c).Index + 1 == len( (*c).bytes ) {
		return false
	}
	(*c).Forward(1)
	return true
}

// Prev Decrements the Cursor by 1. Prev is intended to be used in a loop.
func (c *Cursor) Prev() (b bool) {
	if c.Index == 0 {
		return false
	}
	(*c).Backward(1)
	return true
}
