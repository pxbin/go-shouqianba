# 收钱吧

### 使用
```go
import "github.com/pxbin/go-shouqianba"

config := &shouqianba.Config{}
client := shouqianba.NewClient(config)
```

### 终端激活
```go
terminal, _, err := client.Terminal.Activate(ctx)
```

### 终端签到
```go
checkin, _, err := client.Terminal.Checkin(ctx)
```

### 终端付款
```go
request := &shouqianba.PreCreateRequest{
    TerminalSN: "10298371039",
    ClientSN: "2006101016201512080095793262",
    TotalAmount: "10000", // 金额以 分 为单位。
    Subject: "测试商品",
}

result, _, err := client.UPay.Payment(ctx, request)
```

### 终端退款
```go
request := &shouqianba.UPayRefundRequest{
    TerminalSN: "10298371039",
    SN: "7894259244067218",
    ClientSN: "2006101016201512080095793262",
    RefundAmount: "10000", // 金额以 分 为单位。
    RefundRequestNo: "2006101016201512080095793262",
}

result, _, err := client.UPay.Refund(ctx, request)
```

### 终端查询
```go
request := &shouqianba.UPayQueryRequest{
    TerminalSN: "10298371039",
    SN: "7894259244067218",
    ClientSN: "2006101016201512080095793262",
}

result, _, err := client.UPay.Query(ctx, request)
```

### 终端撤销
```go
request := &shouqianba.CancelRequest{
    TerminalSN: "10298371039",
    SN: "7894259244067218",
    ClientSN: "2006101016201512080095793262",
}

result, _, err := client.UPay.Cancel(ctx, request)
```

### 终端预下单支付
```go
request := &shouqianba.PreCreateRequest{
    TerminalSN: "10298371039",
    TotalAmount: "10000", // 金额以 分 为单位。
    ClientSN: "2006101016201512080095793262",
    Subject: "测试商品",
}

result, _, err := client.UPay.PreCreate(ctx, request)
```

### 跳转支付
```go
request := &shouqianba.GatewayRequest{
    TerminalSN: "10298371039",
    TotalAmount: "10000", // 金额以 分 为单位。
    ClientSN: "2006101016201512080095793262",
    Subject: "测试商品",
}

result, _, err := client.UPay.JumpPay(ctx, request)
```