package main

import (
	"reflect"
	"testing"
)

func TestCreateBerlinClock(t *testing.T) {

	cases := []struct {
		name           string
		twentyFourHour string
		secondBulb     bool
		fiveHourRow    int
		oneHourRow     int
		fiveMinuteRow  int
		oneMinuteRow   int
		causesError    bool
	}{
		{
			name:           "Check all rows 0",
			twentyFourHour: "00:00:00",
			secondBulb:     true,
			fiveHourRow:    0,
			oneHourRow:     0,
			fiveMinuteRow:  0,
			oneMinuteRow:   0,
			causesError:    false,
		},
		{
			name:           "Check 6am",
			twentyFourHour: "06:00:00",
			secondBulb:     true,
			fiveHourRow:    1,
			oneHourRow:     1,
			fiveMinuteRow:  0,
			oneMinuteRow:   0,
			causesError:    false,
		},
		{
			name:           "Check 13:35:51",
			twentyFourHour: "13:35:51",
			secondBulb:     false,
			fiveHourRow:    2,
			oneHourRow:     3,
			fiveMinuteRow:  7,
			oneMinuteRow:   0,
			causesError:    false,
		},
		{
			name:           "Check 23:59:59",
			twentyFourHour: "23:59:59",
			secondBulb:     false,
			fiveHourRow:    4,
			oneHourRow:     3,
			fiveMinuteRow:  11,
			oneMinuteRow:   4,
			causesError:    false,
		},
		{
			name:           "Check 24:00:01 causes error",
			twentyFourHour: "24:00:01",
			causesError:    true,
		},
		{
			name:           "Check 12:61:01 causes error",
			twentyFourHour: "12:61:01",
			causesError:    true,
		},
		{
			name:           "Check 06:59:61 causes error",
			twentyFourHour: "06:59:61",
			causesError:    true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			var causesError bool
			got, err := createBerlinClock(c.twentyFourHour)
			if err != nil {
				causesError = true
			}

			want := &berlinClock{
				SecondBulb:    c.secondBulb,
				FiveHourRow:   c.fiveHourRow,
				OneHourRow:    c.oneHourRow,
				FiveMinuteRow: c.fiveMinuteRow,
				OneMinuteRow:  c.oneMinuteRow,
			}

			if causesError {
				if !reflect.DeepEqual(causesError, c.causesError) {
					t.Errorf("got: %#v\nwant: %#v\n", causesError, c.causesError)
				}
			} else {

				if !reflect.DeepEqual(got, want) {
					t.Errorf("got: %#v\nwant: %#v\n", got, want)
				}
			}

		})
	}
}
