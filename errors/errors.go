// Errors
// Go returns errors as a separate return value (instead of using exceptions), easier to handle errors
// They are the last return value (w/ type error)
// 'nil' means no error

package main

import (
	"errors"
	"fmt"
)

func testNum(num int) (int, error) {
	if num == 42 {
		return -1, errors.New("Can't work with 42!")
	}

	return num + 3, nil
}

func main() {
	nums := []int{2, 42, 46, 1, 42}

	for _, num := range nums {
		result, err := testNum(num)

		if err != nil {
			fmt.Println("testNum failed:", err)
		} else {
			fmt.Println("testNum worked:", result)
		}
	}
}