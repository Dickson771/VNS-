package utils

import (
	"fmt"
	"time"
)

func RetryFunction(f func() error, retries int, delay time.Duration) error {
	var err error
	for i := 0; i < retries; i++ {
		err = f()
		if err == nil {
			return nil
		}
		time.Sleep(delay)
	}
	return fmt.Errorf("failed after %d retries: %w", retries, err)
}
