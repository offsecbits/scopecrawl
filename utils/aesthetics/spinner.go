// spinner.go
package aesthetics

import (
	"fmt"
	"time"
)
// Spinner struct holds the spinner state
type Spinner struct {
	stop chan bool
}

// StartSpinner starts a spinner that runs until the task is complete
func StartSpinner() *Spinner {
	spinSymbols := []string{"x", "-", "+", "*", "+", "-", "x"}
	spinner := &Spinner{stop: make(chan bool)}

	// Start a goroutine for the spinner
	go func() {
		for {
			select {
			case <-spinner.stop:
				// Stop the spinner when the signal is received
				fmt.Print("\rDone!        \n") // Clear the spinner
				return
			default:
				for _, symbol := range spinSymbols {
					// Print the spinner symbol and overwrite it
					fmt.Printf("\r \r \r%s %s %s ", symbol, symbol, symbol)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()

	return spinner
}

// Stop stops the spinner
func (s *Spinner) Stop() {
	s.stop <- true
}
