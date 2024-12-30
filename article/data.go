package article

import "time"

var Articles = []Article{
	{
		ID:        "1",
		Title:     "Article 1",
		Content:   "Article 1 content",
		CreatedAt: time.Date(2024, time.August, 25, 15, 23, 0, 0, time.UTC),
	},
	{
		ID:        "2",
		Title:     "Article 2",
		Content:   "Article 2 content",
		CreatedAt: time.Date(2024, time.December, 12, 12, 47, 0, 0, time.UTC),
	},
}
