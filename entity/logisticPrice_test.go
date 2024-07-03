package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogisticCost(t *testing.T) {
	tests := []struct {
		name     string
		location string
		want     float64
		wantErr  assert.ErrorAssertionFunc
	}{
		{"should return price", "international", logisticPrice["international"], assert.NoError},
		{"should return price", "domestic", logisticPrice["domestic"], assert.NoError},
		{"should not return price", "test", 0, assert.Error},
		{"should return price with lowercase", "Domestic", logisticPrice["domestic"], assert.NoError},
		{"should return price with uppercase", "INTERNATIONAL", logisticPrice["international"], assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LogisticCost(tt.location)
			if !tt.wantErr(t, err, "costLocation") {
				return
			} //??? from P'Tuk Code
			assert.Equal(t, tt.want, got)
		})
	}
}
