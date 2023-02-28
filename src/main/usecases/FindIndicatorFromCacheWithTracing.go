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

const IDX_TRACING_FIND_INDICATOR_CACHE = "find-indicator-from-cache"

const MSG_CACHE_ERROR_TO_FIND_INDICATOR = "Error to try to find indicator with id: %s"

type FindIndicatorFromCacheWithTracing struct {
	indicatorCachedDatabaseGateway main_gateways.IndicatorCachedDatabaseGateway
}

func NewFindIndicatorFromCacheWithTracing(
	indicatorCachedDatabaseGateway *main_gateways.IndicatorCachedDatabaseGateway,
) *FindIndicatorFromCacheWithTracing {
	return &FindIndicatorFromCacheWithTracing{
		indicatorCachedDatabaseGateway: *indicatorCachedDatabaseGateway,
	}
}

func (thisUseCase *FindIndicatorFromCacheWithTracing) Execute(ctxP *context.Context, id string) main_domains.Indicator {

	tr := otel.GetTracerProvider().Tracer(main_domains.APP_INDICATOR_TYPE_FIND_INDICATOR.GetDescription())
	_, span := tr.Start(*ctxP, IDX_TRACING_FIND_INDICATOR_CACHE, trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	indicator, err := thisUseCase.indicatorCachedDatabaseGateway.FindById(id)
	main_utils.FailOnError(err, fmt.Sprintf(MSG_CACHE_ERROR_TO_FIND_INDICATOR, id))
	return indicator
}
