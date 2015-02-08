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
"9753721":{"SUBID":"9753721","os":"Ubuntu 14.04 x64","ram":"768 MB","disk":"Virtual 15 GB","main_ip":"123.456.789.0",
	"vcpu_count":"2","location":"Frankfurt","DCID":"9","default_password":"oops!","date_created":"2017-07-07 07:07:07",
	"pending_charges":0.04,"status":"active","cost_per_month":"5.00","current_bandwidth_gb":7,"allowed_bandwidth_gb":"1000",
	"netmask_v4":"255.255.255.0","gateway_v4":"123.456.789.1","power_status":"running","VPSPLANID":"29","v6_network":"::",
	"v6_main_ip":"","v6_network_size":"0","label":"test alpha","internal_ip":"",
	"kvm_url":"https:\/\/my.vultr.com\/subs\/vps\/novnc\/api.php?data=123","auto_backups":"no"},
"789032":{"SUBID":"789032","os":"CentOs 6.5 i368","ram":"1024 MB","disk":"Virtual 20 GB","main_ip":"192.168.1.2",
	"vcpu_count":"1","location":"Amsterdam","DCID":"21","default_password":"more oops!","date_created":"2011-01-01 01:01:01",
	"pending_charges":0.01,"status":"stopped","cost_per_month":"7.25","current_bandwidth_gb":0,"allowed_bandwidth_gb":"25",
	"netmask_v4":"255.255.254.0","gateway_v4":"192.168.1.1","power_status":"down","VPSPLANID":"31","v6_network":"::",
	"v6_main_ip":"?","v6_network_size":"2","label":"test 002","internal_ip":"10.10.10.10",
	"kvm_url":"https:\/\/my.vultr.com\/subs\/vps\/novnc\/api.php?data=456","auto_backups":"yes"}}`)
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
			case "9753721":
				assert.Equal(t, "test alpha", server.Name)
				assert.Equal(t, "Ubuntu 14.04 x64", server.OS)
				assert.Equal(t, "768 MB", server.RAM)
				assert.Equal(t, "Virtual 15 GB", server.Disk)
				assert.Equal(t, "123.456.789.0", server.MainIP)
				assert.Equal(t, 2, server.VCpus)
				assert.Equal(t, "Frankfurt", server.Location)
				assert.Equal(t, 9, server.RegionID)
				assert.Equal(t, "oops!", server.DefaultPassword)
				assert.Equal(t, "2017-07-07 07:07:07", server.Created)
				assert.Equal(t, "255.255.255.0", server.NetmaskV4)
				assert.Equal(t, "123.456.789.1", server.GatewayV4)
				assert.Equal(t, 7, server.CurrentBandwidth)
				assert.Equal(t, 1000, server.AllowedBandwidth)
			case "789032":
				assert.Equal(t, "test 002", server.Name)
				assert.Equal(t, 0.01, server.PendingCharges)
				assert.Equal(t, "7.25", server.Cost)
				assert.Equal(t, "stopped", server.Status)
				assert.Equal(t, "down", server.PowerStatus)
				assert.Equal(t, 31, server.PlanID)
				assert.Equal(t, "::", server.NetworkV6)
				assert.Equal(t, "?", server.MainIPV6)
				assert.Equal(t, "2", server.NetworkSizeV6)
				assert.Equal(t, "10.10.10.10", server.InternalIP)
				assert.Equal(t, `https://my.vultr.com/subs/vps/novnc/api.php?data=456`, server.KVMUrl)
				assert.Equal(t, "yes", server.AutoBackups)
			default:
				t.Error("Unknown SUBID")
			}
		}
	}
}

func Test_Servers_GetServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	_, err := client.GetServer("789032")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_Servers_GetServer_NoServer(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	s, err := client.GetServer("789032")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, Server{}, s)
}

