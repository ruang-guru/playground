package graderapiclient

import (
	"context"

	"github.com/ruang-guru/playground/cli/assignmentmanager/entity"
)

type GraderApiClient struct {
	ApiUrl       string
	Email        string
	Password     string
	AccessToken  string
	RefreshToken string
}

var GraderApiUrl string = "https://gw.ruangguru.com/api/v3/rg-grader"

func NewGraderApiClientFromEmailAndToken(ctx context.Context, email string, tokenData *entity.TokenData) *GraderApiClient {
	return &GraderApiClient{
		Email:        email,
		AccessToken:  tokenData.AccessToken,
		RefreshToken: tokenData.RefreshToken,
	}
}

func NewGraderApiClientFromEmailAndPassword(ctx context.Context, email, password string) *GraderApiClient {
	return &GraderApiClient{
		Email:    email,
		Password: password,
	}
}
