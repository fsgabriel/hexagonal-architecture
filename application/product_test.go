package application

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	tests := []struct {
		name    string
		p       *Product
		wantErr bool
	}{
		{
			name: "Success",
			p: &Product{
				ID:    "1",
				Name:  "Product 1",
				Price: 10,
			},
			wantErr: false,
		},
		{
			name: "Error",
			p: &Product{
				ID:    "1",
				Name:  "Product 2",
				Price: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Enable(); tt.wantErr {
				require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestProduct_Disable(t *testing.T) {
	tests := []struct {
		name    string
		p       *Product
		wantErr bool
	}{
		{
			name: "Success",
			p: &Product{
				ID:    "1",
				Name:  "Product 1",
				Price: 0,
			},
			wantErr: false,
		},
		{
			name: "Error",
			p: &Product{
				ID:    "2",
				Name:  "Product 2",
				Price: 10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Disable(); tt.wantErr {
				require.Equal(t, "the price must be zero in order to have the product disable", err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestProduct_IsValid(t *testing.T) {
	tests := []struct {
		name    string
		p       *Product
		want    bool
		wantErr bool
	}{
		{
			name: "Valid",
			p: &Product{
				ID:     uuid.NewV4().String(),
				Name:   "P",
				Price:  10,
				Status: ENABLE,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Invalid status",
			p: &Product{
				ID:     uuid.NewV4().String(),
				Name:   "P",
				Price:  10,
				Status: "Invalid",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Price lower than 0",
			p: &Product{
				ID:     uuid.NewV4().String(),
				Name:   "P",
				Price:  -10,
				Status: ENABLE,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Missing field",
			p: &Product{
				Name:   "P",
				Price:  10,
				Status: ENABLE,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.IsValid()
			if (err != nil) != tt.wantErr {
				t.Errorf("Product.IsValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Product.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct_GetID(t *testing.T) {
	tests := []struct {
		name string
		p    *Product
		want string
	}{
		{
			name: "Ok",
			p: &Product{
				ID: "ID",
			},
			want: "ID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetID(); got != tt.want {
				t.Errorf("Product.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct_GetName(t *testing.T) {
	tests := []struct {
		name string
		p    *Product
		want string
	}{
		{
			name: "Ok",
			p: &Product{
				Name: "Name",
			},
			want: "Name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetName(); got != tt.want {
				t.Errorf("Product.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct_GetStatus(t *testing.T) {
	tests := []struct {
		name string
		p    *Product
		want string
	}{
		{
			name: "Ok",
			p: &Product{
				Status: "Status",
			},
			want: "Status",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetStatus(); got != tt.want {
				t.Errorf("Product.GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct_GetPrice(t *testing.T) {
	tests := []struct {
		name string
		p    *Product
		want float64
	}{
		{
			name: "Ok",
			p: &Product{
				Price: 312,
			},
			want: 312,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetPrice(); got != tt.want {
				t.Errorf("Product.GetPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
