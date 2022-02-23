//go:build mysql

package username

import (
	"time"

	"github.com/persiangophers/userland/external"
	"github.com/persiangophers/userland/protocol"
)

var storage mysql
var _ protocol.UsernameService = &storage

func init() {
	// TODO::: check desire table exist or create it
}

type mysql struct{}

func (m *mysql) Create(username protocol.Username) (err error) {
	const query = ""
	_, err = external.MySQL.Exec(query, username.UserID(), username.Username(), username.Status())
	return
}

func (m *mysql) Find(username string) (protocol.Username, error)               { return nil, nil }
func (m *mysql) Last(userID [16]byte) (protocol.Username, error)               { return nil, nil }
func (m *mysql) Lasts(userID [16]byte, count int) ([]protocol.Username, error) { return nil, nil }
func (m *mysql) Meantime(userID [16]byte, start, end time.Time) ([]protocol.Username, error) {
	return nil, nil
}
func (m *mysql) CountVersionUntil(userID [16]byte, until time.Time) (int, error) { return 0, nil }
