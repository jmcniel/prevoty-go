package prevoty

import (
	"fmt"
	"testing"
)

var (
	apiKey = "api key goes here"

	trustedContentConfiguration = "configuration key goes here"
	trustedContentPayload       = "the <script>alert('quick brown fox');</script> jumps over the lazy dog & mouse"

	trustedTokenUser   = "example_user"
	trustedTokenAction = "example_action"
	trustedTokenTTL    = "1000"

	trustedQueryConfiguration = "configuration key goes here"
	trustedQueryPayload       = "SELECT * FROM Commitments"

	link = "http://google.com"
)

func TestAPIKey(t *testing.T) {
	client := NewPrevotyClient(apiKey)
	verified, verifiedErr := client.Verify()
	if !verified {
		t.Error("API key not verified:", verifiedErr)
	}
}

func TestTrustedContentConfiguration(t *testing.T) {
	client := NewPrevotyClient(apiKey)
	verified, verifiedErr := client.VerifyConfigurationKey(trustedContentConfiguration)
	if !verified {
		t.Error("Trusted Content configuration not verified:", verifiedErr)
	}
}

func TestTrustedContent(t *testing.T) {
	client := NewPrevotyClient(apiKey)
	tc, filterErr := client.FilterContent(trustedContentPayload, trustedContentConfiguration)
	if filterErr != nil {
		t.Error("Trusted Content filter error:", filterErr)
	} else {
		fmt.Println("Filtered output:", tc.Output)
	}
}

func TestTrustedTokenLifecycle(t *testing.T) {
	client := NewPrevotyClient(apiKey)
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
}

func TestTrustedTokenDeletion(t *testing.T) {
	client := NewPrevotyClient(apiKey)
	generatedToken, generationErr := client.GenerateTimedToken(trustedTokenAction, trustedTokenUser, trustedTokenTTL)
	if generationErr != nil {
		t.Error("Trusted Token generation error:", generationErr)
	} else {
		fmt.Println("Generated token:", generatedToken.Token)
		deletedToken, deletionErr := client.DeleteTimedToken(trustedTokenAction, trustedTokenUser, generatedToken.Token)
		if deletionErr != nil || !deletedToken.Deleted {
			t.Error("Trusted Token deletion error:", deletionErr)
		} else {
			fmt.Println("Deleted token:", generatedToken.Token)
		}
	}
}

func TestTrustedQuery(t *testing.T) {
	client := NewPrevotyClient(apiKey)
	la, linkErr := client.AnalyzeLink(link)
	if linkErr != nil {
		t.Error("Link Analysis error:", linkErr)
	} else {
		fmt.Println("Link Analysis result:", la)
	}
}
