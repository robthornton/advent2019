package main

import (
	"fmt"
	"strconv"

	"github.com/robthornton/advent2019/password"
)

func main() {
	matches := 0
	for i := int64(109165); i <= 576723; i++ {
		if ok := password.ValidElfPassword2(strconv.FormatInt(i, 10)); ok {
			matches++
		}
	}

	fmt.Println(matches)
}
