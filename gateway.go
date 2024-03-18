package shouqianba

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// GatewayService handles communication with the wap gateway related
// methods of the Shouqianba API.
//
// API docs: https://doc.shouqianba.com/zh-cn/api/wap2.html
type GatewayService service

type GatewayRequest struct {
	// AppID 收钱吧分配的APPID
	TerminalSN   string `mapstructure:"terminal_sn" json:"terminal_sn"`                         // 收钱吧终端ID
	ClientSN     string `mapstructure:"client_sn" json:"client_sn"`                             // 商户系统订单号
	TotalAmount  string `mapstructure:"total_amount" json:"total_amount"`                       // 订单总金额
	Subject      string `mapstructure:"subject" json:"subject"`                                 // 交易概述
	Payway       string `mapstructure:"payway,omitempty" json:"payway,omitempty"`               // 支付方式
	Operator     string `mapstructure:"operator" json:"operator,omitempty"`                     // 门店操作员
	Description  string `mapstructure:"description,omitempty" json:"description,omitempty"`     // 商品详情
	Longitude    string `mapstructure:"longitude,omitempty" json:"longitude,omitempty"`         // 经度
	Latitude     string `mapstructure:"latitude,omitempty" json:"latitude,omitempty"`           // 纬度
	Extended     string `mapstructure:"extended,omitempty" json:"extended,omitempty"`           // 扩展参数集合
	GoodsDetails string `mapstructure:"goods_details,omitempty" json:"goods_details,omitempty"` // 商品详情
	Reflect      string `mapstructure:"reflect,omitempty" json:"reflect,omitempty"`             // 反射参数
	NotifyURL    string `mapstructure:"notify_url,omitempty" json:"notify_url,omitempty"`       // 服务器异步回调 url
	ReturnURL    string `mapstructure:"return_url" json:"return_url"`                           // 页面跳转同步通知页面路径
}

// GetWapURL 收银吧网关支付 WAP支付收银台
// https://doc.shouqianba.com/zh-cn/api/wap2.html
func (s *GatewayService) GetWapURL(ctx context.Context, req *GatewayRequest) (string, error) {
	req.TerminalSN = s.client.config.TerminalSN

	if s.client.config.NotifyURL != "" {
		req.NotifyURL = s.client.config.NotifyURL
	}
	params := map[string]string{}
	mapstructure.Decode(req, &params)

	q := encode(params)

	signstr := fmt.Sprintf("%s&key=%s", q, s.client.config.TerminalKey)
	sum := md5.Sum([]byte(signstr))
	signed := strings.ToUpper(hex.EncodeToString(sum[:]))
	return fmt.Sprintf("%s?%s&sign=%s", apiGatewayBaseURL, q, signed), nil
}

func encode(v map[string]string) string {
	if len(v) == 0 {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := v[k]
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(v)
	}
	return buf.String()
}
