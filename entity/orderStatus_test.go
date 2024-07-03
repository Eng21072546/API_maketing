package entity

import (
	"testing"
)

func TestUpStatus(t *testing.T) {
	tests := []struct {
		name    string
		status  Status
		want    Status
		wantErr error
	}{
		{"Change status New --> Paid", New, Paid, nil},
		{"Change status Paid --> Processing", Paid, Processing, nil},
		{"Change status Processing --> Done", Processing, Done, nil},
		{"Success status Done --> Done", Done, Done, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//assert.Equal(t, tt.want, Order.UpStatus(tt.status))
		})
	}
}
