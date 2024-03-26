package shouqianba

import (
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
	str := strings.Trim(string(data), "\"")
	if str == "" {
		return nil
	}

	millis, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	t.Time = time.UnixMilli(millis)
	return nil
}
