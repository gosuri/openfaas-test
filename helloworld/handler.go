package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Hello, OpenFaaS if freaking cool! You said: %s", string(req))
}
