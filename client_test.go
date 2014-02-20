package prevoty

import (
	"fmt"
	"testing"
)

func TestPrevoty(t *testing.T) {
	apiKey := "api key goes here"
	configurationKey := "configuration key goes here"
	input := "the <script>alert('quick brown fox');</script> jumps over the lazy dog & mouse"

	client := NewPrevotyClient(apiKey)

	// Verify the API key
	verified, verifiedErr := client.Verify()
	if verified {
		// Get API key information
		info, infoErr := client.Info()
		if infoErr == nil {
			fmt.Println("Information:", info.Message)
			// Verify configuration
			verification, verifyErr := client.VerifyConfigurationKey(configurationKey)
			fmt.Println("Verified rule:", verification, verifyErr)
			// Filter XSS
			result, filterErr := client.Filter(input, configurationKey)
			fmt.Println("Filtered output:", result.Output, filterErr)
		} else {
			t.Error("Could not get information")
		}
	} else {
		t.Error("API key not verified", verifiedErr)
	}
}
