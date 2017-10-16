package kellyframework

import (
	"net/http"
	"io"
)

type methodCallLogger struct {
	row *AccessLogRow
}

const ServiceHandlerAccessLogRowFillerContextKey = "kellyframework.ServiceHandlerAccessLogRowFiller"

func (l *methodCallLogger) Record(field string, value string) {
	l.row.SetRowField(field, value)
}

func ServiceHandlerAccessLogRowFillerFactory(row *AccessLogRow) AccessLogRowFiller {
	return &methodCallLogger{row}
}

func NewLoggingServiceMux(pairs []*PathFunctionPair, logWriter io.Writer) (http.Handler, error) {
	serviceMux := http.NewServeMux()
	err := RegisterFunctionsToServeMux(serviceMux, ServiceHandlerAccessLogRowFillerContextKey, pairs)
	if err != nil {
		return nil, err
	}

	return NewAccessLogDecorator(serviceMux, logWriter, ServiceHandlerAccessLogRowFillerContextKey,
		ServiceHandlerAccessLogRowFillerFactory), nil
}
