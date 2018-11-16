package cursor

/*
  The Cursor package makes moving up and down a bit slices easier.
	Cursors are intended to have a short life span, used to perform an operation
	within a function.
*/

// NewCursor is a Cursor constructor
func NewCursor(b []byte, startIndex int) (c Cursor) {
	c.bytes = b
	c.Index = startIndex
	return c
}

// Cursor
type Cursor struct {
	Index     int
	bytes     []byte
	Increment bool
	Decrement bool
}

// GetByte return the document byte at the Cursors current Index
func (c *Cursor) GetByte() byte {
	return (*c).bytes[(*c).Index]
}

// SetByte updates the document byte at the Cursors current Index
func (c *Cursor) SetByte(b byte) {
	(*c).bytes[(*c).Index] = b
}

// sets the Index of the Cursor at the begining of the byte array
func (c *Cursor) ToStart() {
	(*c).Index = 0
	(*c).Increment = false
}

// sets the Index of the Cursor at the end of the byte array
func (c *Cursor) ToEnd() {
	(*c).Index = len((*c).bytes) - 1
	(*c).Decrement = false
}

// DontSkip resets the decr and incr boolean to false
func (c *Cursor) DontSkip() {
	(*c).Increment = false
	(*c).Decrement = false
}

// DontSkip sets the decr and incr boolean to true
func (c *Cursor) Skip() {
	(*c).Increment = true
	(*c).Decrement = true
}


// Forward moves the Cursor forward n bytes. Will return an false if Index is
// larger then document byte array
func (c *Cursor) Forward(n int) bool {
	if !(((*c).Index + n) < len((*c).bytes)) {
		(*c).Decrement = false
		return false
	}
	(*c).Index += n
	return true
}

// Backward moves the Cursor Backward n bytes and return true if the new Cursor
// Index would be is less then zero.
func (c *Cursor) Backward(n int) bool {
	if ((*c).Index - n) < 0 {
		(*c).Increment = false
		return false
	}
	(*c).Index -= n
	return true
}

// Next Increments the Cursor by 1. Next is intended to be used in a loop.
func (c *Cursor) Next() (b bool) {
	if !(*c).Increment {
    (*c).Increment = true
		return true
	}
	b = (*c).Forward(1)
	if !b {
		(*c).Decrement = false
	}
	return b
}

// Prev Decrements the Cursor by 1. Prev is intended to be used in a loop.
func (c *Cursor) Prev() (b bool) {
	if !(*c).Decrement {
		(*c).Decrement = true
		return true
	}
	b = (*c).Backward(1)
	if !b {
		(*c).Increment = false
	}
	return b
}
