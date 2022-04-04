package controller

import (
	"net/http"
	"testing"
)

func TestUserController_FindPassword(t *testing.T) {
	type args struct {
		response http.ResponseWriter
		request  *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "13407315009",
			args: args{response: nil, request: nil},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//uc := &UserController{}
		})
	}
}
