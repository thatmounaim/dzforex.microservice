package exchange

import (
	"context"
	"errors"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/thatmounaim/dzforex.microservice/internal/storage"
	"github.com/thatmounaim/dzforex.microservice/pkg/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ExchangeService struct {
	proto.UnimplementedDzForexServer
	sc *Scrapper
	ds storage.Storage
	l  hclog.Logger
}

func NewExchangeService(sc *Scrapper, ds storage.Storage, l hclog.Logger) *ExchangeService {
	return &ExchangeService{
		sc: sc,
		ds: ds,
		l:  l,
	}
}

func (s *ExchangeService) UpdateData() {
	data, err := s.sc.GetLatestExchangeRates()
	if err != nil {
		s.l.Error("could not update upstream exchange rates")
	}
	s.ds.UpdateData(data)
}

func (s *ExchangeService) GetRate(ctx context.Context, r *proto.RateRequest) (*proto.RateResponse, error) {
	cb := strings.ToLower(r.Currency) + "_buy"
	cs := strings.ToLower(r.Currency) + "_sell"
	rb, err := s.ds.Get(cb)

	if err != nil {
		s.l.Error("could not find buy price for", r.Currency)
		return nil, errors.New("could not find buy price for requested currency")
	}

	rs, err := s.ds.Get(cs)
	if err != nil {
		s.l.Error("could not find sell price for", r.Currency)
		return nil, errors.New("could not find sell price for requested currency")
	}

	return &proto.RateResponse{
		Buy:  rb,
		Sell: rs,
	}, nil
}

func (s *ExchangeService) GetAvailableCurrencies(ctx context.Context, r *emptypb.Empty) (*proto.AvailableCurrenciesResponse, error) {
	d := s.ds.GetAll()
	// NOTE: Quick and Dirty Duplicate, Suffix Removal
	keys := make([]string, 0, len(d)/2)
	pr := make(map[string]bool)
	for k := range d {
		kk := strings.TrimRight(k, "_sell")
		kk = strings.TrimRight(kk, "_buy")
		_, ok := pr[kk]
		if !ok {
			keys = append(keys, kk)
		}
	}
	return &proto.AvailableCurrenciesResponse{
		Currencies: keys,
	}, nil
}
