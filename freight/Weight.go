package freight

import "math"

// CalculateWeightValue calcula o custo do peso.
func CalculateWeightValue(
	realWeight,
	volume,
	cubicFactor,
	pricePerKg float64,
) float64 {
	volumetricWeight := CalculateVolumetricWeight(volume, cubicFactor)
	chargeableWeight := CalculateChargeableWeight(realWeight, volumetricWeight)

	return chargeableWeight * pricePerKg
}

// CalculateVolumetricWeight calcula o peso cúbico.
func CalculateVolumetricWeight(volume, cubicFactor float64) float64 {
	return volume * cubicFactor
}

// O maior valor entre peso real e peso cúbico é utilizado.
func CalculateChargeableWeight(realWeight, volumetricWeight float64) float64 {
	return math.Max(realWeight, volumetricWeight)
}
