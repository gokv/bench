package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type value string

var ErrNotFound = errors.New("key not found")

func (v *value) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*v = value(str)
	return nil
}

func (v value) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(v))
}

type store interface {
	Ping(ctx context.Context) error
	Set(ctx context.Context, k string, v json.Marshaler) error
	Get(ctx context.Context, k string, v json.Unmarshaler) (ok bool, err error)
}

func Do(s store) error {
	ctx := context.Background()
	v := value("This is an example string")

	if err := s.Ping(ctx); err != nil {
		return err
	}

	if err := s.Set(ctx, "the key", v); err != nil {
		return err
	}

	var found value
	if ok, err := s.Get(ctx, "the key", &found); err != nil || !ok {
		if !ok {
			return ErrNotFound
		}
		return err
	}

	if v != found {
		return fmt.Errorf("expected %q, found %q", v, found)
	}

	return nil
}

func main() {}
