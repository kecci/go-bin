package freebinchecker

import (
	"bytes"
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kecci/go-bin/utility"
)

var (
	ErrBinEmpty = errors.New("bin is empty")
)

type (
	BinData struct {
		Bin               string `json:"bin"`
		Prepaid           bool   `json:"prepaid"`
		Commercial        bool   `json:"commercial"`
		CardNetwork       string `json:"card_network"`
		CardType          string `json:"card_type"`
		CardLevel         string `json:"card_level"`
		Currency          string `json:"currency"`
		BankIssuerCode    string `json:"bank_issuer_code"`
		BankIssuerName    string `json:"bank_issuer_name"`
		BankIssuerWebsite string `json:"bank_issuer_website"`
		BankIssuerPhone   string `json:"bank_issuer_phone"`
		BankIssuerCity    string `json:"bank_issuer_city"`
		CountryCode       string `json:"country_code"`
		CountryName       string `json:"country_name"`
		CountryLatitude   string `json:"country_latitude"`
		CountryLongitude  string `json:"country_longitude"`
	}
)

// BinLookup returns the bin data for the given bin.
func BinLookup(bin string) (*BinData, error) {
	if bin == "" {
		return nil, ErrBinEmpty
	}

	htmlByte, err := utility.GetDataFromURL("https://www.freebinchecker.com/bin-lookup/" + bin)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(htmlByte))
	if err != nil {
		return nil, err
	}

	binData := &BinData{}

	doc.Find("table").Has("thead").Find("tbody tr").
		Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0: // Card Information
				{
					s.Find("td").Each(func(i int, s *goquery.Selection) {
						switch i {
						case 0: // bin
							binData.Bin = s.Text()
						case 2: // prepaid
							binData.Prepaid = s.Text() == "true"
						case 3: // commercial
							binData.Commercial = s.Text() == "true"
						case 4: // card network
							binData.CardNetwork = s.Text()
						case 5: // card type
							binData.CardType = s.Text()
						case 6: // card level
							binData.CardLevel = s.Text()
						case 7: // currency
							binData.Currency = s.Text()
						}
					})
				}
			case 1: // Bank Issuer
				{
					s.Find("td").Each(func(i int, s *goquery.Selection) {
						switch i {
						case 0: // bank issuer name
							binData.BankIssuerName = s.Text()
							binData.BankIssuerCode = mappingBankCode(s.Text())
						case 1: // bank issuer website
							binData.BankIssuerWebsite = s.Text()
						case 2: // bank issuer phone
							binData.BankIssuerPhone = s.Text()
						case 3: // bank issuer city
							binData.BankIssuerCity = s.Text()
						}
					})
				}
			case 2: // Country Issuer
				{
					s.Find("td").Each(func(i int, s *goquery.Selection) {
						switch i {
						case 1: // country code
							binData.CountryCode = s.Text()
						case 2: // country name
							binData.CountryName = s.Text()
						case 4: // country latitude
							binData.CountryLatitude = s.Text()
						case 5: // country longitude
							binData.CountryLongitude = s.Text()
						}
					})
				}
			}
		})

	if binData.Bin == "" {
		return nil, errors.New("bin not found")
	}

	return binData, nil
}

// mappingBankCode maps the bank code to the bank name.
func mappingBankCode(bankName string) string {

	if strings.Contains(bankName, "Bank Rakyat Indonesia") {
		return "bri"
	}

	if strings.Contains(bankName, "Bank Negara Indonesia") {
		return "bni"
	}

	if strings.Contains(bankName, "Mandiri") {
		return "mandiri"
	}

	if strings.Contains(bankName, "Permata") {
		return "permata"
	}

	if strings.Contains(bankName, "dbs") {
		return "dbs"
	}

	return ""
}
