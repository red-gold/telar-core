package types

import "github.com/gofrs/uuid"

const UserCtxName = "user"

type UserContext struct {
	UserID      uuid.UUID `json:"uid"`
	Username    string    `json:"email"`
	DisplayName string    `json:"displayName"`
	SocialName  string    `json:"socialName"`
	Avatar      string    `json:"avatar"`
	Banner      string    `json:"banner"`
	TagLine     string    `json:"tagLine"`
	SystemRole  string    `json:"role"`
	CreatedDate int64     `json:"createdDate"`
}
