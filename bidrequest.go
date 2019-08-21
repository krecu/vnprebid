package vnprebid

import "errors"

var (
	ErrInvalidBidNoID   = errors.New("vnprebid: bid is missing ID")
	ErrInvalidBidNoCpm  = errors.New("vnprebid: bid is missing cpm")
	ErrInvalidBidNoCurr = errors.New("vnprebid: bid is missing currency")
	ErrInvalidBidNoSize = errors.New("vnprebid: bid is missing size")
)

type BidRequest struct {
	ID        string  `json:"id"`            // идентификатор для рекламного места в терминах Яндекса (из запроса)
	Cpm       float64 `json:"cpm"`           // ставка, число, больше нуля
	Currency  string  `json:"cur,omitempty"` // валюта ставки 'RUB'
	Size      *Size   `json:"sizes,omitempty"`
	Referer   string  `json:"ref,omitempty"`
	Placement string  `json:"placementId,omitempty"`
}

// Validate required attributes
func (bid *BidRequest) Validate() error {
	if bid.ID == "" {
		return ErrInvalidBidNoID
	} else if bid.Cpm <= 0 {
		return ErrInvalidBidNoCpm
	} else if bid.Currency == "" {
		return ErrInvalidBidNoCurr
	} else if bid.Size == nil {
		return ErrInvalidBidNoSize
	}

	if err := bid.Size.Validate(); err != nil {
		return err
	}

	return nil
}
