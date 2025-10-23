package tea

import "example.com/errors-demo/app"

/*
	func MakeTea(hasTea, hasWater, kettleWorking bool) (string, error) {
		if !hasTea {
			return "", ErrNoTea
		}
		if !hasWater {
			return "", fmt.Errorf("preparation failed: %w", ErrNoWaterTea)
		}
		if !kettleWorking {
			return "", fmt.Errorf("failed to boil water: %w", ErrKettleBroken)
		}
		return "☕ Tea is ready!", nil
	}
*/

func MakeTea(hasTea, hasWater, kettleWorking bool) (string, error) {
	if !hasTea {
		return "", app.WrapError("NO_TEA", "Cannot prepare tea", ErrNoTea)
	}
	if !hasWater {
		return "", app.WrapError("NO_WATER", "Preparation failed", ErrNoWaterTea)
	}
	if !kettleWorking {
		return "", app.WrapError("BROKEN_KETTLE", "Boiling failed", ErrKettleBroken)
	}
	return "☕ Tea ready!", nil
}
