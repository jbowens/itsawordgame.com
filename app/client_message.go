package app

// ClientMessage defines the schema for json messages sent from the client
// to the server.
type ClientMessage struct {
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	MessageType    string `json:"message_type"`
	CellID         string `json:"cell_id"`
}
