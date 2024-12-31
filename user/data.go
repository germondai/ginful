package user

import "time"

var Users = []User{
	{
		ID:        "1",
		Email:     "a@a.a",
		Password:  "abc",
		CreatedAt: time.Date(2024, time.August, 25, 15, 23, 0, 0, time.UTC),
	},
	{
		ID:        "2",
		Email:     "a@a.a",
		Password:  "abc",
		CreatedAt: time.Date(2024, time.December, 12, 12, 47, 0, 0, time.UTC),
	},
}
