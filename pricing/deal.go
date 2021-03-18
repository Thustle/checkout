package pricing

import (
	"errors"
	"sort"
)

type Deal struct {
	Sku string
	Qty int
	Price int
}

type SortedDeals []Deal

func (sd SortedDeals) Len() int           { return len(sd) }
func (sd SortedDeals) Less(i, j int) bool { return sd[i].Qty > sd[j].Qty }
func (sd SortedDeals) Swap(i, j int)      { sd[i], sd[j] = sd[j], sd[i] }

type DealStore struct {
	deals map[string][]Deal
}

func CreateDealStore(deals map[string][]Deal) *DealStore {
	for _, skuDeals := range deals {
		sort.Sort(SortedDeals(skuDeals))
	}
	return &DealStore{deals}
}

func (ds *DealStore) GetBestDealForSkuQty(sku string, qty int) (*Deal, error) {
	for _, deal := range ds.deals[sku] {
		if deal.Qty <= qty {
			return &deal, nil
		}
	}
	return nil, errors.New("no deal found")
}
