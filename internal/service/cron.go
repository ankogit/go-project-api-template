package service

import (
	"github.com/robfig/cron/v3"
)

type CronService struct {
	Scheduler *cron.Cron
	Users     Users
}

func NewCronService(sh *cron.Cron, userService Users) *CronService {
	return &CronService{
		Scheduler: sh,
		Users:     userService,
	}
}

func (c *CronService) Start() {
	c.Scheduler.Start()
}

func (c *CronService) Init() {
}

func (c *CronService) SetJob(notificationCron string) (cron.EntryID, error) {
	return 0, nil
}

func (c *CronService) RemoveJob() {

}
