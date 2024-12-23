# crongen

**Go Library to help you spawn your own Go Cron Job**

## Overview

`crongen` is a lightweight library that simplifies the process of scheduling cron jobs in Go. You can use this library to create flexible and reusable cron job configurations tailored to your application's needs.

## Features

- Easy-to-use configuration for cron job scheduling.
- Customizable parameters for the first invocation and periodic intervals.
- Local and UTC timezone support.
- Lightweight and optimized for performance.

## Installation

To use `crongen`, include it in your project by running:

```bash
go get github.com/golangda/crongen
```

## Usage
Here's how you can use crongen to set up a cron job:

### Example
Below is an example to configure and run a cron job:
```go
package main

import (
	"fmt"
	"time"

	"github.com/golangda/crongen"
)

func main() {
	done := make(chan bool)
	loc, _ := time.LoadLocation("Local")
	now := time.Now()

	// Routine to be invoked as a cron job
	routineToInvoke := func() { fmt.Println("Cron job invoked at:", time.Now()) }

	// Configure the cron job
	crgen := &crongen.CronGen{
		FirstInvokeYear:       now.Year(),
		FirstInvokeMonth:      now.Month(),
		FirstInvokeDay:        now.Day(),
		FirstInvokeHour:       now.Hour(),
		FirstInvokeMin:        now.Minute(),
		FirstInvokeSecond:     30,
		Loc:                   loc,
		InvokeIntervalSeconds: 15,
		RoutineToInvoke:       routineToInvoke,
	}

	// Start the cron job
	go crgen.CreateCronJob()
	<-done
}
```

Configuration Parameters
- **FirstInvokeYear:** The year for the first job execution. Default: current year.
- **FirstInvokeMonth:** The month for the first execution. Default: current month.
- **FirstInvokeDay:** The day for the first execution. Default: current day.
- **FirstInvokeHour/Min/Second/Nanosecond:** The time for the first execution. Defaults to midnight.
- **Loc:** The timezone for execution. Default: UTC.
- **InvokeIntervalHours/Mins/Seconds/Nanoseconds:** The periodic execution interval.


## API Documentation
### Constants
No specific constants are defined in this library.

### Variables
No global variables are defined in this library.

### Types
**CronGen**
`CronGen` is the main type used to configure and manage cron jobs. It includes the following fields:

- **FirstInvokeYear (int):** The year for the first job execution.
- **FirstInvokeMonth (time.Month):** The month for the first job execution.
- **FirstInvokeDay (int):** The day for the first job execution.
- **FirstInvokeHour (int):** The hour for the first job execution.
- **FirstInvokeMin (int):** The minute for the first job execution.
- **FirstInvokeSecond (int):** The second for the first job execution.
- **FirstInvokeNanoSecond (int):** The nanosecond for the first job execution.
- **Loc (*time.Location):** The timezone for execution.
- **InvokeIntervalHours (int):** Hours between periodic executions.
- **InvokeIntervalMins (int):** Minutes between periodic executions.
- **InvokeIntervalSeconds (int):** Seconds between periodic executions.
- **InvokeIntervalNanoSeconds (int):** Nanoseconds between periodic executions.
- **RoutineToInvoke (func()):** The function to be invoked as the cron job.

### Functions
**func (cronGen *CronGen) CreateCronJob()**
This function initiates the cron job using the parameters specified in the CronGen struct. It ensures that the job is executed first at the defined time and then repeatedly at the specified intervals.

#### Helper Functions
**getNextInvokeAT(firstInvokeTime time.Time, invokeInterval time.Duration) time.Time:** Calculates the next execution time based on the provided interval.

## Example Project
Check the `examples/` folder for a fully functional example project.

---
# License
This project is licensed under the MIT License. See the LICENSE file for details.

---

# Contributing
We welcome contributions to make crongen even better! Here's how you can help:

## Report Bugs: If you encounter any issues, please create a GitHub issue with detailed information about the problem.

Request Features: Suggest new features or enhancements by creating a GitHub issue.

## Submit Code:

- Fork the repository.
- Make your changes in a new branch.
- Test your changes thoroughly.
- Submit a pull request to the main branch with a clear description of your changes.
- Improve Documentation: If you find areas where the documentation can be improved, feel free to update it and submit a pull request.

## Guidelines
- Follow the Go coding conventions.
- Write clear and concise commit messages.
- Ensure all tests pass before submitting a pull request.
- Add unit tests for any new functionality.

We appreciate your contributions and look forward to your involvement!

---