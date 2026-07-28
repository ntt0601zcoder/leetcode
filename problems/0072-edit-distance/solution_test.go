package editdistance

import "testing"

func TestEditDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{
			name:  "test 1",
			word1: "horse",
			word2: "ros",
			want:  3,
		},
		// {
		// 	name:  "test 2",
		// 	word1: "intention",
		// 	word2: "execution",
		// 	want:  5,
		// },
	}
	for _, tc := range tests {
		got := minDistance(tc.word1, tc.word2)

		if got != tc.want {
			t.Errorf("Test %s failed (%s, %s, %d, %d)", tc.name, tc.word1, tc.word2, got, tc.want)
		}
	}
}
