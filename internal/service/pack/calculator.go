package pack

import "math"

// OptimalPacking represents the optimal pack combination for a given order
type OptimalPacking struct {
	Packs map[int]int
}

// PackCombination represents a pack combination with metadata
type PackCombination struct {
	Packs     map[int]int
	Total     int
	PackCount int
}

// calculatePacks finds the optimal pack combination for a given order quantity
func calculatePacks(orderItemQty int, packSizes []int) OptimalPacking {
	setPacks := make(map[int]struct{})
	packs := make(map[int]int)

	// Check if exact pack size exists for the order
	for _, v := range packSizes {
		if orderItemQty == v {
			packs[v] = 1
			return OptimalPacking{Packs: packs}
		}

		setPacks[v] = struct{}{}
	}

	maxValue := packSizes[0]                // Largest pack size
	minValue := packSizes[len(packSizes)-1] // Smallest pack size

	// Case 1: Order is smaller than smallest pack - use smallest pack
	if orderItemQty < minValue {
		packs[minValue] = 1
		return OptimalPacking{Packs: packs}
	}

	// Case 2: Order is smaller than largest pack - find optimal combination
	if orderItemQty < maxValue {
		best := findBestPackCombination(orderItemQty, packSizes)
		return OptimalPacking{Packs: best.Packs}
	}

	// Case 3: Order is larger than largest pack
	// Use as many largest packs as possible, then handle remainder
	count := orderItemQty / maxValue
	reminder := orderItemQty % maxValue
	packs[maxValue] = count

	// no remainder
	if reminder == 0 {
		return OptimalPacking{Packs: packs}
	}

	// Remainder matches an existing pack size
	if _, ok := setPacks[reminder]; ok {
		packs[reminder] = 1
		return OptimalPacking{Packs: packs}
	}

	// Remainder is smaller than smallest pack - use smallest pack
	if reminder < minValue {
		packs[minValue] = 1
		return OptimalPacking{Packs: packs}
	}

	// Find optimal combination for remainder
	best := findBestPackCombination(reminder, packSizes)
	for k, v := range best.Packs {
		packs[k] += v
	}

	return OptimalPacking{Packs: packs}
}

// findBestPackCombination uses DFS algorithm to find the optimal pack combination
func findBestPackCombination(orderQty int, packSizes []int) PackCombination {
	// packSizes should be sorted in descending order for efficiency
	best := PackCombination{Total: math.MaxInt64}

	var dfs func(index int, current map[int]int, total int, count int)

	dfs = func(index int, current map[int]int, total int, count int) {
		// Base case: we have enough items
		if total >= orderQty {
			// Update best if this combination is better (fewer total items or same total but fewer packs)
			if total < best.Total || (total == best.Total && count < best.PackCount) {
				newMap := make(map[int]int)
				for k, v := range current {
					newMap[k] = v
				}

				best = PackCombination{Packs: newMap, Total: total, PackCount: count}
			}
			return
		}

		// Base case: no more pack sizes to try
		if index >= len(packSizes) {
			return
		}

		packSize := packSizes[index]
		// Calculate maximum packs needed for this size (with buffer for optimization)
		maxPackCount := (orderQty-total)/packSize + 2

		for i := 0; i <= maxPackCount; i++ {
			if i > 0 {
				current[packSize] = i
			}
			dfs(index+1, current, total+packSize*i, count+i)
			if i > 0 {
				delete(current, packSize)
			}
		}
	}

	dfs(0, map[int]int{}, 0, 0)
	return best
}
