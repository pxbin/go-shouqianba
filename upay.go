package shouqianba

import (
	"context"
	"net/http"
)

// UPayService handles communication with the upayV2 related
// methods of the Shouqianba API.
//
// API docs: https://doc.shouqianba.com/zh-cn/api/wap2.html
type UPayService service

// PreCreateRequest represents the request for create payment request
type PreCreateRequest struct {
	TerminalSN   string       `json:"terminal_sn"`             // 收钱吧终端ID 	不超过32位的纯数字
	ClientSN     string       `json:"client_sn"`               // 商户系统订单号,必须在商户系统内唯一；且长度不超过32字节
	TotalAmount  string       `json:"total_amount"`            // 交易总金额,以分为单位,不超过10位纯数字字符串,超过1亿元的收款请使用银行转账
	Payway       string       `json:"payway,omitempty"`        // 支付方式
	SubPayway    string       `json:"sub_payway,omitempty"`    // 二级支付方式 内容为数字的字符串，如果要使用WAP支付，则必须传 "3"；使用小程序支付，则必须传"4"
	PayerUID     string       `json:"payer_uid,omitempty"`     // 付款人id 消费者在支付通道的唯一id，wap支付，小程序支付必传 ，微信WAP支付必须传open_id,支付宝WAP支付必传用户授权的userId
	DynamicID    string       `json:"dynamic_id,omitempty"`    // 条码内容
	Subject      string       `json:"subject"`                 // 交易简介
	Operator     string       `json:"operator,omitempty"`      // 门店操作员
	Description  string       `json:"description,omitempty"`   // 交易详情
	Longitude    string       `json:"longitude,omitempty"`     // 经度
	Latitude     string       `json:"latitude,omitempty"`      // 纬度
	DeviceID     string       `json:"device_id,omitempty"`     // 设备指纹
	Extended     interface{}  `json:"extended,omitempty"`      // 扩展参数集合
	GoodsDetails []GoodDetail `json:"goods_details,omitempty"` // 商品详情
	Reflect      string       `json:"reflect,omitempty"`       // 业务反射参数
	NotifyURL    string       `json:"notify_url,omitempty"`    // 支付回调的地址
}

// Payment 收钱吧支付接口
// https://doc.shouqianba.com/zh-cn/api/interface/pay.html
func (s *UPayService) Payment(ctx context.Context, req *PreCreateRequest, opts ...RequestOption) (*ApiResponse, *http.Response, error) {
	u := baseURL + "/upay/v2/pay"
	req.TerminalSN = s.client.config.TerminalSN
	req.Subject = s.client.config.subject
	req.Operator = s.client.config.operator

	if s.client.config.NotifyURL != "" {
		req.NotifyURL = s.client.config.NotifyURL
	}

	signed, err := sign(req, s.client.config.TerminalSN, s.client.config.TerminalKey)
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts, WithAuthentication(signed))

	result := new(ApiResponse)
	resp, err := s.client.Request(ctx, http.MethodPost, u, req, result, opts...)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// Precreate 收钱吧预下单接口
// https://doc.shouqianba.com/zh-cn/api/interface/precreate.html
func (s *UPayService) Precreate(ctx context.Context, req *PreCreateRequest, opts ...RequestOption) (*ApiResponse, *http.Response, error) {
	u := baseURL + "/upay/v2/precreate"
	req.TerminalSN = s.client.config.TerminalSN
	req.Subject = s.client.config.subject
	req.Operator = s.client.config.operator

	if s.client.config.NotifyURL != "" {
		req.NotifyURL = s.client.config.NotifyURL
	}

	signed, err := sign(req, s.client.config.TerminalSN, s.client.config.TerminalKey)
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts, WithAuthentication(signed))

	result := new(ApiResponse)
	resp, err := s.client.Request(ctx, http.MethodPost, u, req, result, opts...)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

type UPayRefundRequest struct {
	TerminalSN      string       `json:"terminal_sn"`                 // 收钱吧终端ID 	不超过32位的纯数字
	SN              string       `json:"sn,omitempty"`                // 收钱吧系统订单号
	ClientSN        string       `json:"client_sn,omitempty"`         // 商户自己的订单号
	RefundRequestNo string       `json:"refund_request_no,omitempty"` // 退款序列号
	Operator        string       `json:"operator,omitempty"`          // 门店操作员
	RefundAmount    string       `json:"refund_amount"`               // 退款金额
	Extended        interface{}  `json:"extended,omitempty"`          // 扩展参数集合
	GoodDetails     []GoodDetail `json:"goods_details,omitempty"`     // 商品详情
	Reflect         string       `json:"reflect,omitempty"`           // 业务反射参数
}

// Refund 收钱吧退款接口
// https://doc.shouqianba.com/zh-cn/api/interface/refund.html
func (s *UPayService) Refund(ctx context.Context, req *UPayRefundRequest, opts ...RequestOption) (*ApiResponse, *http.Response, error) {
	u := baseURL + "/upay/v2/refund"
	req.TerminalSN = s.client.config.TerminalSN
	req.Operator = s.client.config.operator

	signed, err := sign(req, s.client.config.TerminalSN, s.client.config.TerminalKey)
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts, WithAuthentication(signed))

	result := new(ApiResponse)
	resp, err := s.client.Request(ctx, http.MethodPost, u, req, result, opts...)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

type UPayQueryRequest struct {
	TerminalSN      string `json:"terminal_sn"`                 // 收钱吧终端ID 	不超过32位的纯数字
	SN              string `json:"sn,omitempty"`                // 收钱吧系统订单号
	ClientSN        string `json:"client_sn,omitempty"`         // 商户自己的订单号
	RefundRequestNo string `json:"refund_request_no,omitempty"` // 退款序列号
}

// Query 收钱吧查询接口
// https://doc.shouqianba.com/zh-cn/api/interface/query.html
func (s *UPayService) Query(ctx context.Context, req *UPayQueryRequest, opts ...RequestOption) (*ApiResponse, *http.Response, error) {
	u := baseURL + "/upay/v2/query"
	req.TerminalSN = s.client.config.TerminalSN

	signed, err := sign(req, s.client.config.TerminalSN, s.client.config.TerminalKey)
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts, WithAuthentication(signed))

	result := new(ApiResponse)
	resp, err := s.client.Request(ctx, http.MethodPost, u, req, result, opts...)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

type UPayCancelRequest struct {
	TerminalSN string `json:"terminal_sn"` // 收钱吧终端ID 	不超过32位的纯数字
	SN         string `json:"sn"`          // 收钱吧系统订单号
	ClientSN   string `json:"client_sn"`   // 商户自己的订单号
}

// Cancel 收钱吧撤销接口
// https://doc.shouqianba.com/zh-cn/api/interface/revoke&cancel.html
func (s *UPayService) Cancel(ctx context.Context, req *UPayCancelRequest, opts ...RequestOption) (*ApiResponse, *http.Response, error) {
	u := baseURL + "/upay/v2/cancel"
	req.TerminalSN = s.client.config.TerminalSN

	signed, err := sign(req, s.client.config.TerminalSN, s.client.config.TerminalKey)
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts, WithAuthentication(signed))

	result := new(ApiResponse)
	resp, err := s.client.Request(ctx, http.MethodPost, u, req, result, opts...)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
