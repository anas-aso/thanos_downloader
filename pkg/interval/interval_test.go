// Copyright 2020 Anas Ait Said Oubrahim

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package interval

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTimeInterval(t *testing.T) {
	var cases = []struct {
		name  string
		start int64
		end   int64
		err   error
	}{
		{
			name:  "good_values",
			start: 0,
			end:   1,
			err:   nil,
		},
		{
			name:  "bad_values",
			start: 1,
			end:   0,
			err:   fmt.Errorf("condition violated: the start of the interval MUST NOT be higher than the end"),
		},
	}

	for _, c := range cases {
		v, err := NewTimeInterval(c.start, c.end)

		if err != nil && err.Error() != c.err.Error() {
			t.Errorf("Test case : %v failed.\nExpected error: %v\nReturned error: %v\n", c.name, c.err, err)
		}

		if err == nil {
			if v.StartTime != c.start || v.EndTime != c.end {
				t.Errorf("Test case : %v failed.\nWrong struct values: %v\n", c.name, v)
			}
		}
	}
}

func TestIntersectWith(t *testing.T) {
	var cases = []struct {
		name           string
		t1             TimeInterval
		t2             TimeInterval
		expectedResult bool
	}{
		{
			name:           "t1_equals_to_t2",
			t1:             TimeInterval{100 * 1000, 200 * 1000},
			t2:             TimeInterval{100 * 1000, 200 * 1000},
			expectedResult: true,
		},
		{
			name:           "t1_completely_within_t2",
			t1:             TimeInterval{100 * 1000, 200 * 1000},
			t2:             TimeInterval{0, time.Now().Unix() * 1000},
			expectedResult: true,
		},
		{
			name:           "t1_includes_t2",
			t1:             TimeInterval{0, time.Now().Unix() * 1000},
			t2:             TimeInterval{100 * 1000, 200 * 1000},
			expectedResult: true,
		},
		{
			name:           "t1_upper_bound_within_t2",
			t1:             TimeInterval{50 * 1000, 150 * 1000},
			t2:             TimeInterval{100 * 1000, 200 * 1000},
			expectedResult: true,
		},
		{
			name:           "t1_lower_bound_within_t2",
			t1:             TimeInterval{150 * 1000, 250 * 1000},
			t2:             TimeInterval{100 * 1000, 200 * 1000},
			expectedResult: true,
		},
		{
			name:           "t1_does_not_intersect_with_t2",
			t1:             TimeInterval{100 * 1000, 200 * 1000},
			t2:             TimeInterval{300 * 1000, 400 * 1000},
			expectedResult: false,
		},
		{
			name:           "t1_intersect_with_on_the_edge_t2",
			t1:             TimeInterval{100 * 1000, 200 * 1000},
			t2:             TimeInterval{200 * 1000, 400 * 1000},
			expectedResult: false,
		},
	}

	for _, c := range cases {
		if c.t1.IntersectWith(c.t2) != c.expectedResult {
			t.Errorf("Test case : %v failed.\n", c.name)
		}
	}
}
