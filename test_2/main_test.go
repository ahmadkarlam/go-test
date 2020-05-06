package main

import "testing"

func Test_orderProduct(t *testing.T) {
	type args struct {
		collection ProductCollection
		name       string
	}
	collection := NewProductCollection()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success order product",
			args: args{
				collection: collection,
				name:       "Baju",
			},
			wantErr: false,
		},
		{
			name: "product out of stock",
			args: args{
				collection: collection,
				name:       "Baju",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := orderProduct(tt.args.collection, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("orderProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
