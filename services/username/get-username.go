package username

import "github.com/persiangophers/userland/protocol"

var GetUsernameService = getUsernameService{
	path: "/apis/username/get-username",
}

type getUsernameService struct {
	path string
}

func (s *getUsernameService) Process(req SetUsernameReq) (protocol.Username, error) {
	return storage.Last(req.UserID)
}

type SetUsernameReq struct {
	UserID [16]byte
}
