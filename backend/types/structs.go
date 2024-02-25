package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	Email    string
	Password string
	OTP      *string
}

type Claims struct {
	Email string
	jwt.RegisteredClaims
}

type JWT struct {
	AccessToken  string
	RefreshToken string
}

type FilterInput struct {
	Field string
	Value []string
}

type ConnectionInput struct {
	Limit     *int
	After     *string
	SortBy    *string
	Direction *AscOrDesc
	Query     *string
	Filters   []*FilterInput
}

type PageInfo struct {
	StartCursor string
	EndCursor   string
	HasNextPage bool
	Count       int
}

type AppleEdge struct {
	Node   *Apple
	Cursor string
}

type AppleConnection struct {
	Edges    []*AppleEdge
	PageInfo *PageInfo
}
