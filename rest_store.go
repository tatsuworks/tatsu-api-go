package tatsu_api

import (
	"context"

	"golang.org/x/xerrors"
)

func (rc *restClient) getStoreListing(ctx context.Context, listingID string) (*StoreListing, error) {
	// Ensure listing ID is not empty.
	if listingID == "" {
		return nil, xerrors.New("listing id was empty")
	}

	var listing *StoreListing

	// Make request.
	err := rc.get(ctx, getStoreListing(listingID), &listing)

	return listing, err
}
