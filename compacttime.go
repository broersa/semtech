package semtech

import "time"

// CompactTime implements time.Time but (un)marshals to and from
// ISO 8601 'compact' format.
type CompactTime time.Time

// MarshalJSON implements the json.Marshaler interface.
func (t CompactTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).UTC().Format(`"` + time.RFC3339Nano + `"`)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *CompactTime) UnmarshalJSON(data []byte) error {
	t2, err := time.Parse(`"`+time.RFC3339Nano+`"`, string(data))
	if err != nil {
		return err
	}
	*t = CompactTime(t2)
	return nil
}
