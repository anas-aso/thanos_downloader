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

package blocks

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortingBlocks(t *testing.T) {
	var cases = []struct {
		name           string
		blocks         Blocks
		expectedResult Blocks
	}{
		{
			name: "differentTimes_sameRes_sameNumSamples",
			blocks: Blocks{
				{
					MinTime: 2,
					MaxTime: 3,
				},
				{
					MinTime: 1,
					MaxTime: 2,
				},
				{
					MinTime: 0,
					MaxTime: 1,
				},
			},
			expectedResult: Blocks{
				{
					MinTime: 0,
					MaxTime: 1,
				},
				{
					MinTime: 1,
					MaxTime: 2,
				},
				{
					MinTime: 2,
					MaxTime: 3,
				},
			},
		},
		{
			name: "sameTimes_differentRes_sameNumSamples",
			blocks: Blocks{
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 1,
				},
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 0,
				},
			},
			expectedResult: Blocks{
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 0,
				},
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 1,
				},
			},
		},
		{
			name: "sameTimes_sameRes_differentNumSamples",
			blocks: Blocks{
				{
					MinTime:    0,
					MaxTime:    1,
					NumSamples: 0,
				},
				{
					MinTime:    0,
					MaxTime:    1,
					NumSamples: 1000,
				},
			},
			expectedResult: Blocks{
				{
					MinTime:    0,
					MaxTime:    1,
					NumSamples: 1000,
				},
				{
					MinTime:    0,
					MaxTime:    1,
					NumSamples: 0,
				},
			},
		},
		{
			name: "sameTimes_differentRes_differentNumSamples",
			blocks: Blocks{
				{
					MinTime:    1,
					MaxTime:    2,
					Resolution: 0,
					NumSamples: 1,
				},
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 1000,
					NumSamples: 1,
				},
				{
					MinTime:    1,
					MaxTime:    2,
					Resolution: 1000,
					NumSamples: 1000,
				},
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 0,
					NumSamples: 1000,
				},
			},
			expectedResult: Blocks{
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 0,
					NumSamples: 1000,
				},
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 1000,
					NumSamples: 1,
				},
				{
					MinTime:    1,
					MaxTime:    2,
					Resolution: 0,
					NumSamples: 1,
				},
				{
					MinTime:    1,
					MaxTime:    2,
					Resolution: 1000,
					NumSamples: 1000,
				},
			},
		},
	}

	for _, c := range cases {
		sort.Sort(c.blocks)

		if !reflect.DeepEqual(c.blocks, c.expectedResult) {
			t.Errorf("Test case : %v failed.\nExpected : %v\nGot : %v\n", c.name, c.expectedResult, c.blocks)
		}
	}
}

func TestDedupBlocks(t *testing.T) {
	var cases = []struct {
		name           string
		blocks         Blocks
		expectedResult Blocks
	}{
		{
			name: "sameTimes_differentRes_sameNumSamples",
			blocks: Blocks{
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 1,
				},
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 0,
				},
			},
			expectedResult: Blocks{
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 0,
				},
			},
		},
		{
			name: "sameTimes_sameRes_differentNumSamples",
			blocks: Blocks{
				{
					MinTime:    0,
					MaxTime:    1,
					NumSamples: 0,
				},
				{
					MinTime:    0,
					MaxTime:    1,
					NumSamples: 1000,
				},
			},
			expectedResult: Blocks{
				{
					MinTime:    0,
					MaxTime:    1,
					NumSamples: 1000,
				},
			},
		},
		{
			name: "sameTimes_differentRes_differentNumSamples",
			blocks: Blocks{
				{
					MinTime:    1,
					MaxTime:    2,
					Resolution: 0,
					NumSamples: 1,
				},
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 1000,
					NumSamples: 1,
				},
				{
					MinTime:    1,
					MaxTime:    2,
					Resolution: 1000,
					NumSamples: 1000,
				},
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 0,
					NumSamples: 1000,
				},
			},
			expectedResult: Blocks{
				{
					MinTime:    0,
					MaxTime:    1,
					Resolution: 0,
					NumSamples: 1000,
				},
				{
					MinTime:    1,
					MaxTime:    2,
					Resolution: 0,
					NumSamples: 1,
				},
			},
		},
	}

	for _, c := range cases {
		c.blocks.DropOverlappingBlocks()

		if !reflect.DeepEqual(c.blocks, c.expectedResult) {
			t.Errorf("Test case : %v failed.\nExpected : %v\nGot : %v\n", c.name, c.expectedResult, c.blocks)
		}
	}
}
