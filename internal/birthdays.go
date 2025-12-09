package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Birthday struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Department  string `json:"department"`
	DateOfBirth string `json:"dateOfBirth"`
}

type BirthdayResponse struct {
	Data []Birthday `json:"data"`
}

func FetchBirthdays() ([]Birthday, error) {
	apiURL := os.Getenv("BIRTHDAYS_API_URL")
	if apiURL == "" {
		return nil, fmt.Errorf("BIRTHDAYS_API_URL environment variable is not set")
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch birthdays: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var birthdayResponse BirthdayResponse
	if err := json.Unmarshal(body, &birthdayResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return birthdayResponse.Data, nil
}

func IsBirthdayToday(dateOfBirth string) (bool, error) {
	if dateOfBirth == "" {
		return false, fmt.Errorf("date of birth is empty")
	}

	birthDate, err := time.Parse(time.RFC3339, dateOfBirth)
	if err != nil {
		return false, fmt.Errorf("failed to parse date: %w", err)
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	thisYearBirthday := time.Date(now.Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, now.Location())

	return today.Month() == thisYearBirthday.Month() && today.Day() == thisYearBirthday.Day(), nil
}

func GetTodayBirthdays(birthdays []Birthday) []Birthday {
	var todayBirthdays []Birthday

	for _, birthday := range birthdays {
		isToday, err := IsBirthdayToday(birthday.DateOfBirth)
		if err != nil {
			log.Printf("Error checking birthday for %s: %v", birthday.Name, err)
			continue
		}

		if isToday {
			todayBirthdays = append(todayBirthdays, birthday)
		}
	}

	return todayBirthdays
}
