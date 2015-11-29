package app

// ServerMessage defines the schema for json messages sent from the server
// to the client.
type ServeMessage struct {
	ServerTime SerializableTime `json:"server_time"`
}
