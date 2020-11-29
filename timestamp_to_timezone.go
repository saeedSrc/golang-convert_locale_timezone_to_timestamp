package main
import (
"time"
"fmt"
)

func main() {
	fmt.Println("canvert timestamp to time zone began ...")
	loc, err := time.LoadLocation("Asia/Kuwait")
	if err != nil {
		fmt.Println(err)
		return
	}

	date := time.Unix(1605611754, 0).In(loc)
	fmt.Println(date.Format("Monday 2006-01-02 15:04:05"))
	fmt.Println("Done")
}
