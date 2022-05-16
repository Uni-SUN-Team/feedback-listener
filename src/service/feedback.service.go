package service

import (
	"unisun/api/feedback-processor-api/src/config"
	"unisun/api/feedback-processor-api/src/model"
)

func GetFeedbackSumAndCount(id string) model.ResultSum {
	var result model.ResultSum
	config.DB.Raw("select sum(f.rate) as sum, count(f.id) as count from feedbacks_course_links fcl join feedbacks f on f.id = fcl.feedback_id where fcl.course_id = ?", id).Scan(&result)
	return result
}
