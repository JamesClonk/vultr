package lib

import (
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Networks_GetNetworks_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	nets, err := client.GetNetworks()
	assert.Nil(t, nets)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_Networks_GetNetworks_NoNets(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	nets, err := client.GetNetworks()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, nets)
}

func Test_Networks_GetNetworks_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{
    "net539626f0798d7": {
        "DCID": "1",
        "NETWORKID": "net539626f0798d7",
        "date_created": "2017-08-25 12:23:45",
        "description": "test1",
        "v4_subnet": "10.99.0.0",
        "v4_subnet_mask": 24
    },
    "net53962b0f2341f": {
        "DCID": "1",
        "NETWORKID": "net53962b0f2341f",
        "date_created": "2014-06-09 17:45:51",
        "description": "vultr",
        "v4_subnet": "0.0.0.0",
        "v4_subnet_mask": 0
    }
}`)
	defer server.Close()

	nets, err := client.GetNetworks()
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, nets) {
		assert.Equal(t, 2, len(nets))

		assert.Equal(t, "net539626f0798d7", nets[0].ID)
		assert.Equal(t, 1, nets[0].RegionID)
		assert.Equal(t, "test1", nets[0].Description)
		assert.Equal(t, "2017-08-25 12:23:45", nets[0].Created)
		assert.Equal(t, "10.99.0.0", nets[0].V4Subnet)
		assert.Equal(t, 24, nets[0].V4SubnetMask)

		assert.Equal(t, "net53962b0f2341f", nets[1].ID)
		assert.Equal(t, 1, nets[1].RegionID)
		assert.Equal(t, "vultr", nets[1].Description)
		assert.Equal(t, "2014-06-09 17:45:51", nets[1].Created)
		assert.Equal(t, "0.0.0.0", nets[1].V4Subnet)
		assert.Equal(t, 0, nets[1].V4SubnetMask)
	}
}

func Test_Networks_CreateNetwork_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	_, subnet, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		t.Error(err)
	}

	net, err := client.CreateNetwork(1, "test", subnet)
	assert.Equal(t, Network{}, net)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_Networks_CreateNetwork_NoNet(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	_, subnet, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		t.Error(err)
	}

	net, err := client.CreateNetwork(1, "test", subnet)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "", net.ID)
}

func Test_Networks_CreateNetwork_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{"NETWORKID": "net59a0526477dd3"}`)
	defer server.Close()

	_, subnet, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		t.Error(err)
	}
	mask, _ := subnet.Mask.Size()

	net, err := client.CreateNetwork(1, "test", subnet)
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, net) {
		assert.Equal(t, "net59a0526477dd3", net.ID)
		assert.Equal(t, "test", net.Description)
		assert.Equal(t, subnet.IP.String(), net.V4Subnet)
		assert.Equal(t, mask, net.V4SubnetMask)
	}
}

func Test_Networks_DeleteNetwork_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.DeleteNetwork("id-1")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_Networks_DeleteNetwork_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.DeleteNetwork("id-1"))
}
