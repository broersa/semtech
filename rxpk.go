package semtech

// RXPK contain a RF packet and associated metadata.
type RXPK struct {
	Time CompactTime `json:"time"` // UTC time of pkt RX, us precision, ISO 8601 'compact' format (e.g. 2013-03-31T16:21:17.528002Z)
	Tmst uint32      `json:"tmst"` // Internal timestamp of "RX finished" event (32b unsigned)
	Freq float64     `json:"freq"` // RX central frequency in MHz (unsigned float, Hz precision)
	Chan uint8       `json:"chan"` // Concentrator "IF" channel used for RX (unsigned integer)
	RFCh uint8       `json:"rfch"` // Concentrator "RF chain" used for RX (unsigned integer)
	Stat int8        `json:"stat"` // CRC status: 1 = OK, -1 = fail, 0 = no CRC
	Modu string      `json:"modu"` // Modulation identifier "LORA" or "FSK"
	DatR DatR        `json:"datr"` // LoRa datarate identifier (eg. SF12BW500) || FSK datarate (unsigned, in bits per second)
	CodR string      `json:"codr"` // LoRa ECC coding rate identifier
	RSSI int16       `json:"rssi"` // RSSI in dBm (signed integer, 1 dB precision)
	LSNR float64     `json:"lsnr"` // Lora SNR ratio in dB (signed float, 0.1 dB precision)
	Size uint16      `json:"size"` // RF packet payload size in bytes (unsigned integer)
	Data string      `json:"data"` // Base64 encoded RF packet payload, padded
}
