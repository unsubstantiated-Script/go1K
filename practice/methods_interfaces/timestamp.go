package methods_interfaces

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) String() string {
	if ts.IsZero() {
		return "unknown"
	}

	const layout = "2006/01"

	return ts.Format(layout)
}

func (ts timestamp) MarshalJSON() (data []byte, _ error) {
	return strconv.AppendInt(data, ts.Unix(), 10), nil
}

func (ts *timestamp) UnmarshalJSON(data []byte) error {
	*ts = toTimestamp(string(data))
	return nil
}

func toTimestamp(v interface{}) (ts timestamp) {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}
