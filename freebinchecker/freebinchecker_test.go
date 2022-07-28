package freebinchecker_test

import (
	"encoding/json"
	"testing"

	"github.com/kecci/go-bin/freebinchecker"
	"github.com/stretchr/testify/assert"
)

func TestGetBin(t *testing.T) {

	t.Run("BNI", func(t *testing.T) {
		res, err := freebinchecker.BinLookup("548988")
		assert.NoError(t, err)
		assert.NotNil(t, res)

		b, _ := json.Marshal(*res)
		t.Log(string(b))
	})

	t.Run("BRI", func(t *testing.T) {
		res, err := freebinchecker.BinLookup("518828")
		assert.NoError(t, err)
		assert.NotNil(t, res)

		b, _ := json.Marshal(*res)
		t.Log(string(b))
	})

	t.Run("DBS", func(t *testing.T) {
		res, err := freebinchecker.BinLookup("460238")
		assert.NoError(t, err)
		assert.NotNil(t, res)

		b, _ := json.Marshal(*res)
		t.Log(string(b))
	})
}
