# package cursor
A Golang package for stepping through byte slices.



```go
package main

import (
	"fmt"
	"github.com/jakobii/cursor"
)

func main() {

	var b = []byte("select * from stu")

	// spawn a new cursor
	var c = cursor.NewCursor(b,0)


	// forward
	for {
		fmt.Println(c.Index, string(c.GetByte()))

		if !c.Next() {
			break
		}
	}

  // backward
	for {
		fmt.Println(c.Index, string(c.GetByte()))

		if !c.Prev() {
			break
		}
	}

}
```
