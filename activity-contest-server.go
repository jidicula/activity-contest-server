package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	StartDate      time.Time
	EndDate        time.Time
	InProgress     bool
	ContestEntries []ContestEntry
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

// score computes the score of an ActivitySummary.
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
	c := Contest{
		StartDate:      time.Now(),
		EndDate:        time.Now().Add(time.Hour * 24 * 7),
		InProgress:     true,
		ContestEntries: []ContestEntry{ce},
	}
	fmt.Printf(
		`ActivitySummary:
  %+v

User:
  %+v

ContestEntry:
  %+v

Contest:
  %+v
`,
		as, u, ce, c)
}
