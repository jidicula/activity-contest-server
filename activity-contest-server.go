package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type contest struct {
	gorm.Model
	startDate      time.Time
	endDate        time.Time
	inProgress     bool
	contestEntries []contestEntry
}

type contestEntry struct {
	gorm.Model
	user  user
	score uint
}

type user struct {
	gorm.Model
	username          string `gorm:"unique"`
	password          string
	activitySummaries []activitySummary
}

type activitySummary struct {
	gorm.Model
	movePercent     uint
	exercisePercent uint
	standPercent    uint
}

// score computes the score from an ActivitySummary.
func (as activitySummary) score() uint {
	return as.movePercent + as.exercisePercent + as.standPercent
}

func main() {
	as := activitySummary{
		movePercent:     100,
		exercisePercent: 30,
		standPercent:    10,
	}
	u := user{
		username:          "test",
		password:          "password",
		activitySummaries: []activitySummary{as},
	}
	ce := contestEntry{
		user:  u,
		score: as.score(),
	}
	c := contest{
		startDate:      time.Now(),
		endDate:        time.Now().Add(time.Hour * 24 * 7),
		inProgress:     true,
		contestEntries: []contestEntry{ce},
	}
	fmt.Printf(
		`activitySummary:
  %+v

user:
  %+v

contestEntry:
  %+v

contest:
  %+v
`,
		as, u, ce, c)
}
