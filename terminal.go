package shouqianba

import (
	"context"
	"fmt"
	"net/http"
)

// TerminalService handles communication with the terminal related
// methods of the Shouqianba API.
//
// API docs: https://doc.shouqianba.com/zh-cn/api/interface/activate.html
type TerminalService service

type TerminalResponse struct {
	BizResponse struct {
		TerminalKey string `json:"terminal_key"`
		TerminalSN  string `json:"terminal_sn"`
	} `json:"biz_response"`
	ResultCode string `json:"result_code"`
	ErrCode    string `json:"error_code,omitempty"`
	ErrMessage string `json:"error_message,omitempty"`
}

type TerminalActivateRequest struct {
	AppID      string `json:"app_id"`
	Code       string `json:"code"`
	DeviceID   string `json:"device_id"`
	ClientSN   string `json:"client_sn,omitempty"`
	Name       string `json:"name,omitempty"`
	OsInfo     string `json:"os_info,omitempty"`
	SDKVersion string `json:"sdk_version,omitempty"`
}

// Activate 收钱吧终端激活接口
// https://doc.shouqianba.com/zh-cn/api/interface/activate.html
func (s *TerminalService) Activate(ctx context.Context) (*TerminalResponse, *http.Response, error) {
	u := baseURL + "/terminal/activate"

	req := &TerminalActivateRequest{
		AppID:    s.client.config.AppID,
		Code:     s.client.config.Code,
		DeviceID: s.client.config.DeviceID,
	}

	signed, err := sign(req, s.client.config.VendorSN, s.client.config.VendorKey)
	if err != nil {
		return nil, nil, fmt.Errorf("md5: activate terminal sign error : %v", err)
	}

	result := new(TerminalResponse)
	resp, err := s.client.Request(ctx, http.MethodPost, u, req, result, WithAuthentication(signed))
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

type TerminalCheckInRequest struct {
	TerminalSn string `json:"terminal_sn"`
	DeviceID   string `json:"device_id"`
	OsInfo     string `json:"os_info,omitempty"`
	SDKVersion string `json:"sdk_version,omitempty"`
}

// CheckIn 收钱吧终端签到接口
// https://doc.shouqianba.com/zh-cn/api/interface/checkin.html
func (s *TerminalService) CheckIn(ctx context.Context) (*TerminalResponse, *http.Response, error) {
	u := baseURL + "/terminal/checkin"

	req := &TerminalCheckInRequest{
		DeviceID:   s.client.config.DeviceID,
		TerminalSn: s.client.config.TerminalSN,
	}

	signed, err := sign(req, s.client.config.TerminalSN, s.client.config.TerminalKey)

	if err != nil {
		return nil, nil, fmt.Errorf("md5: checkin terminal sign error : %v", err)
	}

	result := new(TerminalResponse)
	resp, err := s.client.Request(ctx, http.MethodPost, u, req, result, WithAuthentication(signed))
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
