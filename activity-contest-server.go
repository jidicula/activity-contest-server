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
	activitySummaries []ActivitySummary
}

type ActivitySummary struct {
	gorm.Model
	MovePercent     uint
	ExercisePercent uint
	StandPercent    uint
}

// score computes the score from an ActivitySummary.
func (as ActivitySummary) score() uint {
	return as.MovePercent + as.ExercisePercent + as.StandPercent
}

func main() {
	as := ActivitySummary{
		MovePercent:     100,
		ExercisePercent: 30,
		StandPercent:    10,
	}
	u := user{
		username:          "test",
		password:          "password",
		activitySummaries: []ActivitySummary{as},
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
		`ActivitySummary:
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
