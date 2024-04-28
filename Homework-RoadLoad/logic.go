package main

func Road(maximumWeight uint, vehicleWeights []uint) uint {
	if len(vehicleWeights) == 0 {
		return 0
	}

	if len(vehicleWeights) == 1 {
		if vehicleWeights[0] > maximumWeight {
			return 1
		}

		return 0
	}

	var returnedVehicles uint

	i := 0
	j := 1

	for {
		if j > len(vehicleWeights)-1 || i > len(vehicleWeights)-1 {
			break
		}

		if vehicleWeights[i] > maximumWeight {
			i++
			returnedVehicles++

			if vehicleWeights[j] > maximumWeight {
				i++
				returnedVehicles++
			}

			j++

			continue
		}

		if vehicleWeights[j] > maximumWeight {
			j++
			returnedVehicles++

			continue
		}

		currentWeight := vehicleWeights[i] + vehicleWeights[j]
		if currentWeight <= maximumWeight {
			i++
			j++

			continue
		}

		returnedVehicles++

		if vehicleWeights[i] > vehicleWeights[j] {
			i++
		}

		j++
	}

	return returnedVehicles
}
