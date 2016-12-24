package lib

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// IP on Vultr
type IP struct {
	ID         string `json:"SUBID,string"`
	RegionID   int    `json:"DCID,string"`
	IPType     string `json:"ip_type"`
	Subnet     string `json:"subnet"`
	SubnetSize int    `json:"subnet_size"`
	Label      string `json:"label"`
	AttachedTo string `json:"attached_SUBID,string"`
}

// Implements json.Unmarshaller on IP.
// This is needed because the Vultr API is inconsistent in it's JSON responses.
// Some fields can change type, from JSON number to JSON string and vice-versa.
func (i *IP) UnmarshalJSON(data []byte) (err error) {
	if i == nil {
		*i = IP{}
	}

	var fields map[string]interface{}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	value := fmt.Sprintf("%v", fields["SUBID"])
	if len(value) == 0 || value == "<nil>" || value == "0" {
		i.ID = ""
	} else {
		id, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		i.ID = strconv.FormatFloat(id, 'f', -1, 64)
	}

	value = fmt.Sprintf("%v", fields["DCID"])
	if len(value) == 0 || value == "<nil>" {
		value = "0"
	}
	region, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	i.RegionID = int(region)

	value = fmt.Sprintf("%v", fields["attached_SUBID"])
	if len(value) == 0 || value == "<nil>" || value == "0" || value == "false" {
		i.AttachedTo = ""
	} else {
		attached, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		i.AttachedTo = strconv.FormatFloat(attached, 'f', -1, 64)
	}

	value = fmt.Sprintf("%v", fields["subnet_size"])
	if len(value) == 0 || value == "<nil>" {
		value = "0"
	}
	size, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	i.SubnetSize = int(size)

	i.IPType = fmt.Sprintf("%v", fields["ip_type"])
	i.Subnet = fmt.Sprintf("%v", fields["subnet"])
	i.Label = fmt.Sprintf("%v", fields["label"])

	return
}

func (c *Client) ListReservedIP() ([]IP, error) {
	var ipMap map[string]IP

	err := c.get(`reservedip/list`, &ipMap)
	if err != nil {
		return nil, err
	}

	ips := make([]IP, 0)
	for _, ip := range ipMap {
		ips = append(ips, ip)
	}
	return ips, nil
}

func (c *Client) GetReservedIP(id string) (IP, error) {
	var ipMap map[string]IP

	err := c.get(`reservedip/list`, &ipMap)
	if err != nil {
		return IP{}, err
	}
	if ip, ok := ipMap[id]; ok {
		return ip, nil
	}
	return IP{}, fmt.Errorf("IP with id %v not found.", id)
}

func (c *Client) CreateReservedIP(regionID int, ipType string, label string) (string, error) {
	values := url.Values{
		"DCID":    {fmt.Sprintf("%v", regionID)},
		"ip_type": {ipType},
	}
	if len(label) > 0 {
		values.Add("label", label)
	}

	result := IP{}
	err := c.post(`reservedip/create`, values, &result)
	if err != nil {
		return "", err
	}
	return result.ID, nil
}

func (c *Client) DestroyReservedIP(id string) error {
	values := url.Values{
		"SUBID": {id},
	}
	return c.post(`reservedip/destroy`, values, nil)
}

func (c *Client) AttachReservedIP(ip string, serverId string) error {
	values := url.Values{
		"ip_address":   {ip},
		"attach_SUBID": {serverId},
	}
	return c.post(`reservedip/attach`, values, nil)
}

func (c *Client) ConvertReservedIP(serverId string, ip string) (string, error) {
	values := url.Values{
		"SUBID":      {serverId},
		"ip_address": {ip},
	}

	result := IP{}
	err := c.post(`reservedip/convert`, values, &result)
	if err != nil {
		return "", err
	}
	return result.ID, err
}

func (c *Client) DetachReservedIP(serverId string, ip string) error {
	values := url.Values{
		"ip_address":   {ip},
		"detach_SUBID": {serverId},
	}
	return c.post(`reservedip/detach`, values, nil)
}
