package lib

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Plans_ListReservedIp_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	_, err := client.ListReservedIp()
	if err == nil {
		t.Error(err)
	}
}

func Test_Plans_ListReservedIp_Ok_Empty(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{}`)
	defer server.Close()
	list, err := client.ListReservedIp()
	if err == nil {
		t.Error(err)
	}
	assert.Equal(t, len(list), 0)
}

func Test_Plans_ListReservedIp_Ok(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK,
		`{
      "4":{"SUBID":4,"DCID":5,"ip_type":"v7","subnet":"subnet",
           "subnet_size":8,"label":"label","attached_SUBID":false},
      "9":{"SUBID":9,"DCID":5,"ip_type":"v7","subnet":"subnet",
           "subnet_size":8,"label":"label","attached_SUBID":true}
      }`)
	defer server.Close()
	list, err := client.ListReservedIp()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, list[0].SUBID, 4)
	assert.Equal(t, list[0].DCID, 5)
	assert.Equal(t, list[0].Ip_type, "v7")
	assert.Equal(t, list[0].Subnet, "subnet")
	assert.Equal(t, list[0].Subnet_size, 8)
	assert.Equal(t, list[0].Label, "label")
	assert.Equal(t, list[0].Attached_SUBID, false)

  assert.Equal(t, list[1].SUBID, 9)
  assert.Equal(t, list[1].DCID, 5)
  assert.Equal(t, list[1].Ip_type, "v7")
  assert.Equal(t, list[1].Subnet, "subnet")
  assert.Equal(t, list[1].Subnet_size, 8)
  assert.Equal(t, list[1].Label, "label")
  assert.Equal(t, list[1].Attached_SUBID, true)

}

func Test_Plans_CreateReservedIp_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	_, err := client.CreateReservedIp("dcid", "ip")
	if err == nil {
		t.Error(err)
	}
}

func Test_Plans_CreateReservedIp_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{"SUBID":4711}`)
	defer server.Close()
	subid, err := client.CreateReservedIp("dcid", "ip")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, subid.SUBID, 4711)
}

func Test_Plans_DestroyReservedIp_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	err := client.DestroyReservedIp("subid")
	if err == nil {
		t.Error(err)
	}
}

func Test_Plans_DestroyReservedIp_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, ``)
	defer server.Close()
	err := client.DestroyReservedIp("subid")
	if err != nil {
		t.Error(err)
	}
}

func Test_Plans_AttachReservedIp_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	err := client.AttachReservedIp("ip", "subid")
	if err == nil {
		t.Error(err)
	}
}

func Test_Plans_AttachReservedIp_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, ``)
	defer server.Close()
	err := client.AttachReservedIp("subid", "ip")
	if err != nil {
		t.Error(err)
	}
}

func Test_Plans_ConvertReservedIp_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	_, err := client.ConvertReservedIp("subid", "ip")
	if err == nil {
		t.Error(err)
	}
}

func Test_Plans_ConvertReservedIp_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{"SUBID":4711}`)
	defer server.Close()
	subid, err := client.ConvertReservedIp("subid", "ip")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, subid.SUBID, 4711)
}

func Test_Plans_DetachReservedIp_Fail(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, ``)
	defer server.Close()

	err := client.DetachReservedIp("subid", "ip")
	if err == nil {
		t.Error(err)
	}
}

func Test_Plans_DetachReservedIp_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, ``)
	defer server.Close()

	err := client.DetachReservedIp("subid", "ip")
	if err != nil {
		t.Error(err)
	}
}
