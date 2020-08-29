package cron

import (
	"github.com/robfig/cron/v3"

	"github.com/owarai/zgh/log"
)

// Field name   | Mandatory? | Allowed values  | Allowed special characters
// ----------   | ---------- | --------------  | --------------------------
// Seconds      | Yes        | 0-59            | * / , -
// Minutes      | Yes        | 0-59            | * / , -
// Hours        | Yes        | 0-23            | * / , -
// Day of month | Yes        | 1-31            | * / , - ?
// Month        | Yes        | 1-12 or JAN-DEC | * / , -
// Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?

// @Note Note that this spec is not sorted by minute, hour , day, month, and week.
func ZgCron(spec string, f func()) {
	c := cron.New()
	_, _ = c.AddFunc(spec, func() {
		f()
		log.L().Info("ZgCron", "ZgCron", "Function", f)
	})

	c.Start()

	// go func() {
	//	for {
	//		f()
	//		now := time.Now()
	//		next := now.Add(duration)
	//		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
	//		t := time.NewTimer(next.Sub(now))
	//		<-t.C
	//	}
	// }()
}
