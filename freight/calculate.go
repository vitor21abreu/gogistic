package freight

// CalculateTotalFreight soma todos os componentes do frete.
func CalculateTotalFreight(
	distanceValue,
	weightValue,
	volumeValue,
	baseRate float64,
) float64 {
	return distanceValue +
		weightValue +
		volumeValue +
		baseRate
}

// CalculateFreight calcula o valor total do frete.
func (f Freight) CalculateFreight() float64 {
	volume := CalculateVolume(
		f.Length,
		f.Width,
		f.Height,
	)

	distanceValue := CalculateDistanceValue(
		f.Distance,
		f.KmValue,
	)

	weightValue := CalculateWeightValue(
		f.RealWeight,
		volume,
		f.CubicFactor,
		f.PricePerKg,
	)

	volumeValue := CalculateVolumeValue(
		volume,
		f.PricePerCubicMeter,
	)

	return CalculateTotalFreight(
		distanceValue,
		weightValue,
		volumeValue,
		f.BaseRate,
	)
}
