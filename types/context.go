package types

import "github.com/gofrs/uuid"

const UserCtxName = "user"

type UserContext struct {
	UserID      uuid.UUID `json:"uid"`
	Username    string    `json:"email"`
	SocialName  string    `json:"socialName"`
	Avatar      string    `json:"avatar"`
	DisplayName string    `json:"displayName"`
	SystemRole  string    `json:"role"`
}
