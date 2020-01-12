package interval

import "fmt"

// TimeInterval represent a time interval in millis
type TimeInterval struct {
	// beginning of the time interval
	StartTime int64
	// end of the time interval
	EndTime int64
}

// IntersectWith return true if "t" intereset with "anotherT"
func (t TimeInterval) IntersectWith(anotherT TimeInterval) bool {

	// "t" completely within "anotherT"
	if anotherT.StartTime <= t.StartTime && t.EndTime <= anotherT.EndTime {
		return true
	}

	// "anotherT" upper bound is within "t"
	if t.StartTime < anotherT.EndTime && anotherT.EndTime <= t.EndTime {
		return true
	}

	// "anotherT" lower bound is within "t"
	if t.StartTime <= anotherT.StartTime && anotherT.StartTime < t.EndTime {
		return true
	}

	return false
}

// NewTimeInterval return a TimeInterval if "start < end" is verified
func NewTimeInterval(start, end int64) (TimeInterval, error) {
	if start < end {
		return TimeInterval{start, end}, nil
	}
	return TimeInterval{}, fmt.Errorf("condition violated: the start of the interval MUST NOT be higher than the end")
}
