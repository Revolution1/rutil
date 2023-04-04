package must

import "reflect"

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Two[T any](val T, err error) T {
	Must(err)
	return val
}

func Three[T any, P any](val1 T, val2 P, err error) (T, P) {
	Must(err)
	return val1, val2
}

func Four[T any, P any, Q any](val1 T, val2 P, val3 Q, err error) (T, P, Q) {
	Must(err)
	return val1, val2, val3
}

func Default[T any](value, _default T) T {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Slice:
		if !v.IsNil() && v.Len() == 0 {
			return _default
		}
	}
	if v.IsZero() {
		return _default
	}
	return value
}
