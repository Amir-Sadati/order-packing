package pack

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	packSizes := []int{5000, 2000, 1000, 500, 250}

	tests := []struct {
		name          string
		orderItemQty  int
		expectedPacks map[int]int
		expectedTotal int
		description   string
	}{
		{
			name:         "Exact pack size - 5000",
			orderItemQty: 5000,
			expectedPacks: map[int]int{
				5000: 1,
			},
			expectedTotal: 5000,
			description:   "Should return exactly one 5000 pack when order quantity matches pack size",
		},
		{
			name:         "Exact pack size - 2000",
			orderItemQty: 2000,
			expectedPacks: map[int]int{
				2000: 1,
			},
			expectedTotal: 2000,
			description:   "Should return exactly one 2000 pack when order quantity matches pack size",
		},
		{
			name:         "Exact pack size - 1000",
			orderItemQty: 1000,
			expectedPacks: map[int]int{
				1000: 1,
			},
			expectedTotal: 1000,
			description:   "Should return exactly one 1000 pack when order quantity matches pack size",
		},
		{
			name:         "Exact pack size - 500",
			orderItemQty: 500,
			expectedPacks: map[int]int{
				500: 1,
			},
			expectedTotal: 500,
			description:   "Should return exactly one 500 pack when order quantity matches pack size",
		},
		{
			name:         "Exact pack size - 250",
			orderItemQty: 250,
			expectedPacks: map[int]int{
				250: 1,
			},
			expectedTotal: 250,
			description:   "Should return exactly one 250 pack when order quantity matches pack size",
		},
		{
			name:         "Large number - 15000",
			orderItemQty: 15000,
			expectedPacks: map[int]int{
				5000: 3,
			},
			expectedTotal: 15000,
			description:   "Should return three 5000 packs for large number divisible by largest pack size",
		},
		{
			name:         "Large number with remainder - 12500",
			orderItemQty: 12500,
			expectedPacks: map[int]int{
				5000: 2,
				2000: 1,
				500:  1,
			},
			expectedTotal: 12500,
			description:   "Should return optimal combination for large number with remainder",
		},
		{
			name:         "Number between pack sizes - 750",
			orderItemQty: 750,
			expectedPacks: map[int]int{
				500: 1,
				250: 1,
			},
			expectedTotal: 750,
			description:   "Should return optimal combination for number between pack sizes",
		},
		{
			name:         "Number smaller than smallest pack - 100",
			orderItemQty: 100,
			expectedPacks: map[int]int{
				250: 1,
			},
			expectedTotal: 250,
			description:   "Should return smallest pack when order quantity is less than smallest pack size",
		},
		{
			name:         "Complex combination - 3750",
			orderItemQty: 3750,
			expectedPacks: map[int]int{
				2000: 1,
				1000: 1,
				500:  1,
				250:  1,
			},
			expectedTotal: 3750,
			description:   "Should return optimal combination using multiple pack sizes",
		},
		{
			name:         "Very large number - 124001",
			orderItemQty: 124001,
			expectedPacks: map[int]int{
				5000: 24,
				2000: 2,
				250:  1,
			},
			expectedTotal: 124250,
			description:   "Should handle very large numbers efficiently with optimal combination",
		},
		{
			name:         "Extremely large number - 1324001",
			orderItemQty: 1324001,
			expectedPacks: map[int]int{
				5000: 264,
				2000: 2,
				250:  1,
			},
			expectedTotal: 1324250,
			description:   "Should handle extremely large numbers efficiently with optimal combination",
		},
		{
			name:         "Large number with complex remainder - 123411",
			orderItemQty: 123411,
			expectedPacks: map[int]int{
				5000: 24,
				2000: 1,
				1000: 1,
				500:  1,
			},
			expectedTotal: 123500,
			description:   "Should handle large numbers with complex remainder efficiently",
		},
		{
			name:         "Large number divisible by 5000 - 100000",
			orderItemQty: 100000,
			expectedPacks: map[int]int{
				5000: 20,
			},
			expectedTotal: 100000,
			description:   "Should handle large numbers perfectly divisible by largest pack size",
		},
		{
			name:         "Large number with small remainder - 100001",
			orderItemQty: 100001,
			expectedPacks: map[int]int{
				5000: 20,
				250:  1,
			},
			expectedTotal: 100250,
			description:   "Should handle large numbers with small remainder efficiently",
		},
		{
			name:         "Large number requiring multiple pack types - 98765",
			orderItemQty: 98765,
			expectedPacks: map[int]int{
				5000: 19,
				2000: 2,
			},
			expectedTotal: 99000,
			description:   "Should handle large numbers requiring multiple pack types efficiently",
		},
		{
			name:         "Large number with exact pack match - 50000",
			orderItemQty: 50000,
			expectedPacks: map[int]int{
				5000: 10,
			},
			expectedTotal: 50000,
			description:   "Should handle large numbers that exactly match pack size multiples",
		},
		{
			name:         "Large number with remainder close to pack size - 124999",
			orderItemQty: 124999,
			expectedPacks: map[int]int{
				5000: 25,
			},
			expectedTotal: 125000,
			description:   "Should handle large numbers with remainder close to pack size",
		},
		{
			name:         "Large number requiring all pack types - 8765",
			orderItemQty: 8765,
			expectedPacks: map[int]int{
				5000: 1,
				2000: 2,
			},
			expectedTotal: 9000,
			description:   "Should handle large numbers requiring all available pack types",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculatePacks(tt.orderItemQty, packSizes)

			// Check if packs match expected
			if !reflect.DeepEqual(result.Packs, tt.expectedPacks) {
				t.Errorf("calculatePacks() packs = %v, want %v", result.Packs, tt.expectedPacks)
			}

			// Calculate total from result
			total := 0
			for packSize, count := range result.Packs {
				total += packSize * count
			}

			// Check if total matches expected
			if total != tt.expectedTotal {
				t.Errorf("calculatePacks() total = %v, want %v", total, tt.expectedTotal)
			}

			// Verify that total is >= order quantity (should never be less)
			if total < tt.orderItemQty {
				t.Errorf("calculatePacks() total %v is less than order quantity %v", total, tt.orderItemQty)
			}

			t.Logf("Test passed: %s", tt.description)
		})
	}
}

