package usecase

import (
	"Media/models"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestCulc(t *testing.T) {
	logger := logrus.New()
	usecase := NewCase(logger)

	tests := []struct {
		name     string
		cashInfo models.CashInfo
		expected [][]int
	}{
		{
			name: "Test with 400 amount and multiple banknotes",
			cashInfo: models.CashInfo{
				Amount: 400,
				Banknotes: []int{
					5000, 2000, 1000, 500, 200, 100, 50,
				},
			},
			expected: [][]int{
				{200, 200},
				{200, 100, 100},
				{200, 100, 50, 50},
				{200, 50, 50, 50, 50},
				{100, 100, 100, 100},
				{100, 100, 100, 50, 50},
				{100, 100, 50, 50, 50, 50},
				{100, 50, 50, 50, 50, 50, 50},
				{50, 50, 50, 50, 50, 50, 50, 50},
			},
		},
		{
			name: "Test with 100 amount and single banknote",
			cashInfo: models.CashInfo{
				Amount:    100,
				Banknotes: []int{100},
			},
			expected: [][]int{
				{100},
			},
		},
		{
			name: "Test with 150 amount and no possible combination",
			cashInfo: models.CashInfo{
				Amount:    150,
				Banknotes: []int{200},
			},
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase.Cash = tt.cashInfo
			response := usecase.Culc()

			if !equal(response.Exchanges, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, response.Exchanges)
			}
		})
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
