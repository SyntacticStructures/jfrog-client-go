package services

import (
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory/auth"
	rthttpclient "github.com/jfrog/jfrog-client-go/artifactory/httpclient"
)

type TokenService struct {
	client     *rthttpclient.ArtifactoryHttpClient
	ArtDetails auth.ArtifactoryDetails
	DryRun     bool
}

func (ts *TokenService) getArtifactoryDetails() auth.ArtifactoryDetails {
	return ts.ArtDetails
}

func NewTokenService(client *rthttpclient.ArtifactoryHttpClient) *TokenService {
	return &TokenService{client: client}
}

func (ts *TokenService) isDryRun() bool {
	return ts.DryRun
}

func (ts *TokenService) CreateToken() {
	fmt.Print("TESTING!!!")
}

type CreateTokenParams struct {
	Scope             string
	Username          string
}