package app

// ClientMessage defines the schema for json messages sent from the client
// to the server.
type ClientMessage struct {
	IdempotencyKey string `json:"idempotency_key"`
}
