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

const IDX_TRACING_SAVE_INDICATOR_CACHE = "save-indicator-from-cache"

const MSG_CACHE_ERROR_TO_SAVE_INDICATOR = "Error to try to save indicator with id: %s"

type SaveIndicatorFromCacheWithTracing struct {
	indicatorCachedDatabaseGateway main_gateways.IndicatorCachedDatabaseGateway
}

func NewSaveIndicatorFromCacheWithTracing(
	indicatorCachedDatabaseGateway *main_gateways.IndicatorCachedDatabaseGateway,
) *SaveIndicatorFromCacheWithTracing {
	return &SaveIndicatorFromCacheWithTracing{
		indicatorCachedDatabaseGateway: *indicatorCachedDatabaseGateway,
	}
}

func (thisUseCase *SaveIndicatorFromCacheWithTracing) Execute(ctxP *context.Context, indicator main_domains.Indicator) main_domains.Indicator {

	tr := otel.GetTracerProvider().Tracer(main_domains.APP_INDICATOR_TYPE_FIND_INDICATOR.GetDescription())
	_, span := tr.Start(*ctxP, IDX_TRACING_SAVE_INDICATOR_CACHE, trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	indicator, err := thisUseCase.indicatorCachedDatabaseGateway.Save(indicator)
	main_utils.FailOnError(err, fmt.Sprintf(MSG_CACHE_ERROR_TO_SAVE_INDICATOR, indicator.Id))
	return indicator
}
