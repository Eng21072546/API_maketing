package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder_UpStatus(t *testing.T) {
	tests := []struct {
		name     string
		status   Status
		expected Status
		err      error
	}{
		{name: "should new --> paid", status: New, expected: Paid, err: nil},
		{name: "should paid --> processing", status: Paid, expected: Processing, err: nil},
		{name: "should processing --> done", status: Processing, expected: Done, err: nil},
		{name: "should done", status: Done, expected: Done, err: nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			statusUpdated := Order.UpStatus(Order{}, test.status)
			assert.Equal(t, test.expected, statusUpdated)
		})
	}

}
