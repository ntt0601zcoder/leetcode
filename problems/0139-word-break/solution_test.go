package wordbreak

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{}

	for _, tc := range tests {
		got := wordBreak(tc.s, tc.wordDict)

		if got != tc.want {
			t.Errorf("Testcase %s failed (s: %s, wordDict: %s)", tc.name, tc.s, tc.wordDict)
		}
	}

}
