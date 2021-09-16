package money

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSumOverflow(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1+1",
			args: args{
				a: 1,
				b: 1,
			},
			want: false,
		},
		{
			name: "maxint in result",
			args: args{
				a: math.MaxInt64 - 10,
				b: 10,
			},
			want: false,
		},
		{
			name: "more then maxint",
			args: args{
				a: math.MaxInt64 - 1,
				b: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSumOverflow(tt.args.a, tt.args.b)

			require.Equal(t, tt.want, got)
		})
	}
}
