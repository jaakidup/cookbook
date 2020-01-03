package state

import "errors"

// Process switches on operation type
// then does the work
func Process(wr *WorkRequest) *WorkResponse {
	resp := WorkResponse{Wr: wr}

	switch wr.Operation {
	case Add:
		resp.Result = wr.Value1 + wr.Value2
	case Subtract:
		resp.Result = wr.Value1 - wr.Value2
	case Multiply:
		resp.Result = wr.Value1 * wr.Value2
	case Divide:
		if wr.Value2 == 0 {
			resp.Err = errors.New("Divide by 0")
			break
		}
		resp.Result = wr.Value1 / wr.Value2
	default:
		resp.Err = errors.New("unsupported operation")
	}
	return &resp
}
