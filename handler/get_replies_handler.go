package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ameykpatil/chatbot/service"
)

// getReplyHandler is a struct holding dependencies required for GET /replies api
type getReplyHandler struct {
	replyService service.ReplyService
}

// NewGetReplyHandler creates new instance of handler
func NewGetReplyHandler(rs service.ReplyService) http.Handler {
	return &getReplyHandler{
		replyService: rs,
	}
}

// ServeHTTP serves response based on received request
func (h *getReplyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	botID := r.URL.Query().Get("bot_id")
	message := r.URL.Query().Get("message")

	if botID == "" || strings.TrimSpace(message) == "" {
		http.Error(w, "invalid query parameters", http.StatusBadRequest)
		return
	}

	reply, err := h.replyService.FindReply(ctx, botID, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{"reply": reply.Text})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
