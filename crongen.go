package crongen

import (
	"fmt"
	"time"
)

type CronGen struct {
	FirstInvokeYear           int
	FirstInvokeMonth          time.Month
	FirstInvokeDay            int
	FirstInvokeHour           int
	FirstInvokeMin            int
	FirstInvokeSecond         int
	FirstInvokeNanoSecond     int
	Loc                       *time.Location
	InvokeIntervalHours       int
	InvokeIntervalMins        int
	InvokeIntervalSeconds     int
	InvokeIntervalNanoSeconds int
	RoutineToInvoke           func()
}

func (cronGen *CronGen) CreateCronJob() {
	now := time.Now()
	switch {
	case cronGen.Loc == nil:
		cronGen.Loc, _ = time.LoadLocation("")
	case cronGen.FirstInvokeYear == 0:
		cronGen.FirstInvokeYear = now.Year()
	case cronGen.FirstInvokeMonth == 0:
		cronGen.FirstInvokeMonth = now.Month()
	}
	firstInvokeTime := time.Date(cronGen.FirstInvokeYear, cronGen.FirstInvokeMonth, cronGen.FirstInvokeDay,
		cronGen.FirstInvokeHour, cronGen.FirstInvokeMin, cronGen.FirstInvokeSecond, cronGen.FirstInvokeNanoSecond, cronGen.Loc)
	fmt.Println("firstInvokeTime:", firstInvokeTime)
	invokeInterval := time.Duration(cronGen.InvokeIntervalHours)*time.Hour + time.Duration(cronGen.InvokeIntervalMins)*time.Minute +
		time.Duration(cronGen.InvokeIntervalSeconds)*time.Second + time.Duration(cronGen.InvokeIntervalNanoSeconds)
	fmt.Println("invokeInterval:", invokeInterval)
	for {
		timer := time.NewTimer(invokeInterval)
		<-timer.C
		go cronGen.RoutineToInvoke()
	}
}
