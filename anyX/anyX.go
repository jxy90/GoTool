package anyX

import (
	"github.com/pkg/errors"
	"strconv"
)

func AnyToString(a any) (string, error) {
	switch a.(type) {
	case string:
		return a.(string), nil
	case int:
		return strconv.FormatInt(int64(a.(int)), 10), nil
	case int8:
		return strconv.FormatInt(int64(a.(int8)), 10), nil
	case int16:
		return strconv.FormatInt(int64(a.(int16)), 10), nil
	case int32:
		return strconv.FormatInt(int64(a.(int32)), 10), nil
	case int64:
		return strconv.FormatInt(a.(int64), 10), nil
	case uint:
		return strconv.FormatUint(uint64(a.(uint)), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(a.(uint8)), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(a.(uint16)), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(a.(uint32)), 10), nil
	case uint64:
		return strconv.FormatUint(a.(uint64), 10), nil
	case float32:
		return strconv.FormatFloat(float64(a.(float32)), 'f', -1, 64), nil
	case float64:
		return strconv.FormatFloat(a.(float64), 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(a.(bool)), nil
	}

	return "", UnknownAnyType
}

func MatchAny(a, b any) (bool, error) {
	aa, err := AnyToString(a)
	if err != nil {
		return false, errors.Errorf("%+v, %v", err, a)
	}
	bb, err := AnyToString(b)
	if err != nil {
		return false, errors.Errorf("%+v, %v", err, b)
	}

	if aa == bb {
		return true, nil
	}
	return false, nil
}
