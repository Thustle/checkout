package pricing

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
)

func LoadPrices(reader io.Reader) (*map[string]int, error) {
	csvReader := csv.NewReader(reader)
	prices := map[string]int{}
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if len(line) != 2 {
			return nil, errors.New("unexpected number of fields in line")
		}
		sku := line[0]
		price, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}
		prices[sku] = price
	}
	return &prices, nil
}
