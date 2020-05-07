package main

import (
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestContainer_increaseNumberOfBall(t *testing.T) {
	type fields struct {
		NumberOfBall int
		Verified     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "increase number of ball",
			fields: fields{
				NumberOfBall: 0,
				Verified:     false,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Container{
				NumberOfBall: tt.fields.NumberOfBall,
				Verified:     tt.fields.Verified,
			}

			c.increaseNumberOfBall()

			assert.Equal(t, tt.want, c.NumberOfBall)
		})
	}
}

func TestContainer_setVerified(t *testing.T) {
	type fields struct {
		NumberOfBall int
		Verified     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "set verified to true",
			fields: fields{
				NumberOfBall: 0,
				Verified:     false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Container{
				NumberOfBall: tt.fields.NumberOfBall,
				Verified:     tt.fields.Verified,
			}

			c.setVerified()
			assert.Equal(t, tt.want, c.Verified)
		})
	}
}

func TestNewContainer(t *testing.T) {
	tests := []struct {
		name string
		want Container
	}{
		{
			name: "new container",
			want: Container{
				NumberOfBall: 0,
				Verified:     false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContainer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}
