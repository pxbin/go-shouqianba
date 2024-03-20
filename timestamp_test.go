package shouqianba

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimestamp_GetTime(t *testing.T) {
	type A struct {
		Timestamp Timestamp `json:"finish_time"`
	}

	var str = `{"finish_time": "1710844025035"}`

	a := A{}
	err := json.Unmarshal([]byte(str), &a)
	assert.Nil(t, err)
	assert.Equal(t, time.UnixMilli(1710844025035), a.Timestamp.Time)
}
