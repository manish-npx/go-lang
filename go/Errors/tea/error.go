package tea

import "errors"

var (
	ErrNoTea        = errors.New("No tea leaves available")
	ErrNoWaterTea   = errors.New("No tea water available")
	ErrKettleBroken = errors.New("Tea Kettle broken")
)
