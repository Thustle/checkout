package pricing

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
)

func LoadDeals(reader io.Reader) (*DealStore, error) {
	csvReader := csv.NewReader(reader)
	deals := map[string][]Deal{}
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.New("error reading deals csv file")
		}
		deal, err := convertToDeal(line)
		if err != nil {
			return nil, errors.New("error parsing deals csv file")
		}
		dealsForSku, ok := deals[deal.Sku]
		if ok {
			deals[deal.Sku] = append(dealsForSku, *deal)
		} else {
			deals[deal.Sku] = []Deal{*deal}
		}
	}
	return CreateDealStore(deals), nil
}

func convertToDeal(line []string) (*Deal, error) {
	if len(line) != 3 {
		return nil, errors.New("unexpected number of fields in line")
	}
	qty, err := strconv.Atoi(line[1])
	if err != nil {
		return nil, err
	}
	price, err := strconv.Atoi(line[2])
	if err != nil {
		return nil, err
	}
	return &Deal{
		Sku:   line[0],
		Qty:   qty,
		Price: price,
	}, nil
}
