package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	StartDate time.Time
	EndDate   time.Time
	Users     []User
	Score     uint
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
	c := Contest{
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Hour * 24 * 7),
		Users:     []User{u},
		Score:     u.ActivitySummaries[0].score(),
	}
	fmt.Printf("ActivitySummary:\n  %+v\n\nUser:\n  %+v\n\nContest:\n  %+v\n", as, u, c)
}
