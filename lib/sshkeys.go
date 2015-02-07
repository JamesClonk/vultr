package lib

import "net/url"

type SshKey struct {
	Id      string `json:"SSHKEYID"`
	Name    string `json:"name"`
	Key     string `json:"ssh_key"`
	Created string `json:"date_created"`
}

type SshKeys []SshKey

func (c *Client) GetSshKeys() (keys SshKeys, err error) {
	var vultrKeys map[string]SshKey
	if err := c.get(`/sshkey/list`, &vultrKeys); err != nil {
		return nil, err
	}

	for _, key := range vultrKeys {
		keys = append(keys, key)
	}

	return keys, nil
}

func (c *Client) CreateSshKey(name, key string) (SshKey, error) {
	values := url.Values{
		"name":    {name},
		"ssh_key": {key},
	}

	var sshKey SshKey
	if err := c.post(`/sshkey/create`, values, &sshKey); err != nil {
		return SshKey{}, err
	}
	sshKey.Name = name
	sshKey.Key = key

	return sshKey, nil
}
