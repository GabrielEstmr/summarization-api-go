package main_gateways_http_controllers

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"log"
	main_domains "mpindicator/main/domains"
	main_gateways_rabbitmq_producers "mpindicator/main/gateways/rabbitmq/producers"
	main_usecases "mpindicator/main/usecases"
	main_utils "mpindicator/main/utils"
	"net/http"
	"strings"
)

const IDX_TRACING_FIND_INDICATOR_CONTROLLER = "find-indicator-controller"

const IDX_HEADER_INDICATOR = "indicator"
const PATH_PREFIX = "/indicators/"

type IndicatorController struct {
	apiIndicatorTrigger                 main_usecases.ApiIndicatorTrigger
	findIndicatorByIdWithCacheAndLocked main_usecases.FindIndicatorByIdWithCacheAndLocked
}

func NewIndicatorController(useCaseApiIndicatorTrigger *main_usecases.ApiIndicatorTrigger,
	findIndicatorByIdWithCacheAndLocked *main_usecases.FindIndicatorByIdWithCacheAndLocked) *IndicatorController {
	return &IndicatorController{
		apiIndicatorTrigger:                 *useCaseApiIndicatorTrigger,
		findIndicatorByIdWithCacheAndLocked: *findIndicatorByIdWithCacheAndLocked,
	}
}

func (thisController *IndicatorController) Reprocess(w http.ResponseWriter, r *http.Request) {

	indicator := r.Header.Get(IDX_HEADER_INDICATOR)
	indicatorType, err := main_domains.FindIndicatorTypeByDescription(indicator)
	if err != nil {
		main_utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	thisController.apiIndicatorTrigger.Execute(main_domains.NewIndicatorProcessorTrigger(indicatorType))
	main_utils.JSON(w, http.StatusAccepted, indicatorType)
}

func (thisController *IndicatorController) FindIndicator(w http.ResponseWriter, r *http.Request) {

	ctx := baggage.ContextWithoutBaggage(r.Context())
	tr := otel.GetTracerProvider().Tracer(main_domains.APP_INDICATOR_TYPE_FIND_INDICATOR.GetDescription())
	ctx, controller := tr.Start(ctx, IDX_TRACING_FIND_INDICATOR_CONTROLLER)
	defer controller.End()

	//log := log.WithFields(apmlogrus.TraceContext(req.Context()))

	vars := mux.Vars(r)
	logrus.WithField("vars", vars).Info("handling hello request")

	main_gateways_rabbitmq_producers.Produce(&ctx)

	id := strings.TrimPrefix(r.URL.Path, PATH_PREFIX)
	indicator := thisController.findIndicatorByIdWithCacheAndLocked.Execute(&ctx, id)
	log.Println(id)

	main_utils.JSON(w, http.StatusAccepted, indicator)

}
