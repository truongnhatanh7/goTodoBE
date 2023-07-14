package common

import "fmt"

const (
	CurrentUser = "current_user"
)

type DbType int

const (
	DbTypeItem DbType = 1
	DbTypeUser DbType = 2
)

const (
	PluginDBMain = "mysql"
	PluginJWT    = "jwt"
	PluginPubSub = "pubsub"

	TopicUserLikedItem   = "TopicUserLikedItem"
	TopicUserUnlikedItem = "TopicUserUnlikedItem"
)

func Recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}

type TokenPayload struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p TokenPayload) UserId() int {
	return p.UId
}

func (p TokenPayload) Role() string {
	return p.URole
}

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func IsAdmin(requester Requester) bool {
	return requester.GetRole() == "admin" || requester.GetRole() == "mod"
}