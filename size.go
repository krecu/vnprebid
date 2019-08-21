package vnprebid

import (
	"errors"

	"encoding/json"
)

var (
	ErrInvalidSizeW = errors.New("vnprebid: bad size width")
	ErrInvalidSizeH = errors.New("vnprebid: bad size height")
	ErrInvalidSizeF = errors.New("vnprebid: bad size format")
)

type Item struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type Size struct {
	Items []Item
}

// Validates the `size` object
func (s *Size) Validate() error {
	if len(s.Items) == 0 {
		return ErrInvalidSizeW
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaler
func (s *Size) UnmarshalJSON(data []byte) (err error) {

	var stuff interface{}
	err = json.Unmarshal(data, &stuff)
	if err != nil {
		return err
	}
	switch t := stuff.(type) {
	case []interface{}:

		for _, value := range t {

			if _, ok := value.([]interface{}); !ok {
				return ErrInvalidSizeF
			}

			s.Items = append(s.Items, Item{
				Height: ToInt(value.([]interface{})[0]),
				Width:  ToInt(value.([]interface{})[1]),
			})
		}
	default:

	}

	return nil
}

func ToInt(i interface{}) int {
	if i == nil {
		return 0
	}
	switch i2 := i.(type) {
	default:
		return 0
	case *json.Number:
		i3, _ := i2.Int64()
		return int(i3)
	case json.Number:
		i3, _ := i2.Int64()
		return int(i3)
	case int64:
		return int(i2)
	case float64:
		return int(i2)
	case float32:
		return int(i2)
	case uint64:
		return int(i2)
	case int:
		return int(i2)
	case uint:
		return int(i2)
	case bool:
		if i2 {
			return 1
		} else {
			return 0
		}
	case *bool:
		if i2 == nil {
			return 0
		}
		if *i2 {
			return 1
		} else {
			return 0
		}
	}
	return 0
}
