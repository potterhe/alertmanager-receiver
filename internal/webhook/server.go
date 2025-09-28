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
	slog.Info("Received alert:", "body", string(alertBytes))

	msg := &am_webhook.Message{}
	if err := json.Unmarshal(alertBytes, msg); err != nil {
		slog.Error("Failed to parse webhook message:", "err", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO Handle the webhook message.
	slog.Info("Handling alert:", "body", msg)
	if err := rh.processNotifyMessage(msg); err != nil {
		slog.Error("Failed to process webhook message:", "err", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (rh *ReceiverHandler) processNotifyMessage(msg *am_webhook.Message) error {
	// TODO:
	return nil
}
