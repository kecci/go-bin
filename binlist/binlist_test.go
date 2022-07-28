package binlist_test

import (
	"encoding/json"
	"testing"

	"github.com/kecci/go-bin/binlist"
	"github.com/stretchr/testify/assert"
)

func TestBinLookup(t *testing.T) {
	t.Run("BNI", func(t *testing.T) {
		res, err := binlist.BinLookup("548988")
		assert.NoError(t, err)
		assert.NotNil(t, res)

		b, _ := json.Marshal(*res)
		t.Log(string(b))
	})

	t.Run("BRI", func(t *testing.T) {
		res, err := binlist.BinLookup("518828")
		assert.NoError(t, err)
		assert.NotNil(t, res)

		b, _ := json.Marshal(*res)
		t.Log(string(b))
	})

	t.Run("DBS", func(t *testing.T) {
		res, err := binlist.BinLookup("460238")
		assert.NoError(t, err)
		assert.NotNil(t, res)

		b, _ := json.Marshal(*res)
		t.Log(string(b))
	})
}
