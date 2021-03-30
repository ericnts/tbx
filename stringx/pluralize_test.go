package stringx

import "testing"

func TestPlural(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				str: "table",
			},
			want: "tables",
		},
		{
			name: "case2",
			args: args{
				str: "footTable",
			},
			want: "footTables",
		},
		{
			name: "case3",
			args: args{
				str: "foot",
			},
			want: "feet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Plural(tt.args.str); got != tt.want {
				t.Errorf("Plural() = %v, want %v", got, tt.want)
			}
		})
	}
}
