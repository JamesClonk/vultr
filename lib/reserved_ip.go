package lib

import (
	"net/url"
)

// DNS Domain
type ReservedIP struct {
	ID   string `json:"SUBID"`
	DCID string `json:"DCID"`

	IPType     string `json:"ip_type"`
	Subnet     string `json:"subnet"`
	SubnetSize string `json:"subnet_size"`
	Label      string `json:"label"`
	AttachedID string `json:"attached_SUBID"`
}

func (c *Client) AttachReserveIP(ip, serverId string) error {
	values := url.Values{
		"ip_address":   {ip},
		"attach_SUBID": {serverId},
	}

	if err := c.post(`reservedip/attach`, values, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) ConvertToReservedIP(ip, serverId string) (ReservedIP, error) {
	values := url.Values{
		"SUBID":      {serverId},
		"ip_address": {ip},
	}

	var reservedIP ReservedIP
	if err := c.post(`reservedip/convert`, values, &reservedIP); err != nil {
		return ReservedIP{}, err
	}
	return reservedIP, nil
}

func (c *Client) CreateReservedIP(dcId, ipType, label string) (ReservedIP, error) {
	values := url.Values{
		"DCID":    {dcId},
		"ip_type": {ipType},
	}
	if len(label) > 0 {
		values.Add("label", label)
	}

	var reservedIP ReservedIP
	if err := c.post(`reservedip/create`, values, &reservedIP); err != nil {
		return ReservedIP{}, err
	}
	return reservedIP, nil
}

func (c *Client) DestroyReservedIP(id string) error {
	values := url.Values{
		"SUBID": {id},
	}

	if err := c.post(`reservedip/destroy`, values, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) DetachReservedIP(id string) error {
	values := url.Values{
		"SUBID": {id},
	}

	if err := c.post(`reservedip/detach`, values, nil); err != nil {
		return err
	}
	return nil
}

func (c *Client) ListReservedIPs() (ips []ReservedIP, err error) {
	if err := c.get(`reservedip/list`, &ips); err != nil {
		return nil, err
	}
	return ips, nil
}

func (c *Client) GetReservedIP(id string) (ip ReservedIP, err error) {
	ips := map[string]ReservedIP{}
	if err := c.get(`reservedip/list`, &ips); err != nil {
		return ReservedIP{}, err
	}
	if ip, ok := ips[id]; ok {
		return ip, nil
	}
	return ReservedIP{}, nil
}