func TestCalculatePacksWithLargeNumbers(t *testing.T) {
	packSizes := []int{5000, 2000, 1000, 500, 250}

	tests := []struct {
		name          string
		orderItemQty  int
		expectedPacks map[int]int
		expectedTotal int
		description   string
	}{
		{
			name:         "Million plus number - 124001",
			orderItemQty: 124001,
			expectedPacks: map[int]int{
				5000: 24,
				2000: 2,
				250:  1,
			},
			expectedTotal: 124250,
			description:   "Should handle numbers over 100k efficiently",
		},
		{
			name:         "Million plus number - 1324001",
			orderItemQty: 1324001,
			expectedPacks: map[int]int{
				5000: 264,
				2000: 2,
				250:  1,
			},
			expectedTotal: 1324250,
			description:   "Should handle numbers over 1M efficiently",
		},
		{
			name:         "Complex large number - 123411",
			orderItemQty: 123411,
			expectedPacks: map[int]int{
				5000: 24,
				2000: 1,
				1000: 1,
				500:  1,
			},
			expectedTotal: 123500,
			description:   "Should handle complex large numbers efficiently",
		},
		{
			name:         "Very large number - 500000",
			orderItemQty: 500000,
			expectedPacks: map[int]int{
				5000: 100,
			},
			expectedTotal: 500000,
			description:   "Should handle very large numbers perfectly divisible by 5000",
		},
		{
			name:         "Very large number with remainder - 500001",
			orderItemQty: 500001,
			expectedPacks: map[int]int{
				5000: 100,
				250:  1,
			},
			expectedTotal: 500250,
			description:   "Should handle very large numbers with small remainder",
		},
		{
			name:         "Extremely large number - 1000000",
			orderItemQty: 1000000,
			expectedPacks: map[int]int{
				5000: 200,
			},
			expectedTotal: 1000000,
			description:   "Should handle extremely large numbers efficiently",
		},
		{
			name:         "Large number with complex remainder - 999999",
			orderItemQty: 999999,
			expectedPacks: map[int]int{
				5000: 200,
			},
			expectedTotal: 1000000,
			description:   "Should handle large numbers with complex remainder",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculatePacks(tt.orderItemQty, packSizes)

			// Check if packs match expected
			if !reflect.DeepEqual(result.Packs, tt.expectedPacks) {
				t.Errorf("calculatePacks() packs = %v, want %v", result.Packs, tt.expectedPacks)
			}

			// Calculate total from result
			total := 0
			for packSize, count := range result.Packs {
				total += packSize * count
			}

			// Check if total matches expected
			if total != tt.expectedTotal {
				t.Errorf("calculatePacks() total = %v, want %v", total, tt.expectedTotal)
			}

			// Verify that total is >= order quantity (should never be less)
			if total < tt.orderItemQty {
				t.Errorf("calculatePacks() total %v is less than order quantity %v", total, tt.orderItemQty)
			}

			// Verify that the algorithm doesn't use more packs than necessary
			// For large numbers, we should use mostly 5000 packs
			if tt.orderItemQty >= 100000 {
				expected5000Packs := tt.orderItemQty / 5000

				actual5000Packs := result.Packs[5000]
				if actual5000Packs < expected5000Packs {
					t.Errorf("Expected at least %d packs of 5000, got %d", expected5000Packs, actual5000Packs)
				}
			}

			t.Logf("Test passed: %s - Total: %d, Packs: %v", tt.description, total, result.Packs)
		})
	}
}
