package cakeday

import (
	"testing"
)

func TestGet(t *testing.T) {
	username := "darthlukan"
	_, err := Get(username)
	if err != nil {
		t.Errorf("Get(%v) err = %v, want nil\n", username, err)
	}
}
