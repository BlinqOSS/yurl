package yurllib

import (
	"io"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestEvaluateAASA(t *testing.T) {
	t.Skipf("FIXME: this test should pass")

	jsonFile, err := os.Open("testdata/apple-app-site-association")
	if err != nil {
		t.Fatal("failed to open test file")
	}

	fileContent, err := io.ReadAll(jsonFile)
	if err != nil {
		t.Fatal("failed to read test file")
	}

	contentType := []string{"application/json"}
	bundleIdentifier := "com.example.app"
	teamIdentifier := "ABCDE12345"

	res, errs := evaluateAASA(fileContent, contentType, bundleIdentifier, teamIdentifier)
	if len(errs) > 0 {
		spew.Dump(errs)
		t.Error("Got errors: %w", errs)
	}

	// TODO: validate res
	spew.Dump(res)
}

func TestEvaluateAASA_Invalid(t *testing.T) {
	cases := []struct {
		name        string
		fileContent []byte
		errs        []error
	}{
		{
			name:        "empty_file_content",
			fileContent: []byte{},
		},
		{
			name: "app_id-empty",
			fileContent: []byte(`
			{
				"applinks": {
					"details": [
						{
							"appID": "",
						}
					]
				}
			}
			`),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			contentType := []string{"application/json"}
			bundleIdentifier := "com.example.app"
			teamIdentifier := "ABCDE12345"

			_, errs := evaluateAASA(c.fileContent, contentType, bundleIdentifier, teamIdentifier)
			assert.True(t, len(errs) > 0)

			// spew.Dump(errs)
		})
	}
}
