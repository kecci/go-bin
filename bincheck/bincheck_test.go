package bincheck_test

import (
	"encoding/json"
	"testing"

	"github.com/kecci/go-bin/bincheck"
	"github.com/stretchr/testify/assert"
)

func TestBinDetail(t *testing.T) {
	t.Run("BNI", func(t *testing.T) {
		binData, err := bincheck.BinDetail("548988")
		assert.NoError(t, err)
		assert.NotNil(t, binData)

		b, _ := json.Marshal(*binData)
		t.Log(string(b))
	})

	t.Run("AMERICAN EXPRESS", func(t *testing.T) {
		binData, err := bincheck.BinDetail("370192")
		assert.NoError(t, err)
		assert.NotNil(t, binData)

		b, _ := json.Marshal(*binData)
		t.Log(string(b))
	})
}
