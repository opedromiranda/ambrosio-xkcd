package ambrosio_xkcd

import (
	"regexp"
	"testing"
)

func TestSpecificHandler(t *testing.T) {
	var tests = []struct {
		s    []string
		want string
	}{
		{[]string{"xkcd 999999999", "999999999"}, ERROR_MESSAGE},
	}

	for _, c := range tests {
		got, _ := Specific.Handler(c.s)
		if got != c.want {
			t.Errorf("Specific.Handler(%q) == %q, want %q", c.s, got, c.want)
		}
	}

}

func TestLatestPattern(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"xkcd latest", true},

		{"xck lates", false},
		{"xkcd 123", false},
		{"xkcd random", false},
	}

	for _, c := range tests {
		got, _ := regexp.MatchString(Latest.Pattern, c.s)
		if got != c.want {
			t.Errorf("regexp.MatchString(%q, %q) == %t, want %t", Latest.Pattern, c.s, got, c.want)
		}
	}
}

func TestSpecificPattern(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"xkcd 1", true},
		{"xkcd 99", true},
		{"xkcd 1231299", true},

		{"xck lates", false},
		{"xkcd latest", false},
		{"xkcd random", false},
	}

	for _, c := range tests {
		got, _ := regexp.MatchString(Specific.Pattern, c.s)
		if got != c.want {
			t.Errorf("regexp.MatchString(%q, %q) == %t, want %t", Specific.Pattern, c.s, got, c.want)
		}
	}
}
