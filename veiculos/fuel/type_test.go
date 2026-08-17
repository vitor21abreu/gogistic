package fuel

import "testing"

func TestFuelTypes(t *testing.T) {
	tests := []struct {
		name string
		got  FuelType
		want string
	}{
		{
			name: "gasoline",
			got:  Gasoline,
			want: "GASOLINA",
		},
		{
			name: "ethanol",
			got:  Ethanol,
			want: "ETANOL",
		},
		{
			name: "diesel",
			got:  Diesel,
			want: "OLEO DIESEL",
		},
		{
			name: "gnv",
			got:  GNV,
			want: "GNV",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.got) != tt.want {
				t.Errorf(
					"expected %q, got %q",
					tt.want,
					tt.got,
				)
			}
		})
	}
}
