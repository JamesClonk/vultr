package lib

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SshKeys_GetSSHKeys_Error(t *testing.T) {
	server := getTestServer(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	client := getTestClient(t, server.URL)

	keys, err := client.GetSSHKeys()
	assert.Nil(t, keys)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_SshKeys_GetSSHKeys_NoKeys(t *testing.T) {
	server := getTestServer(http.StatusOK, `[]`)
	defer server.Close()

	client := getTestClient(t, server.URL)

	keys, err := client.GetSSHKeys()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, keys)
}

func Test_SshKeys_GetSSHKeys_Keys(t *testing.T) {
	server := getTestServer(http.StatusOK, `{
"one":{"SSHKEYID":"1","name":"alpha","ssh_key":"aaaa","date_created":null},
"two":{"SSHKEYID":"2","name":"beta","ssh_key":"bbbb","date_created":"2014-12-31 13:34:56"},
"three":{"SSHKEYID":"3","name":"charlie","ssh_key":"cccc"}}`)
	defer server.Close()

	client := getTestClient(t, server.URL)

	keys, err := client.GetSSHKeys()
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, keys) {
		assert.Equal(t, 3, len(keys))
		// keys could be in random order
		for _, key := range keys {
			switch key.ID {
			case "1":
				assert.Equal(t, "alpha", key.Name)
				assert.Equal(t, "", key.Created)
			case "2":
				assert.Equal(t, "beta", key.Name)
				assert.Equal(t, "2014-12-31 13:34:56", key.Created)
			case "3":
				assert.Equal(t, "cccc", key.Key)
			default:
				t.Error("Unknown Key ID")
			}
		}
	}
}

func Test_SshKeys_CreateSSHKey_Error(t *testing.T) {
	server := getTestServer(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	client := getTestClient(t, server.URL)

	key, err := client.CreateSSHKey("delta", "ddddd")
	assert.Equal(t, SSHKey{}, key)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_SshKeys_CreateSSHKey_NoKey(t *testing.T) {
	server := getTestServer(http.StatusOK, `[]`)
	defer server.Close()

	client := getTestClient(t, server.URL)

	key, err := client.CreateSSHKey("delta", "ddddd")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "", key.ID)
}

func Test_SshKeys_CreateSSHKey_KeyId(t *testing.T) {
	server := getTestServer(http.StatusOK, `{"SSHKEYID":"a1b2c3d4"}`)
	defer server.Close()

	client := getTestClient(t, server.URL)

	key, err := client.CreateSSHKey("delta", "ddddd")
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, key) {
		assert.Equal(t, "a1b2c3d4", key.ID)
		assert.Equal(t, "delta", key.Name)
		assert.Equal(t, "ddddd", key.Key)
	}
}
