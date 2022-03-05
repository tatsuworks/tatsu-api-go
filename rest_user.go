package tatsu_api

import (
	"context"
	"golang.org/x/xerrors"
)

func (rc *restClient) getUserProfile(ctx context.Context, userID string) (*User, error) {
	// Ensure user ID is not empty.
	if userID == "" {
		return nil, xerrors.New("user id was empty")
	}

	var user *User

	// Make request.
	err := rc.get(ctx, getUserProfile(userID), &user)

	return user, err
}
