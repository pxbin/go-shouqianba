package shouqianba

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

// GatewayService handles communication with the wap gateway related
// methods of the Shouqianba API.
//
// API docs: https://doc.shouqianba.com/zh-cn/api/wap2.html
type GatewayService service

type GatewayRequest struct {
	// AppID 收钱吧分配的APPID
	TerminalSN   string `url:"terminal_sn" json:"terminal_sn"`               // 收钱吧终端ID
	ClientSN     string `url:"client_sn" json:"client_sn"`                   // 商户系统订单号
	TotalAmount  string `url:"total_amount" json:"total_amount"`             // 订单总金额
	Subject      string `url:"subject" json:"subject"`                       // 交易概述
	Payway       string `url:"payway" json:"payway,omitempty"`               // 支付方式
	Operator     string `url:"operator" json:"operator,omitempty"`           // 门店操作员
	Description  string `url:"description" json:"description,omitempty"`     // 商品详情
	Longitude    string `url:"longitude" json:"longitude,omitempty"`         // 经度
	Latitude     string `url:"latitude" json:"latitude,omitempty"`           // 纬度
	Extended     string `url:"extended" json:"extended,omitempty"`           // 扩展参数集合
	GoodsDetails string `url:"goods_details" json:"goods_details,omitempty"` // 商品详情
	Reflect      string `url:"reflect" json:"reflect,omitempty"`             // 反射参数
	NotifyURL    string `url:"notify_url" json:"notify_url,omitempty"`       // 服务器异步回调 url
	ReturnURL    string `url:"return_url" json:"return_url"`                 // 页面跳转同步通知页面路径
}

func (r *GatewayRequest) signed(terminalKey string) (string, error) {
	v, err := query.Values(r)
	if err != nil {
		return "", err
	}
	v.Del("sign")
	v.Del("sign_type")

	signstr := fmt.Sprintf("%s&key=%s", v.Encode(), terminalKey)
	sum := md5.Sum([]byte(signstr))
	signed := strings.ToUpper(hex.EncodeToString(sum[:]))
	v.Add("sign", signed)
	return v.Encode(), nil
}

// Request 收银吧网关支付 WAP支付收银台
// https://doc.shouqianba.com/zh-cn/api/wap2.html
func (s *GatewayService) Request(ctx context.Context, req *GatewayRequest) (*http.Response, error) {
	req.TerminalSN = s.client.config.TerminalSN

	if s.client.config.NotifyURL != "" {
		req.NotifyURL = s.client.config.NotifyURL
	}
	if s.client.config.ReturnURL != "" {
		req.ReturnURL = s.client.config.ReturnURL
	}

	u, err := url.Parse(apiGatewayBaseURL)
	if err != nil {
		return nil, err
	}

	sv, err := req.signed(s.client.config.TerminalKey)
	if err != nil {
		return nil, err
	}
	u.RawQuery = sv

	resp, err := s.client.Request(ctx, http.MethodGet, u.String(), nil, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
