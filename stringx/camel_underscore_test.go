package stringx

import "testing"

func TestLowerCamelName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case1",
			args: args{
				name: "office_id",
			},
			want: "officeID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerCamelName(tt.args.name); got != tt.want {
				t.Errorf("LowerCamelName() = %v, want %v", got, tt.want)
			}
		})
	}
}
