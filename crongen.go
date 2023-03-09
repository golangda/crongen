package crongen

import (
	"time"
)

// CronGen lets you configure the con job parameters such as
// FirstInvokeYear: The year in which you first want to invoke the cron job. Default is the current year.
// FirstInvokeMonth: The month of the FirstInvokeYear in which you first want to invoke the cron job. Default is the current month.
// FirstInvokeDay: The day of the FirstInvokeMonth on which you first want to invoke the cron job. Default is the present day.
// FirstInvokeHour: The hour of the FirstInvokeDay at which you first want to invoke the cron job. Default is zero(0)th hour.
// FirstInvokeMin: The minute of the FirstInvokeHour at which you first want to invoke the cron job. Default is zero(0)th minute.
// FirstInvokeSecond: The second of the FirstInvokeMin at which you first want to invoke the cron job. Default is zero(0)th second.
// FirstInvokeNanoSecond: The nanosecond of the FirstInvokeSecond at which you first want to invoke the cron job. Default is zero(0)th nanosecond.
// Loc: The time zone in which you want to run the cron job. Default is UTC.
// InvokeIntervalHours: The hour of the day at which you periodically want to invoke the cron job after the first invokation. Default is zero(0)th hour.
// InvokeIntervalMins: The minute of the InvokeIntervalHours at which you periodically want to invoke the cron job after the first invokation. Default is zero(0)th minute.
// InvokeIntervalSeconds: The second of the InvokeIntervalMins at which you periodically want to invoke the cron job after the first invokation. Default is zero(0)th second.
// InvokeIntervalNanoSeconds: The nanosecond of the InvokeIntervalSeconds at which you periodically want to invoke the cron job after the first invokation. Default is zero(0)th nanosecond.
// RoutineToInvoke: The go function to invoke as a cron job
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

// CreateCronJob ... creates a cron using CronGen field values
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
	invokeInterval := time.Duration(cronGen.InvokeIntervalHours)*time.Hour + time.Duration(cronGen.InvokeIntervalMins)*time.Minute +
		time.Duration(cronGen.InvokeIntervalSeconds)*time.Second + time.Duration(cronGen.InvokeIntervalNanoSeconds)
	for {
		timer := time.NewTimer(invokeInterval)
		<-timer.C
		if time.Now().Compare(firstInvokeTime) >= 0 {
			go cronGen.RoutineToInvoke()
		}
	}
}
