package app

import (
	"fmt"
	"time"
)

// SerializableTime wraps a time.Time but marshals it as a RFC3339 timestamp.
type SerializableTime time.Time

func (t SerializableTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))), nil
}
