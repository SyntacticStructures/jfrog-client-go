package services

import (
	"encoding/json"
	"github.com/jfrog/jfrog-client-go/artifactory/auth"
	rthttpclient "github.com/jfrog/jfrog-client-go/artifactory/httpclient"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"net/url"
	"path"
	"strings"
	"strconv"
)

type TokenService struct {
	client     *rthttpclient.ArtifactoryHttpClient
	ArtDetails auth.ArtifactoryDetails
}

func NewTokenService(client *rthttpclient.ArtifactoryHttpClient) *TokenService {
	return &TokenService{client: client}
}

func (ts *TokenService) getArtifactoryDetails() auth.ArtifactoryDetails {
	return ts.ArtDetails
}

func (ts *TokenService) CreateToken(params CreateTokenParams) (CreateTokenResponseData, error) {
	artifactoryUrl := ts.ArtDetails.GetUrl()
	endpointUrl := path.Join("api/security/token")
	requestFullUrl, err := utils.BuildArtifactoryUrl(artifactoryUrl, endpointUrl, make(map[string]string))
	data := buildCreateTokenUrlValues(params)
	httpClientsDetails := ts.getArtifactoryDetails().CreateHttpClientDetails()
	_, body, err := ts.client.SendPostForm(requestFullUrl, data, &httpClientsDetails)
	tokenInfo := CreateTokenResponseData{}
	err = json.Unmarshal(body, &tokenInfo)
	return tokenInfo, err
}

func buildCreateTokenUrlValues(params CreateTokenParams) url.Values {
	data := url.Values{}

	data.Set("refreshable", strconv.FormatBool(params.Refreshable))

	if params.Scope != "" {
		data.Set("scope", params.Scope)
	}

	if params.Username != "" {
		data.Set("username", params.Username)
	}

	if params.GrantType != "" {
		data.Set("grant_type", params.GrantType)
	}

	if len(params.Audience) != 0 {
		// Rest api accepts a space-separated list of strings as audience
		data.Set("audience", strings.Join(params.Audience, " "))
	}

	if params.ExpiresIn != 0 {
		data.Set("expires_in", strconv.Itoa(params.ExpiresIn))
	}

	return data
}

type CreateTokenResponseData struct {
	Scope         string
	Access_Token  string
	Expires_In    int
	Token_Type    string
	Refresh_Token string
}

type CreateTokenParams struct {
	Scope       string
	Username    string
	ExpiresIn   int
	GrantType   string
	Refreshable bool
	Audience    []string
}
