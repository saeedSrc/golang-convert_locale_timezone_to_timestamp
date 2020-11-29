package main
import (
"time"
"fmt"
)

func main() {
	fmt.Println("convert timezone to timestamp began ...")
	loc, err := time.LoadLocation("Asia/Kuwait")
	if err != nil {
		fmt.Println(err)
		return
	}

	date := time.Now().In(loc)
	fmt.Println(date.Unix())
	fmt.Println("Done")
}
