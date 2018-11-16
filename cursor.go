package cursor

/*
  The cursor package makes moving up and down a bit slices easier.
*/

// NewCursor is a cursor constructor
func NewCursor(b []byte) (c cursor) {
	c.bytes = b
	c.decrement = true
	return c
}

// cursor is an arrays current byte index
// the increment property is used to allow for the next func to emit a zero index.
type cursor struct {
	bytes     []byte
	index     int
	increment bool
	decrement bool
}

// GetByte return the document byte at the cursors current index
func (c *cursor) GetByte() byte {
	return (*c).bytes[(*c).index]
}

// SetByte updates the document byte at the cursors current index
func (c *cursor) SetByte(b byte) {
	(*c).bytes[(*c).index] = b
}

// Index returns the cursors current index.
func (c *cursor) Index() int {
	return int((*c).index)
}

// Forward moves the cursor forward n bytes. Will return an false if index is
// larger then document byte array
func (c *cursor) Forward(n int) bool {
	if ((*c).index + n) >= len((*c).bytes) {
		return false
	}
	(*c).index += n
	return true
}

// Backward moves the cursor Backward n bytes and return true if the new cursor
// index would be is less then zero.
func (c *cursor) Backward(n int) bool {
	if ((*c).index - n) < 0 {
		return false
	}
	(*c).index -= n
	return true
}

// Next increments the cursor by 1. if increment is set to false, it will set
// increment to equal true and it will skip incrementing the cursor.
func (c *cursor) Next() bool {
	if c.increment == false {
		(*c).increment = true
		return true
	}
	return (*c).Forward(1)
}

// Prev decrements the cursor by 1. if decrement is set to false, it will set
// decrement to equal true and it will skip decrementing the cursor. when
// using NewCursor the decrement property is set to false.
func (c *cursor) Prev() bool {
	if c.decrement == false {
		(*c).decrement = true
		return true
	}
	return (*c).Backward(1)
}
