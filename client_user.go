package tatsu_api

import (
	"context"

	"golang.org/x/xerrors"
)

// GetUserProfile wraps around GetUserProfileWithContext using the background context.
func (c *Client) GetUserProfile(userID string) (*User, error) {
	return c.GetUserProfileWithContext(context.Background(), userID)
}

// GetUserProfileWithContext gets a user's profile using the specified context.
func (c *Client) GetUserProfileWithContext(ctx context.Context, userID string) (*User, error) {
	user, err := c.rc.getUserProfile(ctx, userID)
	if err != nil {
		return nil, xerrors.Errorf("failed to get user profile: %w", err)
	}

	return user, nil
}
