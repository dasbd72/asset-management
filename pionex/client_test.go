package pionex

import "testing"

func Test_sign(t *testing.T) {
	type args struct {
		secret  string
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			want: "b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad",
		},
		{
			name: "non-empty",
			args: args{
				"sec",
				`{"msg":"test"}`,
			},
			want: "23aa07ef2fc07cd18c30a180feea514b0f2717499cc070d7715c9c1a0e4e7c7a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sign(tt.args.secret, tt.args.message); got != tt.want {
				t.Errorf("sign() = %v, want %v", got, tt.want)
			}
		})
	}
}
