package cakeday

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	username := "darthlukan"
	day, err := Get(username)
	if err != nil {
		t.Errorf("Get(%v) err = %v, want nil\n", username, err)
	}
}
