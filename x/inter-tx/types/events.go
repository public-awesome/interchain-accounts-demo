package types

// InterTx events
const (
	EventTypePacket  = "interchain_account_packet"
	EventTypeTimeout = "timeout"

	AttributeKeyAckSuccess     = "success"
	AttributeKeyAckError       = "error"
	AttributeKeyPacketSequence = "packet_sequence"
)
