package ordersync

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDelayWithJitters(t *testing.T) {
	numCalls := 100
	approxDelay := 10 * time.Second
	jitterAmount := 0.1
	for i := 0; i < numCalls; i++ {
		actualDelay := calculateDelayWithJitter(approxDelay, jitterAmount)
		// 0.1 * 10 seconds is 1 second. So we assert that the actual delay is within 1 second
		// of the approximate delay.
		assert.InDelta(t, approxDelay, actualDelay, float64(1*time.Second), "actualDelay: %s", actualDelay)
	}
}
