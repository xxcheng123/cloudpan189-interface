package client

import "context"

type FamilyListFilesRequest struct {
	PageNum    int    `url:"pageNum"`
	PageSize   int    `url:"pageSize"`
	MediaType  int    `url:"mediaType"`
	FamilyId   String `url:"familyId"`
	FolderId   String `url:"folderId"`
	IconOption int    `url:"iconOption"`
	OrderBy    string `url:"orderBy"`
	Descending bool   `url:"descending"`
}

type WithFamilyListFilesRequestOption func(*FamilyListFilesRequest)

func (c *client) FamilyListFiles(ctx context.Context, familyId String, folderId String, opts ...WithFamilyListFilesRequestOption) (*ListFilesResponse, error) {
	req := &FamilyListFilesRequest{
		PageNum:    1,
		PageSize:   100,
		FamilyId:   familyId,
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

	if err := c.get(ctx, ApiFamilyListFiles, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
