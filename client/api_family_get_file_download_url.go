package client

import (
	"context"
)

type FamilyGetFileDownloadResponse struct {
	FileDownloadUrl string `json:"fileDownloadUrl"`
}

type FamilyGetFileDownloadRequest struct {
	FileId   String `url:"fileId"`
	FamilyId String `url:"familyId"`
}

type FamilyGetFileDownloadOption func(*FamilyGetFileDownloadRequest)

func (c *client) FamilyGetFileDownload(ctx context.Context, familyId String, fileId String, opts ...FamilyGetFileDownloadOption) (*FamilyGetFileDownloadResponse, error) {
	req := &FamilyGetFileDownloadRequest{
		FamilyId: familyId,
		FileId:   fileId,
	}

	for _, opt := range opts {
		opt(req)
	}

	var (
		result = new(FamilyGetFileDownloadResponse)
	)

	if err := c.get(ctx, ApiFamilyGetFileDownloadURL, req, result); err != nil {
		return nil, err
	}

	return result, nil
}
