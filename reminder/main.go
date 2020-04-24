package main

import (
	"github.com/sirupsen/logrus"
	"leetcode213/reminder/pkg/emails"
	"leetcode213/reminder/pkg/macros"
)
import "github.com/robfig/cron/v3"

func main() {
	box := &emails.Box{}

	c := cron.New()
	id, e := c.AddFunc(macros.SendMailSpecTest1, func() {
		box.SendMail("Hi")
	})
	if e == nil {
		logrus.Warnf("entry %d is configured", id)
	}
	c.Run()
}
