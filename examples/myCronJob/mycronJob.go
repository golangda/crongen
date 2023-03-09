package main

// import (
// 	"fmt"
// 	"time"
// 	"github.com\golangda\crongen"
// )

// func main() {
// 	done := make(chan bool)
// 	Loc, _ := time.LoadLocation("Local")
// 	now := time.Now()
// 	fmt.Println("now:", now)
// 	firstInvokeSecond := 10
// 	InvokeIntervalSeconds := 10
// 	message := "myCronJob invoked at:"

// 	routineToInvoke := func() { myCronJob(message) }

// 	crgen := &crongen.CronGen{
// 		FirstInvokeYear:       now.Year(),
// 		FirstInvokeMonth:      now.Month(),
// 		FirstInvokeDay:        now.Day(),
// 		FirstInvokeSecond:     firstInvokeSecond,
// 		Loc:                   Loc,
// 		InvokeIntervalSeconds: InvokeIntervalSeconds,
// 		RoutineToInvoke:       routineToInvoke,
// 	}
// 	go crgen.CreateCronJob()
// 	<-done
// }

// func myCronJob(message string) {
// 	fmt.Println(message, time.Now().UTC())
// }
