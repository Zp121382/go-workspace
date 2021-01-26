/**
* The errors_my program just says Something went wrong
* @author Zp121382
* @version 0.1
* @since 2021-1-25
 */

package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Something went wrong")
	if err != nil {
		fmt.Println(err)
	}
}
