package middleware

import (
	"fmt"
	"net/http"

	"github.com/MSEarn/go-venue-finder/internal/header"
	"github.com/MSEarn/go-venue-finder/internal/response"
	"github.com/MSEarn/go-venue-finder/logz"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewExtractHeader(log *zap.Logger, level zapcore.Level) func(*fiber.Ctx) error {

	return func(c *fiber.Ctx) error {

		tp := ""
		if tp = c.Get(header.Traceparent, ""); tp == "" {
			e := response.Err(response.HeaderRequired, fmt.Sprintf("%s header is required", header.Traceparent))
			log.Error(fmt.Sprintf("headers: %s err: %+v", c.Context().Request.Header.String(), e))
			return c.Status(http.StatusBadRequest).JSON(&e)
		}

		//parent, err := traceparent.Parse(tp)
		//if err != nil {
		//	e := response.Err(response.HeaderRequired, "invalid Traceparent format err: "+err.Error())
		//	log.Error(fmt.Sprintf("headers: %s err: %+v", c.Context().Request.Header.String(), e))
		//	return c.Status(http.StatusBadRequest).JSON(&e)
		//}
		//
		//newParent := parent.NewSpan()
		//parentID := parent.Span.SpanID.String()
		l := log.With(
		//zap.String("trace_id", newParent.Span.TraceID.String()),
		//zap.String("parent_id", parentID),
		//zap.String("span_id", newParent.Span.SpanID.String()),
		)
		logging := logz.NewLogging(l, level)

		c.Context().SetUserValue("log", logging)

		return c.Next()
	}
}
