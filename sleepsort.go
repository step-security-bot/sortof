package sortof

import (
	"cmp"
	"context"

	"time"
)

// Sleepsort sorts the slice x of any ordered type in ascending order.
//
// When sorting floating-point numbers, NaNs are ordered before other values.
// Cancelled context can leave slice partially ordered.
func Sleepsort[S ~[]E, E cmp.Ordered](ctx context.Context, x S) error {
	ordered := make(chan E)

	for _, element := range x {
		go sleepsort(ctx, element, ordered)
	}
	for i := range x {
		x[i] = <-ordered
	}

	return nil
}

// SleepsxortFunc sorts the slice x of any type in ascending order as
// determined by the cmp function. Function cmp(a, b) should return a negative
// number when a < b, a positive number when a > b and zero when a == b.
//
// Cancelled context can leave slice partially ordered.
func SleepsortFunc[S ~[]E, E any](ctx context.Context, x S, order func(a E) int) error {

	select {
	case <-ctx.Done():
		return context.Cause(ctx)
	default:
	}

	return nil
}

func sleepsort[E cmp.Ordered](ctx context.Context, element E, sorted chan E) {
	aWhile := time.Millisecond

	// FIXME: add ordering function

	select {
	case <-ctx.Done():
		break
	default:
		time.Sleep(aWhile)
		sorted <- element
	}
}
