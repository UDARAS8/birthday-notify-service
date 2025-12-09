package main

import (
	"blazeisclone/birthday-notifier/internal"
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

func checkAndNotifyBirthdays() {
	birthdays, err := internal.FetchBirthdays()
	if err != nil {
		log.Printf("Error fetching birthdays: %v", err)
		return
	}

	todayBirthdays := internal.GetTodayBirthdays(birthdays)

	if len(todayBirthdays) == 0 {
		log.Println("No birthdays today")
		return
	}

	for _, birthday := range todayBirthdays {
		subject := "Happy Birthday"
		message := fmt.Sprintf(
			"Happy Birthday %s, We hope you have a wonderful day filled with joy and celebration!",
			birthday.Name,
		)

		internal.SendMail(birthday.Name, birthday.Email, subject, message)
		log.Printf("Sent birthday notification to %s (%s)", birthday.Name, birthday.Email)
	}
}

func main() {
	scheduler := cron.New()
	cronExpression := "0 12 * * *"

	checkAndNotifyBirthdays()

	scheduler.AddFunc(cronExpression, checkAndNotifyBirthdays)

	scheduler.Start()

	log.Println("Scheduler started.")

	select {}
}
