package main

import "testing"

// TestActivitySummaryScore ...
func TestActivitySummaryScore(t *testing.T) {
	tests := map[string]struct {
		MovePercent     uint
		ExercisePercent uint
		StandPercent    uint
		want            uint
	}{
		"correctly constructed ActivitySummary object": {
			MovePercent:     100,
			ExercisePercent: 30,
			StandPercent:    10,
			want:            140,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			as := activitySummary{
				movePercent:     tt.MovePercent,
				exercisePercent: tt.ExercisePercent,
				standPercent:    tt.StandPercent,
			}
			got := as.score()
			if got != tt.want {
				t.Errorf("%s: got %v, want %v", name, got, tt.want)
			}
		})
	}

}
