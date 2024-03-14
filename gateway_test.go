package shouqianba

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-querystring/query"
)

func testValue(t *testing.T, input interface{}, want string) {
	v, err := query.Values(input)
	if err != nil {
		t.Errorf("Values(%q) returned error: %v", input, err)
	}
	if diff := cmp.Diff(want, v.Encode()); diff != "" {
		t.Errorf("Values(%#v) mismatch:\n%s", input, diff)
	}
}

func Test_StructToURLValues(t *testing.T) {
	tests := []struct {
		input interface{}
		want  string
	}{
		{
			input: struct {
				Name string `url:"name"`
				Age  int    `url:"age"`
			}{
				Name: "test",
				Age:  18,
			},
			want: "age=18&name=test",
		},
		{
			input: struct {
				Query   string `url:"q"`
				ShowAll bool   `url:"all"`
				Page    int    `url:"page"`
			}{"foo", true, 2},
			want: "all=true&page=2&q=foo",
		},
		{
			input: struct {
				TerminalSN  string `url:"terminal_sn"`
				ClientSN    string `url:"client_sn"`
				TotalAmount string `url:"total_amount"`
				Subject     string `url:"subject"`
			}{
				TerminalSN:  "test",
				ClientSN:    "test",
				TotalAmount: "10000",
				Subject:     "test",
			},
			want: "client_sn=test&subject=test&terminal_sn=test&total_amount=10000",
		},
	}

	for _, tt := range tests {
		testValue(t, tt.input, tt.want)
	}
}
