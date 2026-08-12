package freight

// CalculateVolume calcula o volume em m³.
func CalculateVolume(length, width, height float64) float64 {
	return length * width * height
}

// CalculateVolumeValue calcula o custo referente ao volume.
func CalculateVolumeValue(volume, pricePerCubicMeter float64) float64 {
	return volume * pricePerCubicMeter
}
