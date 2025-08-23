package client

import "context"

type ListFilesRequest struct {
	PageNum    int    `url:"pageNum"`
	PageSize   int    `url:"pageSize"`
	MediaType  int    `url:"mediaType"`
	FolderId   String `url:"folderId"`
	IconOption int    `url:"iconOption"`
	OrderBy    string `url:"orderBy"`
	Descending bool   `url:"descending"`
}

type WithListFilesRequestOption func(*ListFilesRequest)

func (c *client) ListFiles(ctx context.Context, folderId String, opts ...WithListFilesRequestOption) (*ListFilesResponse, error) {
	req := &ListFilesRequest{
		PageNum:    1,
		PageSize:   100,
		FolderId:   folderId,
		OrderBy:    "lastOpTime",
		IconOption: 5,
		Descending: true,
	}

	for _, opt := range opts {
		opt(req)
	}

	var (
		result = new(ListFilesResponse)
	)

	if err := c.get(ctx, ApiListFiles, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
