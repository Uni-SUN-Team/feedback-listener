package services

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"unisun/api/feedback-processor-api/src/constants"
	"unisun/api/feedback-processor-api/src/models"
	"unisun/api/feedback-processor-api/src/ports"

	"github.com/spf13/viper"
)

type StrapiInformationServiceAdapter struct {
	Utils ports.HttpRequestPort
}

func NewConsumerServiceAdapter(utils ports.HttpRequestPort) *StrapiInformationServiceAdapter {
	return &StrapiInformationServiceAdapter{
		Utils: utils,
	}
}

func (svr *StrapiInformationServiceAdapter) GetInformationFormStrapi(payloadRequest models.ServiceIncomeRequest) (string, error) {
	var serviceIncomeResponse = models.ServiceIncomeResponse{}
	url := strings.Join([]string{viper.GetString("endpoint.strapi-information-gateway.host"), viper.GetString("endpoint.strapi-information-gateway.path")}, "")
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		return "", err
	}
	response, err := svr.Utils.HTTPRequest(url, constants.POST, payload)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	err = json.Unmarshal([]byte(body), &serviceIncomeResponse)
	if err != nil {
		return "", err
	}
	return serviceIncomeResponse.Payload, nil
}
