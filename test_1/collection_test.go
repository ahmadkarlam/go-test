package main

import (
	"reflect"
	"testing"
)

func TestCollection_isThereAFullContainer(t *testing.T) {
	type fields struct {
		containers map[int]Container
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "success found full container",
			fields: fields{
				containers: map[int]Container{
					1: {
						NumberOfBall: 1,
						Verified:     false,
					},
					2: {
						NumberOfBall: 2,
						Verified:     true,
					},
				},
			},
			want: true,
		},
		{
			name: "full container not found",
			fields: fields{
				containers: map[int]Container{
					1: {
						NumberOfBall: 1,
						Verified:     false,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Collection{
				containers: tt.fields.containers,
			}
			if got := c.isThereAFullContainer(); got != tt.want {
				t.Errorf("isThereAFullContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCollection(t *testing.T) {
	tests := []struct {
		name string
		want Collection
	}{
		{
			name: "new collection",
			want: Collection{
				containers: make(map[int]Container),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCollection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_fill(t *testing.T) {
	type fields struct {
		containers map[int]Container
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]Container
	}{
		{
			name: "success fill container",
			fields: fields{
				containers: make(map[int]Container),
			},
			args: args{
				key: 1,
			},
			want: map[int]Container{
				1: {
					NumberOfBall: 1,
					Verified:     false,
				},
			},
		},
		{
			name: "fulling container",
			fields: fields{
				containers: map[int]Container{
					1: {
						NumberOfBall: 1,
						Verified:     false,
					},
				},
			},
			args: args{
				key: 1,
			},
			want: map[int]Container{
				1: {
					NumberOfBall: 2,
					Verified:     true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Collection{
				containers: tt.fields.containers,
			}

			c.fill(tt.args.key)
			got := c.containers
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
