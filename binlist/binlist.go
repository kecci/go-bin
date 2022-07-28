package binlist

import (
	"encoding/json"
	"errors"

	"github.com/kecci/go-bin/utility"
)

var (
	ErrBinEmpty = errors.New("bin is empty")
)

type (
	BinData struct {
		Number struct {
			Length int  `json:"length"`
			Luhn   bool `json:"luhn"`
		} `json:"number"`
		Scheme  string `json:"scheme"`
		Type    string `json:"type"`
		Brand   string `json:"brand"`
		Prepaid bool   `json:"prepaid"`
		Country struct {
			Numeric   string `json:"numeric"`
			Alpha2    string `json:"alpha2"`
			Name      string `json:"name"`
			Emoji     string `json:"emoji"`
			Currency  string `json:"currency"`
			Latitude  int    `json:"latitude"`
			Longitude int    `json:"longitude"`
		} `json:"country"`
		Bank struct {
			Name  string `json:"name"`
			URL   string `json:"url"`
			Phone string `json:"phone"`
		} `json:"bank"`
	}
)

// BinLookup returns the bin data for the given bin.
func BinLookup(bin string) (*BinData, error) {
	if bin == "" {
		return nil, ErrBinEmpty
	}

	dataByte, err := utility.GetDataFromURL("https://lookup.binlist.net/" + bin)
	if err != nil {
		return nil, err
	}

	binData := BinData{}
	err = json.Unmarshal(dataByte, &binData)
	if err != nil {
		return nil, err
	}

	return &binData, nil
}
