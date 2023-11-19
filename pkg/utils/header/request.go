package header

import (
	"os"

	"github.com/google/uuid"

	"github.com/sah4ez/dapr-example/internal/transport"
)

const (
	RequestHeader  = "X-Request-Id"
	UserHeader     = "X-User-Id"
	InstanceHeader = "X-Instance-Id"
)

type TransportOpt func(h transport.Header, value string)

func WithResponseHeader() TransportOpt {
	return func(h transport.Header, value string) {
		h.ResponseKey = RequestHeader
		h.ResponseValue = value
	}
}

func RequestID(opts ...TransportOpt) func(value string) transport.Header {

	return func(value string) (header transport.Header) {
		if value == "" {
			value = uuid.New().String()
		}
		h := transport.Header{
			LogKey:    "requestID",
			LogValue:  value,
			SpanKey:   "requestID",
			SpanValue: value,
		}

		for _, o := range opts {
			o(h, value)
		}

		return h
	}
}

func Hostname(value string) (header transport.Header) {

	if value == "" {
		value, _ = os.Hostname()
	}
	h := transport.Header{
		ResponseKey:   InstanceHeader,
		ResponseValue: value,
	}
	return h
}

func UserID(value string) (header transport.Header) {
	if value == "" {
		value = "empty"
	}
	h := transport.Header{
		LogKey:    "userID",
		LogValue:  value,
		SpanKey:   "userID",
		SpanValue: value,
	}

	return h
}
