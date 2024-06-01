package asteroid

import (
	"errors"
	"time"
	"regexp"
)

var (
	ErrInvalidName = errors.New("invalid name")
	ErrInvalidDiameter = errors.New("invalid diameter")
	ErrInvalidDateFormat = errors.New("invalid date format")
	ErrInvalidDate = errors.New("invalid date")
	ErrInvalidDistance = errors.New("invalid distance")
)

func ValidateAsteroid(asteroid *Asteroid) error {
	if asteroid.Name == "" {
		return ErrInvalidName
	}
	if asteroid.Diameter <= 0 {
		return ErrInvalidDiameter
	}
	if err := validateDate(asteroid.DiscoveryDate); err != nil {
		return err
	}
	for _, distance := range asteroid.Distances {
		if err := validateDate(distance.Date); err != nil {
			return err
		}
		if distance.Distance <= 0 {
			return ErrInvalidDistance
		}
	}
	return nil
}

func validateDate(date string) error {
	match, _ := regexp.MatchString(`^\d{2}-\d{2}-\d{4}$`, date)
	if !match {
		return ErrInvalidDateFormat
	}
	_, err := time.Parse("02-01-2006", date)
	return err
}
