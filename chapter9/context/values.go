package context

import "context"

// it's considered good practice to create a type to represent keys
// then declaring all possible keys as const
type key string

const timeoutKey key = "TimeoutKey"
const deadlineKey key = "DeadlineKey"

// Setup sets some values
func Setup(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, timeoutKey, "Timeout exceeded")
	ctx = context.WithValue(ctx, deadlineKey, "Deadline exceeded")

	return ctx
}

// GetValue grabs a value given a key
// and returns a string representation of the value
func GetValue(ctx context.Context, k key) string {
	if value, ok := ctx.Value(k).(string); ok {
		return value
	}
	return ""
}
