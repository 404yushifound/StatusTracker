package provider

import (
	"fmt"
	"sync"
	"time"
)

type MockEmailProvider struct {
	failureCount   int
	circuitOpen    bool
	mutex          sync.Mutex
	lastFailedTime time.Time
}

func NewMockEmailProvider() *MockEmailProvider {
	return &MockEmailProvider{
		failureCount: 0,
		circuitOpen:  false,
	}
}

func (m *MockEmailProvider) Send(email string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.circuitOpen {
		// Check if cooldown period has passed
		if time.Since(m.lastFailedTime) < 10*time.Second {
			fmt.Printf("[CIRCUIT OPEN] Email sending blocked for %s\n", email)
			return false
		}
		// Cooldown over → close the circuit
		fmt.Println("[CIRCUIT CLOSED] Resuming email sending.")
		m.circuitOpen = false
		m.failureCount = 0
	}

	// Simulate sending email
	fmt.Printf("Sending email to %s...\n", email)
	time.Sleep(1 * time.Second)

	// FAKE FAILURE: Fail email if it contains "fail"
	if email == "fail@example.com" {
		fmt.Println("[SIMULATED FAILURE]")
		m.failureCount++
		if m.failureCount >= 3 {
			m.circuitOpen = true
			m.lastFailedTime = time.Now()
			fmt.Println("[CIRCUIT OPENED] Too many failures.")
		}
		return false
	}

	// Reset failure count on success
	m.failureCount = 0
	return true
}
