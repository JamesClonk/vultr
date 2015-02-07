package lib

import "net/url"

// SSHKey on Vultr account
type SSHKey struct {
	ID      string `json:"SSHKEYID"`
	Name    string `json:"name"`
	Key     string `json:"ssh_key"`
	Created string `json:"date_created"`
}

// SSHKeys on Vultr account
type SSHKeys []SSHKey

func (c *Client) GetSSHKeys() (keys SSHKeys, err error) {
	var vultrKeys map[string]SSHKey
	if err := c.get(`/sshkey/list`, &vultrKeys); err != nil {
		return nil, err
	}

	for _, key := range vultrKeys {
		keys = append(keys, key)
	}

	return keys, nil
}

func (c *Client) CreateSSHKey(name, key string) (SSHKey, error) {
	values := url.Values{
		"name":    {name},
		"ssh_key": {key},
	}

	var sshKey SSHKey
	if err := c.post(`/sshkey/create`, values, &sshKey); err != nil {
		return SSHKey{}, err
	}
	sshKey.Name = name
	sshKey.Key = key

	return sshKey, nil
}
