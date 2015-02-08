package lib

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Servers_GetServers_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	servers, err := client.GetServers()
	assert.Nil(t, servers)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_Servers_GetServers_NoServers(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	servers, err := client.GetServers()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, servers)
}

func Test_Servers_GetServers_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{
"one":{"SSHKEYID":"1","name":"alpha","ssh_key":"aaaa","date_created":null},
"two":{"SSHKEYID":"2","name":"beta","ssh_key":"bbbb","date_created":"2014-12-31 13:34:56"},
"three":{"SSHKEYID":"3","name":"charlie","ssh_key":"cccc"}}`)
	defer server.Close()

	servers, err := client.GetServers()
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, servers) {
		assert.Equal(t, 2, len(servers))
		// servers could be in random order
		for _, server := range servers {
			switch server.ID {
			case "1":
				assert.Equal(t, "alpha", key.Name)
				assert.Equal(t, "", key.Created)
			case "2":
				assert.Equal(t, "beta", key.Name)
				assert.Equal(t, "2014-12-31 13:34:56", key.Created)
			default:
				t.Error("Unknown SUBID")
			}
		}
	}
}

// TODO: add tests for GetServer, CreateServer, DeleteServer
