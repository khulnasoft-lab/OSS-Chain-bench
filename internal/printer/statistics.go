package printer

import (
	"fmt"

	"github.com/khulnasoft-lab/oss-chain-bench/internal/models/checkmodels"
)

type Statistics struct {
	Passed  int `json:"passed,omitempty"`
	Failed  int `json:"failed,omitempty"`
	Unknown int `json:"unknown,omitempty"`
	Total   int `json:"total,omitempty"`
}

// NewStatistics initializes a new Statistics struct.
func NewStatistics() Statistics {
	return Statistics{Passed: 0, Failed: 0, Unknown: 0, Total: 0}
}

// Add increments the value of a specific field as well as the total value.
func (s *Statistics) Add(value checkmodels.ResultStatus) error {
	switch value {
	case checkmodels.Passed:
		s.Passed++
	case checkmodels.Failed:
		s.Failed++
	case checkmodels.Unknown:
		s.Unknown++
	default:
		return fmt.Errorf("unknown statistical value: %s", value)
	}
	s.Total++

	return nil
}

// Sub decrements the value of a specific field as well as the total value.
func (s *Statistics) Sub(value checkmodels.ResultStatus) error {
	switch value {
	case checkmodels.Passed:
		s.Passed--
	case checkmodels.Failed:
		s.Failed--
	case checkmodels.Unknown:
		s.Unknown--
	default:
		return fmt.Errorf("unknown statistical value: %s", value)
	}

	s.Total--

	return nil
}
