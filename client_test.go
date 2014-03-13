package prevoty

import (
	"fmt"
	"testing"
)

func TestPrevoty(t *testing.T) {
	apiKey := "api key goes here"

	trustedContentConfiguration := "configuration key goes here"
	trustedContentPayload := "the <script>alert('quick brown fox');</script> jumps over the lazy dog & mouse"

	trustedTokenUser := "example_user"
	trustedTokenAction := "example_action"
	trustedTokenTTL := "1000"

	client := NewPrevotyClient(apiKey)

	// Verify the API key
	verified, verifiedErr := client.Verify()
	if verified {
		// Get API key information
		info, infoErr := client.Info()
		if infoErr == nil {
			fmt.Println("Information:", info.Message)
			// Verify configuration
			verification, verifyErr := client.VerifyConfigurationKey(trustedContentConfiguration)
			fmt.Println("Verified rule:", verification, verifyErr)

			// Trusted Content
			tc, filterErr := client.FilterContent(trustedContentPayload, trustedContentConfiguration)
			if filterErr != nil {
				t.Error("Trusted Content filter error:", filterErr)
			}
			fmt.Println("Filtered output:", tc.Output, filterErr)

			// Trusted Token
			generatedToken, generationErr := client.GenerateTimedToken(trustedTokenAction, trustedTokenUser, trustedTokenTTL)
			if generationErr != nil {
				t.Error("Trusted Token generation error:", generationErr)
			} else {
				fmt.Println("Generated token:", generatedToken.Token)
				validatedToken, validationErr := client.ValidateTimedToken(trustedTokenAction, trustedTokenUser, generatedToken.Token)
				if validationErr != nil || !validatedToken.Valid {
					t.Error("Trusted Token validation error:", validationErr)
				} else {
					fmt.Println("Validated token:", generatedToken.Token)
				}
			}
		} else {
			t.Error("Could not get information")
		}
	} else {
		t.Error("API key not verified:", verifiedErr)
	}
}
