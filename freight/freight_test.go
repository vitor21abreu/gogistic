package freight

import (
	"math"
	"testing"
)

func TestCalculateTotalFreight(t *testing.T) {
	tests := []struct {
		name          string
		distanceValue float64
		weightValue   float64
		volumeValue   float64
		baseRate      float64
		expected      float64
	}{
		{
			name:          "soma todos os componentes",
			distanceValue: 100,
			weightValue:   50,
			volumeValue:   25,
			baseRate:      10,
			expected:      185,
		},
		{
			name:          "valores zerados",
			distanceValue: 0,
			weightValue:   0,
			volumeValue:   0,
			baseRate:      0,
			expected:      0,
		},
		{
			name:          "valores decimais",
			distanceValue: 10.5,
			weightValue:   20.25,
			volumeValue:   5.75,
			baseRate:      3.5,
			expected:      40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateTotalFreight(
				tt.distanceValue,
				tt.weightValue,
				tt.volumeValue,
				tt.baseRate,
			)

			if got != tt.expected {
				t.Errorf("CalculateTotalFreight() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCalculateDistanceValue(t *testing.T) {
	tests := []struct {
		name       string
		distance   float64
		pricePerKm float64
		expected   float64
	}{
		{
			name:       "calcula custo da distância",
			distance:   100,
			pricePerKm: 2.5,
			expected:   250,
		},
		{
			name:       "distância zero",
			distance:   0,
			pricePerKm: 2.5,
			expected:   0,
		},
		{
			name:       "preço por km zero",
			distance:   100,
			pricePerKm: 0,
			expected:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateDistanceValue(tt.distance, tt.pricePerKm)

			if got != tt.expected {
				t.Errorf("CalculateDistanceValue() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCalculateVolume(t *testing.T) {
	tests := []struct {
		name     string
		length   float64
		width    float64
		height   float64
		expected float64
	}{
		{
			name:     "calcula volume",
			length:   2,
			width:    3,
			height:   4,
			expected: 24,
		},
		{
			name:     "uma dimensão zero",
			length:   2,
			width:    3,
			height:   0,
			expected: 0,
		},
		{
			name:     "valores decimais",
			length:   1.5,
			width:    2,
			height:   0.5,
			expected: 1.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateVolume(tt.length, tt.width, tt.height)

			if got != tt.expected {
				t.Errorf("CalculateVolume() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCalculateVolumeValue(t *testing.T) {
	tests := []struct {
		name               string
		volume             float64
		pricePerCubicMeter float64
		expected           float64
	}{
		{
			name:               "calcula custo do volume",
			volume:             10,
			pricePerCubicMeter: 20,
			expected:           200,
		},
		{
			name:               "volume zero",
			volume:             0,
			pricePerCubicMeter: 20,
			expected:           0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateVolumeValue(
				tt.volume,
				tt.pricePerCubicMeter,
			)

			if got != tt.expected {
				t.Errorf("CalculateVolumeValue() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCalculateVolumetricWeight(t *testing.T) {
	tests := []struct {
		name        string
		volume      float64
		cubicFactor float64
		expected    float64
	}{
		{
			name:        "calcula peso cúbico",
			volume:      2,
			cubicFactor: 300,
			expected:    600,
		},
		{
			name:        "volume zero",
			volume:      0,
			cubicFactor: 300,
			expected:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateVolumetricWeight(
				tt.volume,
				tt.cubicFactor,
			)

			if got != tt.expected {
				t.Errorf("CalculateVolumetricWeight() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCalculateChargeableWeight(t *testing.T) {
	tests := []struct {
		name             string
		realWeight       float64
		volumetricWeight float64
		expected         float64
	}{
		{
			name:             "peso real maior",
			realWeight:       100,
			volumetricWeight: 80,
			expected:         100,
		},
		{
			name:             "peso cúbico maior",
			realWeight:       80,
			volumetricWeight: 100,
			expected:         100,
		},
		{
			name:             "pesos iguais",
			realWeight:       100,
			volumetricWeight: 100,
			expected:         100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateChargeableWeight(
				tt.realWeight,
				tt.volumetricWeight,
			)

			if got != tt.expected {
				t.Errorf("CalculateChargeableWeight() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCalculateWeightValue(t *testing.T) {
	tests := []struct {
		name        string
		realWeight  float64
		volume      float64
		cubicFactor float64
		pricePerKg  float64
		expected    float64
	}{
		{
			name:        "peso real maior que peso cúbico",
			realWeight:  100,
			volume:      0.2,
			cubicFactor: 300,
			pricePerKg:  5,
			expected:    500,
		},
		{
			name:        "peso cúbico maior que peso real",
			realWeight:  50,
			volume:      0.5,
			cubicFactor: 300,
			pricePerKg:  5,
			expected:    750,
		},
		{
			name:        "pesos iguais",
			realWeight:  100,
			volume:      1.0 / 3.0,
			cubicFactor: 300,
			pricePerKg:  5,
			expected:    500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateWeightValue(
				tt.realWeight,
				tt.volume,
				tt.cubicFactor,
				tt.pricePerKg,
			)

			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("CalculateWeightValue() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFreightCalculateFreight(t *testing.T) {
	tests := []struct {
		name     string
		freight  Freight
		expected float64
	}{
		{
			name: "calcula frete completo usando peso real",
			freight: Freight{
				Distance:           100,
				KmValue:            2,
				RealWeight:         50,
				CubicFactor:        100,
				PricePerKg:         5,
				Length:             1,
				Width:              1,
				Height:             0.5,
				PricePerCubicMeter: 100,
				BaseRate:           20,
			},
			// Distância: 100 * 2 = 200
			// Volume: 1 * 1 * 0.5 = 0.5
			// Peso cúbico: 0.5 * 100 = 50
			// Peso cobrado: max(50, 50) = 50
			// Peso: 50 * 5 = 250
			// Volume: 0.5 * 100 = 50
			// Total: 200 + 250 + 50 + 20 = 520
			expected: 520,
		},
		{
			name: "calcula frete completo usando peso cúbico",
			freight: Freight{
				Distance:           50,
				KmValue:            3,
				RealWeight:         20,
				CubicFactor:        200,
				PricePerKg:         4,
				Length:             1,
				Width:              1,
				Height:             1,
				PricePerCubicMeter: 50,
				BaseRate:           10,
			},
			// Distância: 50 * 3 = 150
			// Volume: 1
			// Peso cúbico: 1 * 200 = 200
			// Peso cobrado: 200
			// Peso: 200 * 4 = 800
			// Volume: 1 * 50 = 50
			// Total: 150 + 800 + 50 + 10 = 1010
			expected: 1010,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.freight.CalculateFreight()

			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("CalculateFreight() = %v, want %v", got, tt.expected)
			}
		})
	}
}
