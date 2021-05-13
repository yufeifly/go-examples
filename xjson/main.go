package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	str := `"{\"field0\":\"ZHRUQnFlZEROc3l0dlJFdlpNVUtJZ3FreWVvWUN0Q2FFU3RncldSS2lFdG9aeWFsblJUblhKdGpHaUNqRVZEQ2FxZHZleFlCWFFYZ0VFYWZscHdQWkxwcFRBY2hGbW1lQmdkbg==\",\"field1\":\"S2N5SXBPVk94SVJhd21VSHBZelp1ZHZ4YVJZbHhUUUlvaXB6WVRjU3BjckJEV1VnSlNFV0l2bkNtTUJwVU5wQmNKaGpleGFNZGh3RnlnT0ZTSktDcWRqbFNueHpmU0JpcnhPdQ==\",\"field2\":\"U0tQblBJaW1ZaHdmUGdnWVhHSHhua0RhaVhScE5sa25aR05CZXBSb21DcEtlbnhIampPTmNaWGtyanNWU3V1UVhXRUZUdmZUZ2pIYW13V2h1UWdFc0RackNOeG5Fa05IWmZMcQ==\",\"field3\":\"U3BYbkt0bWViUFpJcnRNTWFEZXlmSnh2dnNIaEpUclN4RGlCVXBLcnpCcUxHbGpraVFnd014cW9FUWxXbHRETm5LVERlR3JJT3lBalp3dGxkb3JiUVF1c0lLdXVqR1Noa1ZuWA==\",\"field4\":\"Q2VCTkhhckFYQlVTd0RKSE5CcEhqcWRyU0dBTmJDRlJjd0dQdmpidk5za0dGZW1Gc0NUeEt0dmdHWWJKeVdBZ1pYRllYSGlOR2hXTEx2eVlqYXRoeE5jVFp5bExUeEhQYnZGag==\",\"field5\":\"Y3BEQUpMVkhSaWd5SWlUeGRsZXNoakZ5c1RCR2ZXbkFKdGNKVmhIckNlQWxFQ3ltbndJUVhRSnlteEJHYnNNcW9Md0F3ZHpXQ3RLUWthdUdQbXdubnJhRlNTSHlaUmFwRXB5Tw==\",\"field6\":\"YXNjbkRFWm1WcFJDdmZlZUdlaGhjb05naWlKQlhVYktEYlVmZ1FteFlacElnR2prcFVtc0pGYmdTbUdVRWN3YkRvcnhkcHlQUGl6aU1pZVJUVmt5WndsVnprclJVYkZ6dVVxZw==\",\"field7\":\"Qm5HSVlnbHJTdlJyTUpnSXB3R0p0T0xpWGtzUUxFd2JjWWlYQnNpdkluaEpHSkJPTkZXRFZGZ2NaREttb25EWmFYTWREcXprU01hSXpseEJQb3FGRnhwbklxT1d3cnV4THJVSA==\",\"field8\":\"ekhlZ0F5cVZ0bXVZQ0xxYlZDam50dVhod2dveUlReG1FdWt6ZHlZd0lqRURPdUhYc3ptYVRPVEZaRUdGZ3JFQklVTnZLQUh3bmhYeUpJYWdZRHFHQnJueXdwbGNMRUdTdkd2bg==\",\"field9\":\"TG1lYnJvVW9YR0dYdE5kcXlmcGRnYmptVk5lZ1FwamNPbGpneFVJUERsUlVmYkFCS2R3S2RGSFVRSXVZVXJiV0lnWHNUQWtnc3Z0bXBwcmFRVmh0cEhVTVdpeFpiTmh1dnNnZQ==\"}"`
	var buf bytes.Buffer
	for _, ch := range str {
		if ch == '\\' {
			continue
		}
		buf.WriteRune(ch)
	}
	ans := buf.String()
	ans = ans[1 : len(ans)-1]
	//ans := strings.Trim(str, "\"")
	fmt.Printf("ans: %v\n", ans)
	var data map[string][]byte
	//fmt.Printf("str: %v\n", str)
	err := json.Unmarshal([]byte(ans), &data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("data: %v\n", data)
}
