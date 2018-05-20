package portal

type ConnectionType int

const (
// CONNECTION_STANDARD
// CONNECTION_BARRIER_ANCHOR_TO_EXTENT
// CONNECTION_BOX_DROPPER
// CONNECTION_CATAPULT_TARGET

)

type Connection struct {
	Sender   int
	Receiver int
	Type     ConnectionType
}
