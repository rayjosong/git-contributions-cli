package scan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getDotFilePath(t *testing.T) {
	filePath := GetDotFilePath()
	assert.Equal(t, "/Users/r.ong.4/.gogitlocalstats", filePath)
}

func Test_sliceContains(t *testing.T) {
	tests := []struct{
		name string
		sliceInput []string
		value string
		expected bool
	}{
		{
			name: "when slice has expected element, then return true.",
			sliceInput: []string{"one", "two", "three"},
			value: "one",
			expected: true,
		},
		{
			name: "when slice does not have expected element, then return false.",
			sliceInput: []string{"one", "two", "three"},
			value: "john wick",
			expected: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := sliceContains(tc.sliceInput, tc.value)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func Test_joinSlices(t *testing.T) {
	existing := []string{"this", "is", "a"}
	new := []string{"slice"}


	got := joinSlices(new, existing)
	assert.Equal(t, []string{"this", "is", "a", "slice"}, got)
}

// func Test_addNewSliceOfElementsToFile(t *testing.T) {
// 	filePath := "/.testFile"
// 	newRepos := []string{
// 		"test/repo_1",
// 		"test/repo_2",
// 		"test/repo_3",
// 	}

// 	addNewSliceElementsToFile(filePath, newRepos)

// 	dat, _ := os.ReadFile("./.testFile")
// 	fmt.Println(string(dat))
// }
