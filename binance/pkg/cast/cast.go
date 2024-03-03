package cast

import (
	"strconv"
	"strings"
	"time"
)

type (
	JSONFloat64 float64
	JSONInt64   int64
	JSONTime    time.Time
)

func (t *JSONFloat64) Float64() float64 { return float64(*t) }

func (t *JSONFloat64) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseFloat(r, 64)
	if err != nil {
		return err
	}
	*(*float64)(t) = q
	return
}

func (t *JSONInt64) Int64() int64 { return int64(*t) }

func (t *JSONInt64) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*int64)(t) = q
	return
}

func (t *JSONTime) Time() time.Time { return time.Time(*t) }

func (t *JSONTime) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	if r == "" {
		return
	}

	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.UnixMilli(q)
	return
}

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(*t).UnixMilli(), 10)), nil
}

func (t *JSONTime) String() string { return (time.Time)(*t).String() }
