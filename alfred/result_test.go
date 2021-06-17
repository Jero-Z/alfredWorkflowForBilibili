package alfred

import (
	"testing"
)

func TestResult_Append(t *testing.T) {
	result := NewResult()
	for i := 0; i < 10; i++ {
		result.Append(&Item{
			Valid: true,
			Title: "定军山",
		})
	}
}
