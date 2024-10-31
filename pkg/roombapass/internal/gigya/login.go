package gigya

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/vbargl/go-roomba/pkg/roombapass/httputil"
)

const defaultApiKey = "3_rWtvxmUKwgOzu3AUPTMLnM46lj-LxURGflmu5PcE_sGptTbD-wMeshVbLvYpq01K"

func Login(req LoginRequest) (res LoginResponse, err error) {
	req.withDefaults()

	query := make(url.Values)
	query.Add("apiKey", req.APIKey)
	query.Add("loginID", req.LoginID)
	query.Add("password", req.Password)
	query.Add("format", req.Format)
	query.Add("targetEnv", req.TargetEnv)

	resp, err := http.Post("https://accounts.us1.gigya.com/accounts.login?"+query.Encode(), "", nil)
	if err := httputil.CheckAndLimitBody(resp, err); err != nil {
		return LoginResponse{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	return res, err
}

type LoginRequest struct {
	APIKey    string
	LoginID   string
	Password  string
	Format    string
	TargetEnv string
}

func (req *LoginRequest) withDefaults() {
	if req.APIKey == "" {
		req.APIKey = defaultApiKey
	}
	if req.Format == "" {
		req.Format = "json"
	}
	if req.TargetEnv == "" {
		req.TargetEnv = "mobile"
	}
}

type LoginResponse struct {
	CallID                     string                    `json:"callId"`
	ErrorCode                  int                       `json:"errorCode"`
	APIVersion                 int                       `json:"apiVersion"`
	StatusCode                 int                       `json:"statusCode"`
	StatusReason               string                    `json:"statusReason"`
	Time                       time.Time                 `json:"time"`
	RegisteredTimestamp        int64                     `json:"registeredTimestamp"`
	UID                        string                    `json:"UID"`
	UIDSignature               string                    `json:"UIDSignature"`
	SignatureTimestamp         string                    `json:"signatureTimestamp"`
	Created                    time.Time                 `json:"created"`
	CreatedTimestamp           int64                     `json:"createdTimestamp"`
	IsActive                   bool                      `json:"isActive"`
	IsRegistered               bool                      `json:"isRegistered"`
	IsVerified                 bool                      `json:"isVerified"`
	LastLogin                  time.Time                 `json:"lastLogin"`
	LastLoginTimestamp         int64                     `json:"lastLoginTimestamp"`
	LastUpdated                time.Time                 `json:"lastUpdated"`
	LastUpdatedTimestamp       int64                     `json:"lastUpdatedTimestamp"`
	LoginProvider              string                    `json:"loginProvider"`
	OldestDataUpdated          time.Time                 `json:"oldestDataUpdated"`
	OldestDataUpdatedTimestamp int64                     `json:"oldestDataUpdatedTimestamp"`
	Profile                    LoginResponse_Profile     `json:"profile"`
	Registered                 time.Time                 `json:"registered"`
	SocialProviders            string                    `json:"socialProviders"`
	Verified                   time.Time                 `json:"verified"`
	VerifiedTimestamp          int64                     `json:"verifiedTimestamp"`
	NewUser                    bool                      `json:"newUser"`
	SessionInfo                LoginResponse_SessionInfo `json:"sessionInfo"`
}

type LoginResponse_Profile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Country   string `json:"country"`
	Email     string `json:"email"`
}

type LoginResponse_SessionInfo struct {
	SessionToken  string `json:"sessionToken"`
	SessionSecret string `json:"sessionSecret"`
	ExpiresIn     string `json:"expires_in"`
}
