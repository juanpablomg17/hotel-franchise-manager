package utils

import (
	"encoding/json"
	"fmt"
	"reflect"

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

func StructToMap(src any) (map[string]interface{}, error) {
	var target map[string]interface{}
	data, _ := json.Marshal(src)

	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	return target, nil
}

func PrintStructInfoWithoutZeroValues[T interface{}](obj T) string {
	str := "["
	objetctMap, _ := StructToMap(obj)

	for k, v := range objetctMap {
		if v == nil {
			continue
		}
		if !reflect.ValueOf(objetctMap[k]).IsZero() && reflect.DeepEqual(objetctMap[k], v) {
			str += fmt.Sprintf("%s: %v, ", k, v)
		}
	}
	str += "]"
	return str
}
