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

	// ReAdminFediversePre matches the admin fediverse page prefix
	ReAdminFediversePre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s`, PartAdmin, PartFediverse))
	// ReAdminFediverseAccountsPre matches the admin fediverse page prefix
	ReAdminFediverseAccountsPre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s/%s`, PartAdmin, PartFediverse, PartAccounts))
	// ReAdminFediverseInstancesPre matches the admin fediverse page prefix
	ReAdminFediverseInstancesPre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s/%s`, PartAdmin, PartFediverse, PartInstances))

	// ReAdminOauthPre matches the admin oauth page prefix
	ReAdminOauthPre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s`, PartAdmin, PartOauth))
	// ReAdminOauthClientsPre matches the admin oauth clients page prefix
	ReAdminOauthClientsPre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s/%s`, PartAdmin, PartOauth, PartClients))

	// ReAdminSystemPre matches the admin system page prefix
	ReAdminSystemPre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s`, PartAdmin, PartSystem))
	// ReAdminSystemApplicationTokensPre matches the admin system application tokens page prefix
	ReAdminSystemApplicationTokensPre = regexp.MustCompile(fmt.Sprintf(`^?/%s/%s/%s`, PartAdmin, PartSystem, PartApplicationTokens))
)
