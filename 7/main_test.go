package main

import "testing"

func TestIsPrimeTrue(t *testing.T) {
	got := isPrime(11)
	if got != true {
		t.Errorf("isPrime(11) = %t; want true", got)
	}
}

func TestIsPrimeFalse(t *testing.T) {
	got := isPrime(8)
	if got != false {
		t.Errorf("isPrime(8) = %t; want false", got)
	}
}
