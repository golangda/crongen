package main

import (
	"fmt"
	"time"

	"github.com/golangda/crongen"
)

func main() {
	done := make(chan bool)
	Loc, _ := time.LoadLocation("Local")
	now := time.Now()
	fmt.Println("now:", now)
	firstInvokeSecond := 15
	InvokeIntervalSeconds := 10
	message := "myGoFunction invoked at:"

	// routine to invoke as a cron job
	routineToInvoke := func() { myGoFunction(message) }

	// configure cron job
	crgen := &crongen.CronGen{
		// FirstInvokeYear:       now.Year(),
		FirstInvokeMonth:      now.Month(),
		FirstInvokeDay:        now.Day(),
		FirstInvokeHour:       now.Hour(),
		FirstInvokeMin:        now.Minute(),
		FirstInvokeSecond:     firstInvokeSecond,
		FirstInvokeNanoSecond: 0,
		Loc:                   Loc,
		InvokeIntervalSeconds: InvokeIntervalSeconds,
		RoutineToInvoke:       routineToInvoke,
	}
	go crgen.CreateCronJob()
	<-done
}

// I want to convert this Go function into a cron job
func myGoFunction(message string) {
	fmt.Println(message, time.Now().Local())
}
