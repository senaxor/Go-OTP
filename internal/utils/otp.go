package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Length of OTP (default is 6-digit)
const OTPDigits = 6

// OTP expiry duration (optional if you use an expiry mechanism)
const OTPExpiry = 2 * time.Minute

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateOTP returns a numeric string of length OTPDigits
func GenerateOTP() string {
	min := int64(1)
	for i := 1; i < OTPDigits; i++ {
		min *= 10
	}
	max := min*10 - 1
	num := rand.Int63n(max-min+1) + min
	return fmt.Sprintf("%0*d", OTPDigits, num)
}
