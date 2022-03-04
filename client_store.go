package tatsu_api

import (
	"context"

	"golang.org/x/xerrors"
)

// GetStoreListing wraps around GetStoreListingWithContext using the background context.
func (c *Client) GetStoreListing(listingID string) (*StoreListing, error) {
	return c.GetStoreListingWithContext(context.Background(), listingID)
}

// GetStoreListingWithContext gets a store listing using the specified context.
func (c *Client) GetStoreListingWithContext(ctx context.Context, listingID string) (*StoreListing, error) {
	listing, err := c.rc.getStoreListing(ctx, listingID)
	if err != nil {
		return nil, xerrors.Errorf("failed to get store listing: %w", err)
	}

	return listing, nil
}
