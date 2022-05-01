package path

import (
	"fmt"
	"regexp"
)

const (
	// reToken is regex to match a token
	reToken = `[a-zA-Z0-9_]{16,}`
)

var (
	// ReAdmin matches the admin page
	ReAdmin = regexp.MustCompile(fmt.Sprintf(`^?/%s$`, PartAdmin))
	// ReAdminOauthPre matches the admin oauth page prefix
	ReAdminOauthPre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s`, PartAdmin, PartOauth))
	// ReAdminOauthClientsPre matches the admin oauth clients page prefix
	ReAdminOauthClientsPre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s/%s`, PartAdmin, PartOauth, PartClients))
)
