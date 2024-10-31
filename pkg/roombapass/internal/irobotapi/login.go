package irobotapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/vbargl/go-roomba/pkg/roombapass/httputil"
)

func Login(req LoginRequest) (res LoginResponse, err error) {
	err = req.validate()
	if err != nil {
		return res, err
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(req)
	if err != nil {
		return res, err
	}

	resp, err := http.Post("https://unauth2.prod.iot.irobotapi.com/v2/login", "application/json", &buf)
	if err := httputil.CheckAndLimitBody(resp, err); err != nil {
		return res, err
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	return res, err
}

type LoginRequest struct {
	AppID                string             `json:"app_id"`
	AssumeRobotOwnership bool               `json:"assume_robot_ownership"`
	Gigya                LoginRequest_Gigya `json:"gigya"`
}

func (req *LoginRequest) validate() error {
	if req.AppID == "" {
		req.AppID = "ANDROID-C7FB240E-DF34-42D7-AE4E-A8C17079A294"
	}
	if err := req.Gigya.validate(); err != nil {
		return fmt.Errorf("Gigya.%w", err)
	}
	return nil
}

type LoginRequest_Gigya struct {
	UID       string `json:"uid"`
	Signature string `json:"signature"`
	Timestamp string `json:"timestamp"`
}

func (req *LoginRequest_Gigya) validate() error {
	if req.UID == "" {
		return errors.New("UID is required")
	}
	if req.Signature == "" {
		return errors.New("Signature is required")
	}
	if req.Timestamp == "" {
		return errors.New("Timestamp is required")
	}
	return nil
}

type LoginResponse struct {
	Credentials       LoginResponse_Credentials      `json:"credentials"`
	Robots            map[string]LoginResponse_Robot `json:"robots"`
	IotToken          string                         `json:"iot_token"`
	IotClientID       string                         `json:"iot_clientid"`
	IotSignature      string                         `json:"iot_signature"`
	IotAuthorizerName string                         `json:"iot_authorizer_name"`
}

type LoginResponse_Credentials struct {
	AccessKeyID  string    `json:"AccessKeyId"`
	SecretKey    string    `json:"SecretKey"`
	SessionToken string    `json:"SessionToken"`
	Expiration   time.Time `json:"Expiration"`
	CognitoID    string    `json:"CognitoId"`
}

type LoginResponse_Robot struct {
	Password     string                         `json:"password"`
	SKU          string                         `json:"sku"`
	SoftwareVer  string                         `json:"softwareVer"`
	Name         string                         `json:"name"`
	Capabilities LoginResponse_Capabilities     `json:"cap"`
	DigiCap      LoginResponse_DigiCapabilities `json:"digiCap"`
	SvcDeplID    string                         `json:"svcDeplId"`
	UserCert     bool                           `json:"user_cert"`
}

type LoginResponse_Capabilities struct {
	BinFullDetect int `json:"binFullDetect"`
	Omode         int `json:"oMode"`
	DockComm      int `json:"dockComm"`
	Edge          int `json:"edge"`
	Maps          int `json:"maps"`
	Pmaps         int `json:"pmaps"`
	MC            int `json:"mc"`
	Tline         int `json:"tLine"`
	Area          int `json:"area"`
	Eco           int `json:"eco"`
	MultiPass     int `json:"multiPass"`
	Team          int `json:"team"`
	Pp            int `json:"pp"`
	Lang          int `json:"lang"`
	FiveGhz       int `json:"5ghz"`
	Prov          int `json:"prov"`
	Sched         int `json:"sched"`
	SvcConf       int `json:"svcConf"`
	Ota           int `json:"ota"`
	Log           int `json:"log"`
	LangOta       int `json:"langOta"`
	AddOnHw       int `json:"addOnHw"`
}

type LoginResponse_DigiCapabilities struct {
	AppVer int `json:"appVer"`
}
