package company_service

import "github.com/aglyzov/go-patch"

func HasChanges[T any](oldData, newData *T) (bool, error) {
	changed, err := patch.Struct(oldData, newData)
	if err != nil {
		return false, err
	}
	return changed, nil
}
