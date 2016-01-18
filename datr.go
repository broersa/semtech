package semtech

import (
	"strconv"
	"strings"
)

// DatR implements the data rate which can be either a string (LoRa identifier)
// or an unsigned integer in case of FSK (bits per second).
type DatR struct {
	LoRa string
	FSK  uint32
}

// MarshalJSON implements the json.Marshaler interface.
func (d DatR) MarshalJSON() ([]byte, error) {
	if d.LoRa != "" {
		return []byte(`"` + d.LoRa + `"`), nil
	}
	return []byte(strconv.FormatUint(uint64(d.FSK), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *DatR) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseUint(string(data), 10, 32)
	if err != nil {
		d.LoRa = strings.Trim(string(data), `"`)
		return nil
	}
	d.FSK = uint32(i)
	return nil
}
