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
	contestEntries []ContestEntry
}

type ContestEntry struct {
	gorm.Model
	User  User
	Score uint
}

type User struct {
	gorm.Model
	Username          string `gorm:"unique"`
	Password          string
	ActivitySummaries []ActivitySummary
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
	u := User{
		Username:          "test",
		Password:          "password",
		ActivitySummaries: []ActivitySummary{as},
	}
	ce := ContestEntry{
		User:  u,
		Score: as.score(),
	}
	c := contest{
		startDate:      time.Now(),
		endDate:        time.Now().Add(time.Hour * 24 * 7),
		inProgress:     true,
		contestEntries: []ContestEntry{ce},
	}
	fmt.Printf(
		`ActivitySummary:
  %+v

User:
  %+v

ContestEntry:
  %+v

contest:
  %+v
`,
		as, u, ce, c)
}
