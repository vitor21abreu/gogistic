package freight

// Freight contém os dados necessários para calcular o frete.
type Freight struct {
	// Distância
	Distance float64
	KmValue  float64

	// Peso
	RealWeight  float64
	CubicFactor float64
	PricePerKg  float64

	// Dimensões
	Length             float64
	Width              float64
	Height             float64
	PricePerCubicMeter float64

	// Taxa fixa
	BaseRate float64
}
