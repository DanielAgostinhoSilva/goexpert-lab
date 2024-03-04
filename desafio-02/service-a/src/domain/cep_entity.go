package domain

import "errors"

var (
	errInvalidZipCode = errors.New("invalid zipcode")
)

type Cep struct {
	numero string
}

func NewCep(numero string) (*Cep, error) {
	if len(numero) != 8 {
		return nil, errInvalidZipCode
	}
	return &Cep{numero: numero}, nil
}

func (c *Cep) Numero() string {
	return c.numero
}
