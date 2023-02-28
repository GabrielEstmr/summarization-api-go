package main_usecases

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"log"
	main_domains "mpindicator/main/domains"
	main_utils "mpindicator/main/utils"
	"time"
)

const IDX_TRACING_FIND_INDICATOR_LOCK = "find-indicator-lock"

const IDX_FIND_INDICATOR_LOCK = "index-find-indicator-lock2"
const MSG_NOT_ABLE_TO_LOCK = "Not able to process due to distributed lock"
const MSG_NOT_ABLE_TO_UNLOCK = "Not able to unlock"
const MSG_NOT_ABLE_TO_EXTEND = "Not able to extend lock"

type FindIndicatorByIdWithCacheAndLocked struct {
	findIndicatorById FindIndicatorByIdWithCache
}

func NewFindIndicatorByIdWithCacheAndLocked(findIndicatorById *FindIndicatorByIdWithCache,
) *FindIndicatorByIdWithCacheAndLocked {
	return &FindIndicatorByIdWithCacheAndLocked{
		findIndicatorById: *findIndicatorById,
	}
}

func (thisUseCase *FindIndicatorByIdWithCacheAndLocked) Execute(ctxP *context.Context, id string) main_domains.Indicator {

	tr := otel.GetTracerProvider().Tracer(main_domains.APP_INDICATOR_TYPE_FIND_INDICATOR.GetDescription())
	ctx, span := tr.Start(*ctxP, IDX_TRACING_FIND_INDICATOR_LOCK, trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	singleLock := main_domains.NewSingleLock(IDX_FIND_INDICATOR_LOCK, 8*time.Second)
	err := singleLock.Lock()
	main_utils.FailOnError(err, MSG_NOT_ABLE_TO_LOCK)

	indicator := thisUseCase.findIndicatorById.Execute(&ctx, id)

	extend, err := singleLock.Extend()
	evaluateLock(extend, err, MSG_NOT_ABLE_TO_EXTEND)

	unlock, err := singleLock.Unlock()
	evaluateLock(unlock, err, MSG_NOT_ABLE_TO_UNLOCK)

	return indicator
}

func evaluateLock(result bool, err error, msg string) {
	if !result || err != nil {
		log.Printf("%s: %s", msg, err)
	}
}
