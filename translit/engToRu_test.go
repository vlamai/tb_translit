package translit

import "testing"

func TestEngToRu(t *testing.T) {
	type args struct {
		inp string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				inp: "q",
			},
			want: "й",
		},
		{
			name: "",
			args: args{
				inp: "Q",
			},
			want: "Й",
		},
		{
			name: "",
			args: args{
				inp: "Rfrjq-nj ntrcn/ D rjnjhjv! tcnm: yt njkmrj ,erds/?",
			},
			want: "Какой-то текст. В котором! есть: не только буквы.?",
		},
		{
			name: "",
			args: args{
				inp: "14 'nj xbckj/ Gjl]tpl 'nj ckjdj/",
			},
			want: "14 это число. Подъезд это слово.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EngToRu(tt.args.inp); got != tt.want || len(got) != len(tt.want) {
				t.Errorf("EngToRu() = \n%v \nwant \n%v", got, tt.want)
			}
		})
	}
}
