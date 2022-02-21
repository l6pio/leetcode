package M_134_Gas_Station

func canCompleteCircuit(gas []int, cost []int) int {
	gasLen := len(gas)
	for i := 0; i < gasLen; {
		j := 0
		totalGas := 0
		totalCost := 0
		for ; j < gasLen; j++ {
			totalGas += gas[(i+j)%gasLen]
			totalCost += cost[(i+j)%gasLen]
			if totalGas < totalCost {
				break
			}
		}
		if j == gasLen {
			return i
		} else {
			i += j + 1
		}
	}
	return -1
}
