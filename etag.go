package etag

import (
	"crypto/sha1"
	"fmt"
)

// Generate etag
func Generate(body []byte, weak bool) string {
	hash := sha1.Sum(body)
	etag := fmt.Sprintf("\"%d-%x\"", int(len(hash)), hash)

	if weak {
		etag = "W/" + etag
	}

	return etag
}
