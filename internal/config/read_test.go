package config

import "testing"

func TestFileExists(t *testing.T) {

	configPath, _ := getConfigFilePath()
	cases := []struct {
		filePath string
		expected bool
	}{
		{
			filePath: "read_test.go",
			expected: true,
		},
		{
			filePath: configPath,
			expected: true,
		},
	}

	for i := 0; i < len(cases); i++ {
		actual := checkFileExists(cases[i].filePath)

		if actual != cases[i].expected {
			t.Errorf("Actual:\t%t != Expected:\t%t", actual, cases[i].expected)
		}
	}

}

func TestReadConfig(t *testing.T) {
	actual, _ := Read()

	expected := "postgres://example"

	if actual.DbURL != expected {
		t.Errorf("Actual:\t%s != Expected:\t%s", actual, expected)
	}

}
