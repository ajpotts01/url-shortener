package urls

import (
	"testing"
)

// func TestHash(t *testing.T) {
// 	input := "https://ajpcloudblog.com"

// 	// hashUrl will return []byte so it can be base64 encoded later.
// 	// This can't be compared to the SHA-256 hash we expect
// 	// As it's actually a stringified version of the bytes
// 	// fmt.Sprintf is necessary here
// 	hashBytes, err := hashUrl(input)
// 	got := fmt.Sprintf("%x", hashBytes)
// 	want := "1c4a873ddc5063c2914d7e364f75905c75b185aa792e116dba512076f10664d9"

// 	if err != nil {
// 		t.Errorf("Failed to hash URL: %v", err)
// 	}

// 	if got != want {
// 		t.Errorf("got %q, want %q", got, want)
// 	}
// }

func TestGetShortenedUrl(t *testing.T) {
	inputSimple := "https://ajpcloudblog.com"
	inputComplex := "https://blog.boot.dev/clean-code/youre-not-qualified-for-tech-opinions/"

	// To identify expected values for this test case, generate the base64 using OpenSSL in Bash:
	// echo -n foo | openssl dgst -binary -sha256 | openssl base64
	// https://stackoverflow.com/questions/3358420/generating-a-sha-256-hash-from-the-linux-command-line
	// If you got an example from an SHA-256 generator online, you will likely get given a STRING format
	// of the raw bytes. This will not convert to the expected base64 used by the URL shortening method.
	want := "HEqHPdx"
	got, err := GetShortenedUrl(inputSimple)

	if err != nil {
		t.Errorf("Failed to shorten URL: %v", err)
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	want = "/evxfyN"

	got, err = GetShortenedUrl(inputComplex)

	if err != nil {
		t.Errorf("Failed to shorten URL: %v", err)
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
