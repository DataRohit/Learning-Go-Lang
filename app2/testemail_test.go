package app2

import "testing"

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Errorf("'hello' is not a valid email, but no error was returned")
	}

	_, err = IsEmail("rohit.vilas.ingole@gmail.com")
	if err != nil {
		t.Errorf("'rohit.vilas.ingole@gmail.com' is a valid email, but an error was returned: %v", err)
	}
}
