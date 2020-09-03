package main

import "time"

// SetTimezone Seta timezone na aplicação
func SetTimezone(timeZone string) error {
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		return err
	}

	time.Now().In(location)

	return nil
}
