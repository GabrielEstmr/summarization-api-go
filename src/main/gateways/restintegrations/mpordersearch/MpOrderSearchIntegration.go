package main_gateways_restintegrations_mpordersearch

import (
	"encoding/json"
	"go.elastic.co/apm/module/apmhttp"
	"io"
	"log"
	main_configurations_yml "mpindicator/main/configurations/yml"
	main_gateways_restintegrations_mpordersearch_responses "mpindicator/main/gateways/restintegrations/mpordersearch/responses"
	main_utils "mpindicator/main/utils"
	"net/http"
	"net/url"
	"strconv"
)

const MSG_MP_ORDER_SEARCH_INTEGRATION_REQUEST_PARAMS = "Integration with mp-order-search has been made. URL: %s"

const MP_ORDER_SEARCH_INTEGRATION_URI_INDEX = "Integration.mp-order-search.url"

type MpOrderSearchIntegration struct {
}

func NewMpOrderSearchIntegration() *MpOrderSearchIntegration {
	return &MpOrderSearchIntegration{}
}

func (thisGateway *MpOrderSearchIntegration) GetOrder(
	orderBeginDate string,
	orderEndDate string,
	invoiceStatus string,
	page int64,
	pageSize int64,
	sellerId string,
	orderStatusList []string) main_gateways_restintegrations_mpordersearch_responses.OrderSearchPageResponse {

	requestURI := main_configurations_yml.GetBeanPropertyByName(MP_ORDER_SEARCH_INTEGRATION_URI_INDEX)

	// TODO: isolate this dependency from gateway
	client := apmhttp.WrapClient(&http.Client{})

	URLQueryPart := make(url.Values)

	URLQueryPart.Add("orderBeginDate", orderBeginDate)
	URLQueryPart.Add("orderEndDate", orderEndDate)
	if invoiceStatus != "" {
		URLQueryPart.Add("invoiceStatus", invoiceStatus)
	}
	URLQueryPart.Add("page", strconv.FormatInt(page, 10))
	URLQueryPart.Add("pageSize", strconv.FormatInt(pageSize, 10))
	if sellerId != "" {
		URLQueryPart.Add("sellerId", sellerId)
	}
	for _, value := range orderStatusList {
		URLQueryPart.Add("orderStatus", value)
	}

	fullURL := requestURI + "/v1/orders" + "?" + URLQueryPart.Encode()
	log.Printf(MSG_MP_ORDER_SEARCH_INTEGRATION_REQUEST_PARAMS, fullURL)

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	main_utils.FailOnError(err, "err.Error()")

	resp, err := client.Do(req)
	main_utils.FailOnError(err, "err.Error()")
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	main_utils.FailOnError(err, "err.Error()")

	var responseObject main_gateways_restintegrations_mpordersearch_responses.OrderSearchPageResponse
	errorJSON := json.Unmarshal(bodyBytes, &responseObject)
	main_utils.FailOnError(errorJSON, "errorJSON.Error()")

	return responseObject

}
