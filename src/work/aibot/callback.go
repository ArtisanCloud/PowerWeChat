package aibot

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

// CallbackMessage contains the common fields of a verified WeCom callback.
// Raw preserves the full decrypted JSON, including message-specific and future fields.
type CallbackMessage struct {
	MsgType     string `json:"msgtype"`
	MsgID       string `json:"msgid"`
	AIBotID     string `json:"aibotid"`
	ChatID      string `json:"chatid"`
	ChatType    string `json:"chattype"`
	ResponseURL string `json:"response_url"`
	From        struct {
		UserID string `json:"userid"`
	} `json:"from"`
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Stream struct {
		ID string `json:"id"`
	} `json:"stream"`
	Event struct {
		EventType string `json:"eventtype"`
	} `json:"event"`
	Raw json.RawMessage `json:"-"`
}

// CallbackFunc returns a JSON-serializable passive reply, or nil to acknowledge receipt.
type CallbackFunc func(context.Context, *CallbackMessage) (any, error)

type callbackHandler struct {
	crypt  *kernel.Encryptor
	handle CallbackFunc
}

// NewCallbackHandler handles URL verification, incoming messages and events for
// an enterprise-created intelligent bot (empty ReceiveId).
func NewCallbackHandler(token, encodingAESKey string, handle CallbackFunc) (http.Handler, error) {
	if token == "" || len(encodingAESKey) != 43 || strings.ContainsAny(encodingAESKey, "=\r\n") || handle == nil {
		return nil, errors.New("aibot: token, 43-character EncodingAESKey and callback are required")
	}
	crypt, err := kernel.NewEncryptor("", token, encodingAESKey)
	if err != nil {
		return nil, err
	}
	return &callbackHandler{crypt: crypt, handle: handle}, nil
}

func (h *callbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ciphertext string
	query := r.URL.Query()
	switch r.Method {
	case http.MethodGet:
		ciphertext = query.Get("echostr")
	case http.MethodPost:
		body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, 10<<20))
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			http.Error(w, "callback envelope too large", http.StatusRequestEntityTooLarge)
			return
		}
		var envelope struct {
			Encrypt string `json:"encrypt"`
		}
		if err != nil || json.Unmarshal(body, &envelope) != nil {
			http.Error(w, "invalid callback envelope", http.StatusBadRequest)
			return
		}
		ciphertext = envelope.Encrypt
	default:
		w.Header().Set("Allow", "GET, POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	signature, timestamp, nonce := query.Get("msg_signature"), query.Get("timestamp"), query.Get("nonce")
	if ciphertext == "" || signature == "" || timestamp == "" || nonce == "" {
		http.Error(w, "missing callback parameters", http.StatusBadRequest)
		return
	}
	plain, cryptErr := h.crypt.OpenCallback(ciphertext, signature, nonce, timestamp)
	if cryptErr != nil {
		http.Error(w, "invalid encrypted callback", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write(plain)
		return
	}
	msg := &CallbackMessage{Raw: plain}
	if err := json.Unmarshal(plain, msg); err != nil || msg.MsgType == "" {
		http.Error(w, "invalid callback message", http.StatusBadRequest)
		return
	}
	reply, err := h.handle(r.Context(), msg)
	if err != nil {
		http.Error(w, "callback failed", http.StatusInternalServerError)
		return
	}
	if reply == nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = io.WriteString(w, "success")
		return
	}

	data, err := json.Marshal(reply)
	if err != nil {
		http.Error(w, "invalid callback reply", http.StatusInternalServerError)
		return
	}
	now := time.Now().Unix()
	sealed, cryptErr := h.crypt.EncryptEnvelope(string(data), nonce, strconv.FormatInt(now, 10))
	if cryptErr != nil {
		http.Error(w, "cannot encrypt callback reply", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(struct {
		Encrypt      string `json:"encrypt"`
		MsgSignature string `json:"msgsignature"`
		Timestamp    int64  `json:"timestamp"`
		Nonce        string `json:"nonce"`
	}{sealed.Encrypt.Value, sealed.Signature.Value, now, nonce})
}
