package main

import (
	"fmt"

	"time"
)

func main() {
	var ii int64
	ii = int64(time.Now().Unix())
	tm := time.Unix(ii, 0)

	fmt.Println(tm)
}
