package configurations_test

import (
	"encoding/json"
	"testing"

	"github.com/free5gc/nrf/pkg/configurations"
	"github.com/stretchr/testify/assert"
)

func TestDynamicConfigJSON(t *testing.T) {
	// Create a sample DynamicConfig object

	type testcaset struct {
		name  string
		input configurations.DynamicConfig
	}

	var testcases = []testcaset{
		{
			name: "Complete Config",
			input: configurations.DynamicConfig{
				Sbi: configurations.Sbi{
					OAuth2: configurations.OAuth2{
						Enable: true,
						Period: 3600,
					},
				},
			},
		},
		{
			name:  "Empty Config",
			input: configurations.DynamicConfig{},
		},
		{
			name: "No OAuth2 Config",
			input: configurations.DynamicConfig{
				Sbi: configurations.Sbi{},
			},
		},
		{
			name: "Only OAuth2 Period",
			input: configurations.DynamicConfig{
				Sbi: configurations.Sbi{
					OAuth2: configurations.OAuth2{
						Period: 163163,
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		originalConfig := tc.input
		expectedConfig := tc.input

		// Marshal the object to JSON
		jsonData, err := json.Marshal(originalConfig)
		if err != nil {
			t.Fatalf("Failed to marshal DynamicConfig: %v", err)
		}

		// Unmarshal the JSON back to a DynamicConfig object
		var unmarshalledConfig configurations.DynamicConfig
		err = json.Unmarshal(jsonData, &unmarshalledConfig)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON to DynamicConfig: %v", err)
		}

		// Compare the original and expected objects
		assert.Equal(t, originalConfig, expectedConfig)
	}
}
