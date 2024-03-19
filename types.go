package shouqianba

import (
	"encoding/json"
	"time"
)

type (
	// JSONTime overrides MarshalJson method to format in ISO8601
	JSONTime time.Time

	// ApiResponse is the response from Shouqianba base biz response API
	ApiResponse struct {
		ResultCode  string      `json:"result_code"`
		ErrCode     string      `json:"error_code,omitempty"`
		ErrMessage  string      `json:"error_message,omitempty"`
		BizResponse BizResponse `json:"biz_response,omitempty"`
	}

	// BizResponse 业务响应数据
	BizResponse struct {
		ResultCode string  `json:"result_code"`
		ErrCode    string  `json:"error_code,omitempty"`
		ErrMessage string  `json:"error_message,omitempty"`
		BizData    BizData `json:"data,omitempty"`
	}

	// BizData 业务响应数据报文
	// https://doc.shouqianba.com/zh-cn/api/annex/responseParams.html#biz_response.data内字段列表
	BizData struct {
		SN                string      `json:"sn,omitempty"`
		ClientSN          string      `json:"client_sn,omitempty"`
		Status            string      `json:"status,omitempty"`
		OrderStatus       string      `json:"order_status,omitempty"`
		PayerUID          string      `json:"payer_uid,omitempty"`
		PayLogin          string      `json:"pay_login,omitempty"`
		TradeNo           string      `json:"trade_no,omitempty"`
		QrCode            string      `json:"qr_code,omitempty"`
		TotalAmount       json.Number `json:"total_amount,omitempty"`
		NetAmount         json.Number `json:"net_amount,omitempty"`
		Payway            string      `json:"payway,omitempty"`
		SubPayway         string      `json:"sub_payway,omitempty"`
		FinishTime        string      `json:"finish_time,omitempty"`
		ChannelFinishTime string      `json:"channel_finish_time,omitempty"`
		TerminalSN        string      `json:"terminal_sn,omitempty"`
		StoreID           string      `json:"store_id,omitempty"`
		Subject           string      `json:"subject,omitempty"`
		Description       string      `json:"description,omitempty"`
		Reflect           string      `json:"reflect,omitempty"`
		Operator          string      `json:"operator,omitempty"`
		WapPayRequest     string      `json:"wap_pay_request,omitempty"`
	}

	// GoodDetail represents the goods detail for payment
	GoodDetail struct {
		GoodsId       string `json:"goods_id"`       // 商品ID
		GoodsName     string `json:"goods_name"`     // 商品名称
		Price         int    `json:"price"`          // 商品单价
		Quantity      int    `json:"quantity"`       // 商品数量
		PromotionType int    `json:"promotion_type"` // 优惠类型，0表示没有优惠，1表示支付机构优惠，为1会把相关信息送到支付机构
	}
)

// MarshalJSON for JSONTime
func (t *JSONTime) MarshalJSON() ([]byte, error) {

	time.Time(*t).UTC().Format(time.RFC3339)

	return []byte(time.Time(*t).Format(`"2006-01-02T15:04:05Z"`)), nil
}

// UnmarshalJSON for JSONTime
func (t *JSONTime) UnmarshalJSON(b []byte) error {
	tt, err := time.Parse(`"2006-01-02T15:04:05Z"`, string(b))
	if err != nil {
		return err
	}
	*t = JSONTime(tt)
	return nil
}
