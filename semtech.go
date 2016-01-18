package semtech

import "errors"

// PacketType defines the packet type.
type PacketType byte

// Available packet types
const (
	PushData PacketType = iota
	PushACK
	PullData
	PullResp
	PullACK
)

// Protocol versions
const (
	ProtocolVersion1 uint8 = 0x01
)

// GetPacketType returns the packet type for the given packet data.
func GetPacketType(data []byte) (PacketType, error) {
	if len(data) < 4 {
		return PacketType(0), errors.New("lorawan/semtech: at least 4 bytes of data are expected")
	}

	if data[0] != ProtocolVersion1 {
		return PacketType(0), errors.New("lorawan/semtech: unknown protocol version")
	}

	return PacketType(data[3]), nil
}
