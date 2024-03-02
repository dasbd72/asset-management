package cast

import (
	"reflect"
	"testing"
)

func TestIfToNilOrString(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want NilOrString
	}{
		{
			name: "String",
			args: args{value: "test"},
			want: NilOrString{value: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfToNilOrString(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IfToNilOrString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrString_String(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Text",
			fields: fields{value: "test"},
			want:   "test",
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   "",
		},
		{
			name:   "Number",
			fields: fields{value: 123},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := IfToNilOrString(tt.fields.value)
			if got := v.String(); got != tt.want {
				t.Errorf("NilOrString.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrString_Value(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name:   "Text",
			fields: fields{value: "test"},
			want:   "test",
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   nil,
		},
		{
			name:   "Number",
			fields: fields{value: 123},
			want:   123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrString{
				value: tt.fields.value,
			}
			if got := v.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NilOrString.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrString_IsNil(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Text",
			fields: fields{value: "test"},
			want:   false,
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrString{value: tt.fields.value}
			if got := v.IsNil(); got != tt.want {
				t.Errorf("NilOrString.IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrString_Valid(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Text",
			fields: fields{value: "test"},
			want:   true,
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   true,
		},
		{
			name:   "Number",
			fields: fields{value: 123},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrString{
				value: tt.fields.value,
			}
			if got := v.Valid(); got != tt.want {
				t.Errorf("NilOrString.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrString_MarshalJSON(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "Text",
			fields:  fields{value: "test"},
			want:    []byte(`"test"`),
			wantErr: false,
		},
		{
			name:    "Nil",
			fields:  fields{value: nil},
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name:    "Number",
			fields:  fields{value: 123},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrString{
				value: tt.fields.value,
			}
			got, err := v.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("NilOrString.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NilOrString.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfToNilOrFloat64(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want NilOrFloat64
	}{
		{
			name: "Float",
			args: args{value: 123.45},
			want: NilOrFloat64{value: 123.45},
		},
		{
			name: "Nil",
			args: args{value: nil},
			want: NilOrFloat64{value: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfToNilOrFloat64(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IfToNilOrFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrFloat64_Float64(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Float",
			fields: fields{value: 123.45},
			want:   123.45,
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   0,
		},
		{
			name:   "Text",
			fields: fields{value: "test"},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrFloat64{
				value: tt.fields.value,
			}
			if got := v.Float64(); got != tt.want {
				t.Errorf("NilOrFloat64.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrFloat64_Value(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name:   "Float",
			fields: fields{value: 123.45},
			want:   123.45,
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrFloat64{
				value: tt.fields.value,
			}
			if got := v.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NilOrFloat64.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrFloat64_IsNil(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Float",
			fields: fields{value: 123.45},
			want:   false,
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   true,
		},
		{
			name:   "Text",
			fields: fields{value: "test"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrFloat64{
				value: tt.fields.value,
			}
			if got := v.IsNil(); got != tt.want {
				t.Errorf("NilOrFloat64.IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrFloat64_Valid(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Float",
			fields: fields{value: 123.45},
			want:   true,
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   true,
		},
		{
			name:   "Text",
			fields: fields{value: "test"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrFloat64{
				value: tt.fields.value,
			}
			if got := v.Valid(); got != tt.want {
				t.Errorf("NilOrFloat64.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrFloat64_MarshalJSON(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "Float",
			fields:  fields{value: 123.45},
			want:    []byte("123.45"),
			wantErr: false,
		},
		{
			name:    "Nil",
			fields:  fields{value: nil},
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name:    "Text",
			fields:  fields{value: "test"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrFloat64{
				value: tt.fields.value,
			}
			got, err := v.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("NilOrFloat64.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), string(tt.want)) {
				t.Errorf("NilOrFloat64.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestIfToNilOrInt(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want NilOrInt
	}{
		{
			name: "Number",
			args: args{value: 123},
			want: NilOrInt{value: 123},
		},
		{
			name: "Nil",
			args: args{value: nil},
			want: NilOrInt{value: nil},
		},
		{
			name: "Text",
			args: args{value: "test"},
			want: NilOrInt{value: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfToNilOrInt(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IfToNilOrInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrInt_Int(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Number",
			fields: fields{value: 123},
			want:   123,
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   0,
		},
		{
			name:   "Text",
			fields: fields{value: "test"},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrInt{
				value: tt.fields.value,
			}
			if got := v.Int(); got != tt.want {
				t.Errorf("%v NilOrInt.Int() = %v, want %v", v, got, tt.want)
			}
		})
	}
}

func TestNilOrInt_Value(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name:   "Number",
			fields: fields{value: 123},
			want:   int(123),
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrInt{
				value: tt.fields.value,
			}
			if got := v.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NilOrInt.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrInt_IsNil(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Number",
			fields: fields{value: 123},
			want:   false,
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrInt{
				value: tt.fields.value,
			}
			if got := v.IsNil(); got != tt.want {
				t.Errorf("NilOrInt.IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrInt_Valid(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Number",
			fields: fields{value: 123},
			want:   true,
		},
		{
			name:   "Nil",
			fields: fields{value: nil},
			want:   true,
		},
		{
			name:   "Text",
			fields: fields{value: "test"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrInt{
				value: tt.fields.value,
			}
			if got := v.Valid(); got != tt.want {
				t.Errorf("NilOrInt.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNilOrInt_MarshalJSON(t *testing.T) {
	type fields struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "Number",
			fields:  fields{value: 123},
			want:    []byte("123"),
			wantErr: false,
		},
		{
			name:    "Nil",
			fields:  fields{value: nil},
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name:    "Text",
			fields:  fields{value: "test"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NilOrInt{
				value: tt.fields.value,
			}
			got, err := v.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("NilOrInt.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NilOrInt.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
