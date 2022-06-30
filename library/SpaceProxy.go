package SpaceProxy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func getAddress(command string, apiKey string) string {
	addr := "https://panel.spaceproxy.net/api/"
	return addr + command + "/" + "?api_key=" + apiKey
}

func getData(apiKey string, command string, params string, flag bool) map[string]interface{} {
	resp, err := http.Get(getAddress(command, apiKey) + params)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		panic(err)
	}

	str := string(data)
	if flag == true {
		str = `{"data":` + str + "}"
	}

	mp := make(map[string]interface{})
	err = json.Unmarshal([]byte(str), &mp)
	if err != nil {
		panic(err)
	}

	return mp
}

func Proxies(
	apiKey string,
	status string,
	ipType string,
	ipVersion string,
	countries []string,
	dateAfter string,
	dateBefore string,
	dateAfterEnd string,
	dateBeforeEnd string,
	orders []string,
	limit string,
	offset string) map[string]interface{} {
	params := "&status=" + status + "&type=" + ipType + "&ip_version=" + ipVersion + "&limit=" + limit + "&offset=" + offset
	params += "&date_after=" + dateAfter + "&date_before=" + dateBefore
	params += "&date_end_after=" + dateAfterEnd + "&date_end_before=" + dateBeforeEnd

	for _, country := range countries {
		params += "&country=" + country
	}

	for _, order := range orders {
		params += "&orders=" + order
	}

	return getData(apiKey, "proxies", params, false)
}

func Orders(apiKey string, dateAfter string, dateBefore string) map[string]interface{} {
	params := "&date_after=" + dateAfter + "&date_before=" + dateBefore
	return getData(apiKey, "orders", params, true)
}

func Renew(apiKey string, proxies string, period string, coupon string) map[string]interface{} {
	data := url.Values{
		"proxies": {proxies},
		"period":  {period},
		"coupon":  {coupon},
	}

	resp, err := http.PostForm(getAddress("renew", apiKey), data)
	if err != nil {
		panic(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func NewOrder(apiKey string,
	ipType string,
	ipVersion string,
	country string,
	quantity string,
	period string,
	coupon string,
	ipList []string) map[string]interface{} {
	data := url.Values{
		"type":       {ipType},
		"ip_version": {ipVersion},
		"country":    {country},
		"quantity":   {quantity},
		"period":     {period},
		"coupon":     {coupon},
		"ip_list":    ipList,
	}

	resp, err := http.PostForm(getAddress("new-order", apiKey), data)
	if err != nil {
		panic(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func Countries(apiKey string) map[string]interface{} {
	return getData(apiKey, "countries", "", true)
}

func Ips(apiKey string, ipType string, ipVersion string, ipCountry string) map[string]interface{} {
	params := "&type=" + ipType + "&ip_version=" + ipVersion + "&country=" + ipCountry
	return getData(apiKey, "ips", params, true)
}

func IpsCount(apiKey string, ipType string, ipVersion string, ipCountry string) map[string]interface{} {
	params := "&type=" + ipType + "&ip_version=" + ipVersion + "&country=" + ipCountry
	return getData(apiKey, "ips-count", params, false)
}

func NewOrderAmount(apiKey string,
	ipType string,
	ipVersion string,
	country string,
	quantity string,
	period string,
	coupon string,
	ipList []string) map[string]interface{} {
	data := url.Values{
		"type":       {ipType},
		"ip_version": {ipVersion},
		"country":    {country},
		"quantity":   {quantity},
		"period":     {period},
		"coupon":     {coupon},
		"ip_list":    ipList,
	}

	resp, err := http.PostForm(getAddress("new-order-amount", apiKey), data)
	if err != nil {
		panic(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func Balance(apiKey string) map[string]interface{} {
	return getData(apiKey, "balance", "", false)
}
