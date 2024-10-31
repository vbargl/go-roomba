package roombapass

import (
	"fmt"

	"github.com/vbargl/go-roomba/pkg/roombapass/internal/gigya"
	"github.com/vbargl/go-roomba/pkg/roombapass/internal/irobotapi"
)

const mebibyte = 1_048_576

func GetPasswordFromCloud(options ...Option) ([]RobotPassword, error) {
	c := getPasswordCloudConfig{}
	c.apply(options...)

	gigyaRes, err := gigya.Login(gigya.LoginRequest{
		LoginID:  c.username,
		Password: c.password,
		APIKey:   c.apikey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to login to Gigya: %w", err)
	}

	res, err := irobotapi.Login(irobotapi.LoginRequest{
		Gigya: irobotapi.LoginRequest_Gigya{
			UID:       gigyaRes.UID,
			Signature: gigyaRes.UIDSignature,
			Timestamp: gigyaRes.SignatureTimestamp,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to login to iRobot API: %w", err)
	}

	robotPasswords := make([]RobotPassword, 0, len(res.Robots))
	for id, robot := range res.Robots {
		robotPasswords = append(robotPasswords, RobotPassword{
			UID:      id,
			Name:     robot.Name,
			Password: robot.Password,
		})
	}


	return robotPasswords, nil
}

type RobotPassword struct {
	UID      string
	Name     string
	Password string
}

type getPasswordCloudConfig struct {
	username string
	password string
	apikey   string
}

func (cfg *getPasswordCloudConfig) apply(opts ...Option) {
	for _, o := range opts {
		o(cfg)
	}
}

type Option func(*getPasswordCloudConfig)

func WithCredentials(username, password string) Option {
	return func(c *getPasswordCloudConfig) {
		c.username = username
		c.password = password
	}
}

func WithAPIKey(apikey string) Option {
	return func(c *getPasswordCloudConfig) {
		c.apikey = apikey
	}
}
