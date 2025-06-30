package color

import (
	"testing"
)

func TestSprint(t *testing.T) {
	type args struct {
		str   string
		attrs []Attribute
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				str:   "hello",
				attrs: nil,
			},
			want: "\x1b[m" + "hello" + "\x1b[0m",
		},
		{
			name: "red",
			args: args{
				str:   "red",
				attrs: []Attribute{FgRed},
			},
			want: "\x1b[31m" + "red" + "\x1b[0m",
		},
		{
			name: "red+bold",
			args: args{
				str:   "redBold",
				attrs: []Attribute{FgRed, Bold},
			},
			want: "\x1b[31;1m" + "redBold" + "\x1b[0m",
		},
		{
			name: "red+bold+underline",
			args: args{
				str:   "redBoldUnderline",
				attrs: []Attribute{FgRed, Bold, Underline},
			},
			want: "\x1b[31;1;4m" + "redBoldUnderline" + "\x1b[0m",
		},
		{
			name: "red+bold+underline+italic",
			args: args{
				str:   "redBoldUnderlineItalic",
				attrs: []Attribute{FgRed, Bold, Underline, Italic},
			},
			want: "\x1b[31;1;4;3m" + "redBoldUnderlineItalic" + "\x1b[0m",
		},
		{
			name: "red+bold+underline+italic+inverse",
			args: args{
				str:   "redBoldUnderlineItalicInverse",
				attrs: []Attribute{FgRed, Bold, Underline, Italic, ReverseVideo},
			},
			want: "\x1b[31;1;4;3;7m" + "redBoldUnderlineItalicInverse" + "\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(tt.args.attrs...)
			t.Log(c.Sprint(tt.args.str))
			if got := c.Sprint(tt.args.str); got != tt.want {
				t.Errorf("Color.Sprintf() = %v, want = %v", got, tt.want)
			}
		})
	}
}
