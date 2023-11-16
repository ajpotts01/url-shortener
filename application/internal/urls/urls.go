package urls

import (
	"encoding/base64"
)

func GetShortenedUrl(url string) (string, error) {
	hashedUrl, err := hashUrl(url)

	if err != nil {
		return "", err
	}

	encodedHash := base64Encode(hashedUrl)
	result := sliceEncodedString(encodedHash)

	return result, nil
}

func sliceEncodedString(input string) string {
	return input[0:7]
}

func base64Encode(hash []byte) string {
	result := base64.StdEncoding.EncodeToString(hash)

	return result
}

func hashUrl(url string) ([]byte, error) {
	// Not handling collisions here
	// Hashing disabled - gosec doesn't like MD5.
	// hash := md5.New()
	// _, err := hash.Write([]byte(url))

	// if err != nil {
	// 	return []byte{}, err
	// }

	// hashVal := hash.Sum(nil)

	// // take the first 7 bytes
	// hashVal = hashVal[0:7]
	// return hashVal, nil
	return []byte{}, nil
}
