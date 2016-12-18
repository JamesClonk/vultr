package lib

import "net/url"
// import "fmt"

// Ips on Vultr account

// "1313044": {
//     "SUBID": 1313044,
//     "DCID": 1,
//     "ip_type": "v4",
//     "subnet": "10.234.22.53",
//     "subnet_size": 32,
//     "label": "my first reserved ip",
//     "attached_SUBID": 123456
// },

type SubId struct {
	SUBID int `json:"SUBID"`
}

type Ip struct {
	SUBID          int    `json:"SUBID"`
	DCID           int    `json:"DCID"`
	Ip_type        string `json:"ip_type"`
	Subnet         string `json:"subnet"`
	Subnet_size    int    `json:"subnet_size"`
	Label          string `json:"label"`
	Attached_SUBID bool   `json:"attached_SUBID"`
}

func (c *Client) ListReservedIp() (ips []Ip, err error) {
	var ipMap map[string]Ip
	err = c.get(`reservedip/list`, &ipMap)
	if err != nil {
		return nil, err
	}
	for _, ip := range ipMap {
		ips = append(ips, ip)
	}
	return ips, nil
}

func (c *Client) CreateReservedIp(dcid string, ip_type string) (subid SubId, err error) {
	values := url.Values{
		"DCID":    {dcid},
		"ip_type": {ip_type},
	}
	err = c.post(`reservedip/create`, values, &subid)
	if err != nil {
		return subid, err
	}
	return subid, nil
}

func (c *Client) DestroyReservedIp(subid string) (err error) {
	values := url.Values{
		"SUBID": {subid},
	}
	return c.post(`reservedip/destroy`, values, nil)
}

func (c *Client) AttachReservedIp(ip_address string, attach_subid string) (err error) {
	values := url.Values{
		"ip_address":   {ip_address},
		"attach_SUBID": {attach_subid},
	}
	return c.post(`reservedip/attach`, values, nil)
}

func (c *Client) ConvertReservedIp(subid string, ip_address string) (subId SubId, err error) {
	values := url.Values{
		"SUBID":      {subid},
		"ip_address": {ip_address},
	}
  // fmt.Printf("%s:%s\n", subid, ip_address)
	err = c.post(`reservedip/convert`, values, &subId)
	return subId, err
}

func (c *Client) DetachReservedIp(detach_subid string, ip_address string) (err error) {
	values := url.Values{
		"ip_address":   {ip_address},
		"detach_SUBID": {detach_subid},
	}
	return c.post(`reservedip/detach`, values, nil)
}
