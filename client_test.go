package prevoty

import (
	"fmt"
	"github.com/prevoty/prevoty-go"
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
			// Verify rule
			verification, verifyErr := client.VerifyConfigurationKey(ruleKey)
			fmt.Println("Verified rule:", verification, verifyErr)
			// Filter XSS
			result, filterErr := client.Filter(input, ruleKey)
			fmt.Println("Filtered output:", result.Output, filterErr)
		} else {
			fmt.Println("Could not get information")
		}
	} else {
		fmt.Println("API key not verified", verifiedErr)
	}
}
