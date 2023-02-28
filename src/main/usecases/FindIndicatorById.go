package main_usecases

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	main_domains "mpindicator/main/domains"
	main_gateways "mpindicator/main/gateways"
	main_utils "mpindicator/main/utils"
)

const MSG_ERROR_TO_FIND_INDICATOR = "Error to try to find indicator with id: %s"

const IDX_TRACING_FIND_INDICATOR_DATABASE_GATEWAY = "find-indicator-database-gateway"

type FindIndicatorById struct {
	indicatorDatabaseGateway main_gateways.IndicatorDatabaseGateway
}

func NewFindIndicatorById(indicatorDatabaseGateway *main_gateways.IndicatorDatabaseGateway) *FindIndicatorById {
	return &FindIndicatorById{indicatorDatabaseGateway: *indicatorDatabaseGateway}
}

func (thisUseCase *FindIndicatorById) Execute(ctxP *context.Context, id string) main_domains.Indicator {
	tr := otel.GetTracerProvider().Tracer(main_domains.APP_INDICATOR_TYPE_FIND_INDICATOR.GetDescription())
	ctx, span := tr.Start(*ctxP, IDX_TRACING_FIND_INDICATOR_DATABASE_GATEWAY, trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	result, err := thisUseCase.indicatorDatabaseGateway.FindById(ctx, id)
	main_utils.FailOnError(err, fmt.Sprintf(MSG_ERROR_TO_FIND_INDICATOR, id))
	return result
}
