package shouqianba

import (
	"context"
	"encoding/json"
	"testing"
)

var config = &Config{}

func Test_TerminalActivate(t *testing.T) {
	c := NewClient(config)
	ctx := context.Background()

	res, _, err := c.Terminal.Activate(ctx)
	if err != nil {
		t.Errorf("TerminalActivate error: %v", err)
	}
	t.Log(res)
}

func Test_TerminalCheckIn(t *testing.T) {
	c := NewClient(config)
	ctx := context.Background()

	res, _, err := c.Terminal.CheckIn(ctx)
	if err != nil {
		t.Errorf("TerminalActivate error: %v", err)
	}
	t.Log(res)
}

func Test_UPayQuery(t *testing.T) {
	c := NewClient(config)
	ctx := context.Background()

	req := &UPayQueryRequest{
		TerminalSN: config.TerminalSN,
		SN:         "7895229858211347",
		ClientSN:   "e10adc3949ba59abbe56e057f20f883e",
	}

	res, _, err := c.UPay.Query(ctx, req)

	if err != nil {
		t.Errorf("UPayQuery error: %v", err)
	}
	t.Log(res)
}

func Test_UPayPrecreate(t *testing.T) {
	c := NewClient(config)
	ctx := context.Background()

	req := &PreCreateRequest{
		TerminalSN:  config.TerminalSN,
		ClientSN:    "e10adc3949ba59abbe56e057f20f883e",
		TotalAmount: 100,
		Payway:      "3",
		Subject:     "测试",
		Operator:    "Obama",
		Reflect:     "test",
	}

	res, _, err := c.UPay.Precreate(ctx, req)

	if err != nil {
		t.Errorf("Precreate error: %v", err)
	}
	pprint(t, res)

}

func pprint(t *testing.T, v interface{}) {
	bs, _ := json.Marshal(v)
	t.Logf(string(bs))
}
