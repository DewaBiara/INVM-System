package dto

import (
	"reflect"
	"testing"

	"github.com/DewaBiara/INVM-System/pkg/entity"
)

func TestUserSignUpRequest_ToEntity(t *testing.T) {
	tests := []struct {
		name string
		u    *UserSignUpRequest
		want *entity.User
	}{
		{
			name: "All field filled",
			u: &UserSignUpRequest{
				Username: "username",
				Password: "password",
				Name:     "name",
				Telp:     "telp",
			},
			want: &entity.User{
				Username: "username",
				Password: "password",
				Name:     "name",
				Telp:     "telp",
			},
		},
		{
			name: "All field empty",
			u:    &UserSignUpRequest{},
			want: &entity.User{},
		},
		{
			name: "Partial field filled",
			u: &UserSignUpRequest{
				Username: "username",
				Password: "password",
			},
			want: &entity.User{
				Username: "username",
				Password: "password",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.ToEntity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserSignUpRequest.ToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
