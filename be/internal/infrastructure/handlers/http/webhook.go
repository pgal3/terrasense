package emqx_webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	protobuf_adapter "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/protobuf"
	"google.golang.org/protobuf/proto"
)

type EMQXPayload struct {
	Username  string            `json:"username"`
	Topic     string            `json:"topic"`
	Timestamp int64             `json:"timestamp"`
	Peerhost  string            `json:"peerhost"`
	Payload   []byte            `json:"payload"`
	Node      string            `json:"node"`
	Metadata  map[string]string `json:"metadata"`
	ID        string            `json:"id"`
	Event     string            `json:"event"`
	ClientID  string            `json:"clientid"`
}

func StartWebhook() {
	http.HandleFunc("POST /emqx/webhook", handleWebhook)
	http.ListenAndServe(":3000", nil)
	fmt.Println("server up and running")
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("Post request received, starting elaboration")
	var payload EMQXPayload
	measurements := &protobuf_adapter.Measurements{}
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &payload)
	fmt.Println(payload)
	proto.Unmarshal(payload.Payload, measurements)
	fmt.Println(measurements)
	fmt.Println("Post request finished")
}
