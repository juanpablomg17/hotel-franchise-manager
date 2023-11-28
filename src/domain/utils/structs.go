package utils

import (
	"github.com/mashingan/smapping"
	"github.com/pkg/errors"
)

func AdapStructs[T any, T2 any](target *T, dataSource T2) error {
	dataMapped := smapping.MapFields(dataSource)
	err := smapping.FillStruct(target, dataMapped)
	if err != nil {
		return errors.Wrap(err, "mapping error")
	}
	return nil
}
