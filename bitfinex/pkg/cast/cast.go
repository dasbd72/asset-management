package cast

import (
	"fmt"
	"strconv"
)

type (
	NilOrString struct {
		value interface{}
	}

	NilOrFloat64 struct {
		value interface{}
	}

	NilOrInt struct {
		value interface{}
	}
)

// NilOrString is a helper type to convert nil or string

func IfToNilOrString(value interface{}) NilOrString {
	return NilOrString{value: value}
}

func (v NilOrString) String() string {
	if v.value == nil {
		return ""
	}
	if _, ok := v.value.(string); !ok {
		return ""
	}
	return v.value.(string)
}

func (v NilOrString) Value() interface{} {
	return v.value
}

func (v NilOrString) IsNil() bool {
	return v.value == nil
}

func (v NilOrString) Valid() bool {
	if v.value == nil {
		return true
	}
	if _, ok := v.value.(string); !ok {
		return false
	}
	return true
}

func (v NilOrString) MarshalJSON() ([]byte, error) {
	if v.value == nil {
		return []byte("null"), nil
	}
	// check type
	_, ok := v.value.(string)
	if !ok {
		return nil, fmt.Errorf("value is not string type: %v", v.value)
	}
	return []byte(`"` + v.value.(string) + `"`), nil
}

// NilOrFloat64 is a helper type to convert nil or float64

func IfToNilOrFloat64(value interface{}) NilOrFloat64 {
	return NilOrFloat64{value: value}
}

func (v NilOrFloat64) Float64() float64 {
	if v.value == nil {
		return 0
	}
	if _, ok := v.value.(float64); !ok {
		return 0
	}
	return v.value.(float64)
}

func (v NilOrFloat64) Value() interface{} {
	return v.value
}

func (v NilOrFloat64) IsNil() bool {
	return v.value == nil
}

func (v NilOrFloat64) Valid() bool {
	if v.value == nil {
		return true
	}
	if _, ok := v.value.(float64); !ok {
		return false
	}
	return true
}

func (v NilOrFloat64) MarshalJSON() ([]byte, error) {
	if v.value == nil {
		return []byte("null"), nil
	}
	// check type
	_, ok := v.value.(float64)
	if !ok {
		return nil, fmt.Errorf("value is not float64 type: %v", v.value)
	}
	return []byte(strconv.FormatFloat(v.value.(float64), 'f', -1, 64)), nil
}

// NilOrInt is a helper type to convert nil or int

func IfToNilOrInt(value interface{}) NilOrInt {
	return NilOrInt{value: value}
}

func (v NilOrInt) Int() int {
	if v.value == nil {
		return 0
	}
	if _, ok := v.value.(int); !ok {
		return 0
	}
	return v.value.(int)
}

func (v NilOrInt) Value() interface{} {
	return v.value
}

func (v NilOrInt) IsNil() bool {
	return v.value == nil
}

func (v NilOrInt) Valid() bool {
	if v.value == nil {
		return true
	}
	if _, ok := v.value.(int); !ok {
		return false
	}
	return true
}

func (v NilOrInt) MarshalJSON() ([]byte, error) {
	if v.value == nil {
		return []byte("null"), nil
	}
	// check type
	_, ok := v.value.(int)
	if !ok {
		return nil, fmt.Errorf("value is not int type: %v", v.value)
	}
	return []byte(strconv.FormatInt(int64(v.value.(int)), 10)), nil
}
