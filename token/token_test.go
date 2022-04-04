package token

import (
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {
	type args struct {
		key        string
		expiration int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "test token",
			args: args{
				key:        "15014077321",
				expiration: 60,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetToken(tt.args.key, tt.args.expiration)
			fmt.Println(got, got1)
			token, err := GetTokenByRefreshToken(got1)
			err2 := CheckToken(token)
			fmt.Println(err, err2)
		})
	}
}
