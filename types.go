package shouqianba

type (
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
		TerminalSN        string     `json:"terminal_sn,omitempty"`
		SN                string     `json:"sn,omitempty"`
		ClientSN          string     `json:"client_sn,omitempty"`
		Status            string     `json:"status,omitempty"`
		OrderStatus       string     `json:"order_status,omitempty"`
		PayerUID          string     `json:"payer_uid,omitempty"`
		PayerLogin        string     `json:"payer_login,omitempty"`
		TradeNo           string     `json:"trade_no,omitempty"`
		QrCode            string     `json:"qr_code,omitempty"`
		TotalAmount       int64      `json:"total_amount,string,omitempty"`
		NetAmount         int64      `json:"net_amount,string,omitempty"`
		Payway            string     `json:"payway,omitempty"`
		PaywayName        string     `json:"payway_name,omitempty"`
		SubPayway         string     `json:"sub_payway,omitempty"`
		FinishTime        *Timestamp `json:"finish_time,omitempty"`
		ChannelFinishTime *Timestamp `json:"channel_finish_time,omitempty"`
		StoreID           string     `json:"store_id,omitempty"`
		Subject           string     `json:"subject,omitempty"`
		Description       string     `json:"description,omitempty"`
		Reflect           string     `json:"reflect,omitempty"`
		Operator          string     `json:"operator,omitempty"`
		WapPayRequest     string     `json:"wap_pay_request,omitempty"`
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
