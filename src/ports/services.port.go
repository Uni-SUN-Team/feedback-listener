package ports

import "unisun/api/feedback-processor-api/src/models"

type ServicesFeedbackPort interface {
	GetInformationFormStrapi(payloadRequest models.ServiceIncomeRequest) (string, error)
}
