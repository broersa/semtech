package semtech

// TXPK contains a RF packet to be emitted and associated metadata.
type TXPK struct {
	Imme bool         `json:"imme"`           // Send packet immediately (will ignore tmst & time)
	Tmst uint32       `json:"tmst,omitempty"` // Send packet on a certain timestamp value (will ignore time)
	Time *CompactTime `json:"time,omitempty"` // Send packet at a certain time (GPS synchronization required)
	Freq float64      `json:"freq"`           // TX central frequency in MHz (unsigned float, Hz precision)
	RFCh uint8        `json:"rfch"`           // Concentrator "RF chain" used for TX (unsigned integer)
	Powe uint8        `json:"powe"`           // TX output power in dBm (unsigned integer, dBm precision)
	Modu string       `json:"modu"`           // Modulation identifier "LORA" or "FSK"
	DatR DatR         `json:"datr"`           // LoRa datarate identifier (eg. SF12BW500) || FSK datarate (unsigned, in bits per second)
	CodR string       `json:"codr,omitempty"` // LoRa ECC coding rate identifier
	FDev uint16       `json:"fdev,omitempty"` // FSK frequency deviation (unsigned integer, in Hz)
	IPol bool         `json:"ipol"`           // Lora modulation polarization inversion
	Prea uint16       `json:"prea,omitempty"` // RF preamble size (unsigned integer)
	Size uint16       `json:"size"`           // RF packet payload size in bytes (unsigned integer)
	NCRC bool         `json:"ncrc,omitempty"` // If true, disable the CRC of the physical layer (optional)
	Data string       `json:"data"`           // Base64 encoded RF packet payload, padding optional
}
