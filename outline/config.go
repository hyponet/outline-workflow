package outline

import (
	"os"
	"strings"
)

var (
	Host     string
	ApiToken string
)

func init() {
	Host = strings.TrimSpace(os.Getenv("OUTLINE_HOST"))
	ApiToken = strings.TrimSpace(os.Getenv("OUTLINE_API_TOKEN"))
}