func Test_Servers_GetServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `
{"SUBID":"9753721","os":"Ubuntu 14.04 x64","ram":"768 MB","disk":"Virtual 15 GB","main_ip":"123.456.789.0",
	"vcpu_count":"2","location":"Frankfurt","DCID":"9","default_password":"oops!","date_created":"2017-07-07 07:07:07",
	"pending_charges":0.04,"status":"active","cost_per_month":"5.00","current_bandwidth_gb":7,"allowed_bandwidth_gb":"1000",
	"netmask_v4":"255.255.255.0","gateway_v4":"123.456.789.1","power_status":"running","VPSPLANID":"29","v6_network":"::",
	"v6_main_ip":"","v6_network_size":"0","label":"test alpha","internal_ip":"",
	"kvm_url":"https:\/\/my.vultr.com\/subs\/vps\/novnc\/api.php?data=123","auto_backups":"no"}`)
	defer server.Close()

	s, err := client.GetServer("789032")
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, s) {
		assert.Equal(t, "test alpha", s.Name)
		assert.Equal(t, "Ubuntu 14.04 x64", s.OS)
		assert.Equal(t, "768 MB", s.RAM)
		assert.Equal(t, "Virtual 15 GB", s.Disk)
		assert.Equal(t, "123.456.789.0", s.MainIP)
		assert.Equal(t, 2, s.VCpus)
		assert.Equal(t, "Frankfurt", s.Location)
		assert.Equal(t, 9, s.RegionID)
		assert.Equal(t, "oops!", s.DefaultPassword)
		assert.Equal(t, "2017-07-07 07:07:07", s.Created)
		assert.Equal(t, "255.255.255.0", s.NetmaskV4)
		assert.Equal(t, "123.456.789.1", s.GatewayV4)
		assert.Equal(t, 7, s.CurrentBandwidth)
		assert.Equal(t, 1000, s.AllowedBandwidth)
		assert.Equal(t, 0.04, s.PendingCharges)
		assert.Equal(t, "5.00", s.Cost)
		assert.Equal(t, "active", s.Status)
		assert.Equal(t, "running", s.PowerStatus)
		assert.Equal(t, 29, s.PlanID)
		assert.Equal(t, "::", s.NetworkV6)
		assert.Equal(t, "", s.MainIPV6)
		assert.Equal(t, "0", s.NetworkSizeV6)
		assert.Equal(t, "", s.InternalIP)
		assert.Equal(t, `https://my.vultr.com/subs/vps/novnc/api.php?data=123`, s.KVMUrl)
		assert.Equal(t, "no", s.AutoBackups)
	}
}

func Test_Servers_CreateServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	s, err := client.CreateServer("test", 1, 2, 3, nil)
	assert.Equal(t, Server{}, s)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_Servers_CreateServer_NoServer(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	s, err := client.CreateServer("test", 1, 2, 3, nil)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "", s.ID)
}

func Test_Servers_CreateServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{"SUBID":"123456789"}`)
	defer server.Close()

	s, err := client.CreateServer("test", 1, 2, 3, nil)
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, s) {
		assert.Equal(t, "123456789", s.ID)
		assert.Equal(t, "test", s.Name)
		assert.Equal(t, 1, s.RegionID)
		assert.Equal(t, 2, s.PlanID)
	}

	options := &ServerOptions{
		IPXEChainURL:      "...",
		ISO:               1,
		Script:            2,
		Snapshot:          "alpha",
		SSHKey:            "key123",
		IPV6:              true,
		PrivateNetworking: true,
		AutoBackups:       true,
	}
	s2, err := client.CreateServer("test2", 4, 5, 6, options)
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, s2) {
		assert.Equal(t, "123456789", s2.ID)
		assert.Equal(t, "test2", s2.Name)
		assert.Equal(t, 4, s2.RegionID)
		assert.Equal(t, 5, s2.PlanID)
	}

}

func Test_Servers_DeleteServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.DeleteServer("123456789")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_Servers_DeleteServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.DeleteServer("123456789"))
}
