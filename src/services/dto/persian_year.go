package dto

import "time"

type CreatePersianYear struct {
	PersianTitle string
	Year         int
	StartAt      time.Time
	EndAt        time.Time
}

type UpdatePersianYear struct {
	PersianTitle string
	Year         int
	StartAt      time.Time
	EndAt        time.Time
}

type PersianYear struct {
	Id           int
	PersianTitle string
	Year         int
	StartAt      time.Time
	EndAt        time.Time
}
