package hackerrank

import "log"

func HRTructTour() {
	log.Println(tructTour([][]int32{{1, 5}, {10, 3}, {3, 4}}))
}

func tructTour(petrolpumps [][]int32) int32 {
	for i := 0; i < len(petrolpumps); i++ {
		petrol := 0
		isSkip := false
		if petrolpumps[i][0] < petrolpumps[i][1] {
			continue
		}
		petrol += int(petrolpumps[i][0]) - int(petrolpumps[i][1])

		for k := i + 1; k < len(petrolpumps); k++ {
			if petrolpumps[k][0]+int32(petrol) < petrolpumps[k][1] {
				isSkip = true
				break
			} else {
				petrol += int(petrolpumps[k][0]) - int(petrolpumps[k][1])
			}
		}

		if isSkip {
			continue
		}

		for j := 0; j < i; j++ {
			if petrolpumps[j][0]+int32(petrol) < petrolpumps[j][1] {
				isSkip = true
				break
			} else {
				petrol += int(petrolpumps[j][0]) - int(petrolpumps[j][1])
			}
		}

		if isSkip {
			continue
		}

		return int32(i)
	}

	return 0
}
