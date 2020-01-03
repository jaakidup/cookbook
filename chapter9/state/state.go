package state

type op string

const (
	// Add values
	Add op = "add"
	// Subtract values
	Subtract op = "sub"
	// Multiply values
	Multiply op = "mult"
	// Divide values
	Divide op = "div"
)

// WorkRequest perform an op on two values
type WorkRequest struct {
	Operation op
	Value1    int64
	Value2    int64
}

// WorkResponse returns the result and any errors
type WorkResponse struct {
	Wr     *WorkRequest
	Result int64
	Err    error
}
