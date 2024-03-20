package shouqianba

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	return t.Time.String()
}

// GetTime returns a std time.Time.
func (t *Timestamp) GetTime() *time.Time {
	if t == nil {
		return nil
	}
	return &t.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	fmt.Println("Timestamp UnmarshalJSON:", data)
	millis, err := strconv.ParseInt(strings.Trim(string(data), "\""), 10, 64)
	if err != nil {
		return err
	}
	t.Time = time.UnixMilli(millis)
	return nil
}
