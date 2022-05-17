package model

type FeedbacksCourseLinks struct {
	FeedbackId int
	CourseId   int
}

type ResultSum struct {
	Sum   int8    `json:"sum"`
	Count int8    `json:"count"`
	Score float32 `json:"Score"`
}
