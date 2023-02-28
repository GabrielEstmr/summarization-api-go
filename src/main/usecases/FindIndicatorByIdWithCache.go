package main_usecases

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"log"
	main_domains "mpindicator/main/domains"
)

const IDX_TRACING_FIND_INDICATOR_CACHE_MANAGEMENT = "find-indicator-cache-management"

const MSG_FIND_INDICATOR_FROM_DATABASE = "Indicator From cache not found: Find indicator from database"
const MSG_FIND_INDICATOR_FROM_CACHE = "Find indicator from cache"

type FindIndicatorByIdWithCache struct {
	findIndicatorById   FindIndicatorById
	findIndicatorCached FindIndicatorFromCacheWithTracing
	saveIndicatorCached SaveIndicatorFromCacheWithTracing
}

func NewFindIndicatorByIdWithCache(
	findIndicatorById *FindIndicatorById,
	findIndicatorCached *FindIndicatorFromCacheWithTracing,
	saveIndicatorCached *SaveIndicatorFromCacheWithTracing,
) *FindIndicatorByIdWithCache {
	return &FindIndicatorByIdWithCache{
		findIndicatorById:   *findIndicatorById,
		findIndicatorCached: *findIndicatorCached,
		saveIndicatorCached: *saveIndicatorCached,
	}
}

func (thisUseCase *FindIndicatorByIdWithCache) Execute(ctxP *context.Context, id string) main_domains.Indicator {

	tr := otel.GetTracerProvider().Tracer(main_domains.APP_INDICATOR_TYPE_FIND_INDICATOR.GetDescription())
	ctx, span := tr.Start(*ctxP, IDX_TRACING_FIND_INDICATOR_CACHE_MANAGEMENT, trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	indicator := thisUseCase.findIndicatorCached.Execute(&ctx, id)
	log.Println(MSG_FIND_INDICATOR_FROM_CACHE)
	if indicator.Id == "" {
		log.Println(MSG_FIND_INDICATOR_FROM_DATABASE)
		notCachedIndicator := thisUseCase.findIndicatorById.Execute(&ctx, id)
		indicatorFromCache := thisUseCase.saveIndicatorCached.Execute(&ctx, notCachedIndicator)
		return indicatorFromCache
	}
	return indicator
}
