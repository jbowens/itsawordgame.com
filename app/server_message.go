package app

// ServerMessage defines the schema for json messages sent from the server
// to the client.
type ServerMessage struct {
	ServerTime SerializableTime `json:"server_time"`
}
