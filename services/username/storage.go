//go:build !mysql

package username

import (
	"github.com/persiangophers/userland/protocol"
)

var storage protocol.UsernameService

func init() {
	panic("[username] Don't support requested storage type")
}
