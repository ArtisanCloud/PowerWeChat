package aibot

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/aibot/request"
	"github.com/gorilla/websocket"
)

type fakeApplication struct {
	config *kernel.Config
}

func (f *fakeApplication) GetContainer() *kernel.ServiceContainer { return nil }
func (f *fakeApplication) GetAccessToken() *kernel.AccessToken    { return nil }
func (f *fakeApplication) GetConfig() *kernel.Config              { return f.config }
func (f *fakeApplication) GetComponent(name string) interface{}   { return nil }

// Reproduce: heartbeat Send uses a ctx with deadline, leaving a write deadline
// on the connection. A later Send without deadline (common business-reply case,
// e.g. context.Background()) hits the already-expired deadline and fails with
// i/o timeout. After the fix, Send should clear the stale deadline instead.
func TestSendWithoutDeadlineClearsStaleWriteDeadline(t *testing.T) {
	up := websocket.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := up.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer c.Close()
		for {
			if _, _, err := c.ReadMessage(); err != nil {
				return
			}
		}
	}))
	defer srv.Close()

	client, err := NewClient(&fakeApplication{config: kernel.NewConfig(&object.HashMap{})})
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}
	client.Endpoint = "ws" + strings.TrimPrefix(srv.URL, "http")
	client.Dialer = &websocket.Dialer{HandshakeTimeout: 5 * time.Second}

	ctx := context.Background()
	if err := client.Connect(ctx); err != nil {
		t.Fatalf("connect failed: %v", err)
	}
	defer client.Close()

	// 1. Simulate heartbeat: send with 300ms deadline, must succeed
	hbCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	err = client.Send(hbCtx, request.NewPing("req-heartbeat"))
	cancel()
	if err != nil {
		t.Fatalf("heartbeat send failed: %v", err)
	}

	// 2. Wait past the heartbeat write deadline (simulates slow business handling)
	time.Sleep(600 * time.Millisecond)

	// 3. Simulate business reply: send without deadline, must succeed after fix;
	//    before fix the leaked expired deadline causes i/o timeout
	err = client.Send(ctx, request.NewPing("req-reply"))
	if err != nil {
		t.Fatalf("reply send without deadline failed (stale write deadline leaked by heartbeat send): %v", err)
	}
}
