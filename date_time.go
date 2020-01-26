package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	septDate := current.AddDate(0, 1, 0)

	fmt.Println(current.String())
	fmt.Println("MM-DD-YYYY", current.Format("01-02-2006"))
	fmt.Println("YYYY-MM-DD hh:mm:ss", current.Format("2006-01-02 15:04:05"))
	fmt.Println(septDate.String())
	beforeDate := current.AddDate(-1, 0, 0)
	fmt.Println(beforeDate.String())

	tenMoreMinutes := beforeDate.Add(10 * time.Minute)
	fmt.Println(tenMoreMinutes)
}
