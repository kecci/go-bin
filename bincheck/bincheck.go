package bincheck

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kecci/go-bin/utility"
)

type (
	BinData struct {
		Bin               string `json:"bin"`
		CardBrand         string `json:"card_brand"`
		CardType          string `json:"card_type"`
		CardLevel         string `json:"card_level"`
		BankIssuerName    string `json:"bank_issuer_name"`
		BankIssuerWebsite string `json:"bank_issuer_website"`
		BankIssuerPhone   string `json:"bank_issuer_phone"`
		CountryName       string `json:"country_name"`
		CountryCode       string `json:"country_code"` // Kode Negara ISO A2
		Currency          string `json:"currency"`     // IDR
	}
)

func BinDetail(bin string) (*BinData, error) {
	htmlByte, err := utility.GetDataFromURL("https://bincheck.io/id/details/" + bin)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(htmlByte))
	if err != nil {
		return nil, err
	}

	binData := &BinData{}

	index := 1

	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		s.Find("tr").Each(func(j int, t *goquery.Selection) {
			t.Find("td").Each(func(k int, u *goquery.Selection) {
				if k == 1 {
					switch index {
					case 1:
						binData.Bin = strings.TrimSpace(u.Text())
					case 2:
						binData.CardBrand = strings.TrimSpace(u.Text())
					case 3:
						binData.CardType = strings.TrimSpace(u.Text())
					case 4:
						binData.CardLevel = strings.TrimSpace(u.Text())
					case 5:
						binData.BankIssuerName = strings.TrimSpace(u.Text())
					case 6:
						binData.BankIssuerWebsite = strings.TrimSpace(u.Text())
					case 7:
						binData.BankIssuerPhone = strings.TrimSpace(u.Text())
					case 8:
						binData.CountryName = strings.TrimSpace(u.Text())
					case 10:
						binData.CountryCode = strings.TrimSpace(u.Text())
					case 12:
						binData.Currency = strings.TrimSpace(u.Text())
					}

					index++
				}
			})
		})
	})

	return binData, nil
}
