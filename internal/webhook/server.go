package webhook

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	am_webhook "github.com/prometheus/alertmanager/notify/webhook"
)

type ReceiverHandler struct{}

func (rh *ReceiverHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	slog.Info("Received request:", "req", req)

	// Read request body.
	alertBytes, err := io.ReadAll(req.Body)
	if err != nil {
		slog.Error("Failed to read request body:", "err", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	msg := &am_webhook.Message{}
	if err := json.Unmarshal(alertBytes, msg); err != nil {
		slog.Error("Failed to parse webhook message:", "err", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO Handle the webhook message.
	slog.Info("Handling alert:", "body", msg)
	rw.WriteHeader(http.StatusOK)
}
