package lib

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BareMetalServers_GetBareMetalServers_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	bareMetalServers, err := client.GetBareMetalServers()
	assert.Nil(t, bareMetalServers)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_GetBareMetalServers_NoBareMetalServers(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	bareMetalServers, err := client.GetBareMetalServers()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, bareMetalServers)
}

func Test_BareMetalServers_GetBareMetalServers_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{
"789032":{"SUBID":"789032","os":"CentOs 6.5 i368","ram":"1024 MB","disk":"Virtual 20 GB","main_ip":"192.168.1.2",
	"cpu_count":1,"location":"Amsterdam","DCID":"21","default_password":"more oops!","date_created":"2011-01-01 01:01:01",
	"status":"stopped","netmask_v4":"255.255.254.0","gateway_v4":"192.168.1.1","METALPLANID":31,
	"v6_networks": [{"v6_network": "2002:DB9:1000::", "v6_main_ip": "2000:DB8:1000::0000", "v6_network_size": 32 }],
	"label":"test beta","OSID": "127","APPID": "2"},
"9753721":{"SUBID":"9753721","os":"Ubuntu 14.04 x64","ram":"768 MB","disk":"Virtual 15 GB","main_ip":"123.456.789.0",
	"cpu_count":2,"location":"Frankfurt","DCID":"9","default_password":"oops!","date_created":"2017-07-07 07:07:07",
	"status":"active","netmask_v4":"255.255.255.0","gateway_v4":"123.456.789.1","METALPLANID":29,
	"label":"test alpha","APPID": "0"}}`)
	defer server.Close()

	bareMetalServers, err := client.GetBareMetalServers()
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, bareMetalServers) {
		assert.Equal(t, 2, len(bareMetalServers))

		assert.Equal(t, "9753721", bareMetalServers[0].ID)
		assert.Equal(t, "test alpha", bareMetalServers[0].Name)
		assert.Equal(t, "Ubuntu 14.04 x64", bareMetalServers[0].OS)
		assert.Equal(t, "768 MB", bareMetalServers[0].RAM)
		assert.Equal(t, "Virtual 15 GB", bareMetalServers[0].Disk)
		assert.Equal(t, "123.456.789.0", bareMetalServers[0].MainIP)
		assert.Equal(t, 2, bareMetalServers[0].CPUs)
		assert.Equal(t, "Frankfurt", bareMetalServers[0].Location)
		assert.Equal(t, 9, bareMetalServers[0].RegionID)
		assert.Equal(t, "oops!", bareMetalServers[0].DefaultPassword)
		assert.Equal(t, "2017-07-07 07:07:07", bareMetalServers[0].Created)
		assert.Equal(t, "255.255.255.0", bareMetalServers[0].NetmaskV4)
		assert.Equal(t, "123.456.789.1", bareMetalServers[0].GatewayV4)
		assert.Equal(t, 0, len(bareMetalServers[0].V6Networks))
		assert.Equal(t, "", bareMetalServers[0].OSID)
		assert.Equal(t, "", bareMetalServers[0].AppID)

		assert.Equal(t, "789032", bareMetalServers[1].ID)
		assert.Equal(t, "test beta", bareMetalServers[1].Name)
		assert.Equal(t, "stopped", bareMetalServers[1].Status)
		assert.Equal(t, 31, bareMetalServers[1].PlanID)
		assert.Equal(t, 1, len(bareMetalServers[1].V6Networks))
		assert.Equal(t, "2002:DB9:1000::", bareMetalServers[1].V6Networks[0].Network)
		assert.Equal(t, "2000:DB8:1000::0000", bareMetalServers[1].V6Networks[0].MainIP)
		assert.Equal(t, "32", bareMetalServers[1].V6Networks[0].NetworkSize)
		assert.Equal(t, "127", bareMetalServers[1].OSID)
		assert.Equal(t, "2", bareMetalServers[1].AppID)
	}
}

func Test_BareMetalServers_GetBareMetalServersByTag_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	bareMetalServers, err := client.GetBareMetalServersByTag("Unknown")
	assert.Nil(t, bareMetalServers)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_GetBareMetalServersByTag_NoBareMetalServers(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	bareMetalServers, err := client.GetBareMetalServersByTag("Nothing")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, bareMetalServers)
}

func Test_BareMetalServers_GetBareMetalServersByTag_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{
"789032":{"SUBID":"789032","os":"CentOs 6.5 i368","ram":"1024 MB","disk":"Virtual 20 GB","main_ip":"192.168.1.2",
	"cpu_count":"1","location":"Amsterdam","DCID":"21","default_password":"more oops!","date_created":"2011-01-01 01:01:01",
	"status":"stopped","netmask_v4":"255.255.254.0","gateway_v4":"192.168.1.1","METALPLANID":"31",
	"v6_networks": [{"v6_network": "2001:DB8:1000::", "v6_main_ip": "2001:DB8:1000::100", "v6_network_size": "64" }],
	"label":"test 002","tag":"Database"}}`)
	defer server.Close()

	bareMetalServers, err := client.GetBareMetalServersByTag("Database")
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, bareMetalServers) {
		assert.Equal(t, 1, len(bareMetalServers))
		assert.Equal(t, "test 002", bareMetalServers[0].Name)
		assert.Equal(t, "stopped", bareMetalServers[0].Status)
		assert.Equal(t, 31, bareMetalServers[0].PlanID)
		assert.Equal(t, 1, len(bareMetalServers[0].V6Networks))
		assert.Equal(t, "2001:DB8:1000::", bareMetalServers[0].V6Networks[0].Network)
		assert.Equal(t, "2001:DB8:1000::100", bareMetalServers[0].V6Networks[0].MainIP)
		assert.Equal(t, "64", bareMetalServers[0].V6Networks[0].NetworkSize)
	}
}

func Test_BareMetalServers_GetBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	_, err := client.GetBareMetalServer("789032")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_GetBareMetalServer_NoBareMetalServer(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	s, err := client.GetBareMetalServer("789032")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, BareMetalServer{}, s)
}

func Test_BareMetalServers_GetBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `
{"SUBID":"9753721","os":"Ubuntu 14.04 x64","ram":"768 MB","disk":"Virtual 15 GB","main_ip":"123.456.789.0",
	"cpu_count":"2","location":"Frankfurt","DCID":"9","default_password":"oops!","date_created":"2017-07-07 07:07:07",
	"status":"active","netmask_v4":"255.255.255.0","gateway_v4":"123.456.789.1","METALPLANID":"29","label":"test alpha",
	"v6_networks": [{"v6_network": "::", "v6_main_ip": "", "v6_network_size": "0" }]}`)
	defer server.Close()

	s, err := client.GetBareMetalServer("789032")
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, s) {
		assert.Equal(t, "test alpha", s.Name)
		assert.Equal(t, "Ubuntu 14.04 x64", s.OS)
		assert.Equal(t, "768 MB", s.RAM)
		assert.Equal(t, "Virtual 15 GB", s.Disk)
		assert.Equal(t, "123.456.789.0", s.MainIP)
		assert.Equal(t, 2, s.CPUs)
		assert.Equal(t, "Frankfurt", s.Location)
		assert.Equal(t, 9, s.RegionID)
		assert.Equal(t, "oops!", s.DefaultPassword)
		assert.Equal(t, "2017-07-07 07:07:07", s.Created)
		assert.Equal(t, "255.255.255.0", s.NetmaskV4)
		assert.Equal(t, "123.456.789.1", s.GatewayV4)
		assert.Equal(t, "active", s.Status)
		assert.Equal(t, 29, s.PlanID)
		assert.Equal(t, 1, len(s.V6Networks))
		assert.Equal(t, "::", s.V6Networks[0].Network)
		assert.Equal(t, "", s.V6Networks[0].MainIP)
		assert.Equal(t, "0", s.V6Networks[0].NetworkSize)
	}
}

func Test_BareMetalServers_CreateBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	s, err := client.CreateBareMetalServer("test", 1, 2, 3, nil)
	assert.Equal(t, BareMetalServer{}, s)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_CreateBareMetalServer_NoBareMetalServer(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	s, err := client.CreateBareMetalServer("test", 1, 2, 3, nil)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "", s.ID)
}

func Test_BareMetalServers_CreateBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{"SUBID":"123456789",
		"cpu_count":"1",
		"DCID":17,
		"VPSPLANID":"29"}`)
	defer server.Close()

	s, err := client.CreateBareMetalServer("test", 1, 2, 3, nil)
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, s) {
		assert.Equal(t, "123456789", s.ID)
		assert.Equal(t, "test", s.Name)
		assert.Equal(t, 1, s.RegionID)
		assert.Equal(t, 2, s.PlanID)
	}

	options := &BareMetalServerOptions{
		Script:   2,
		Snapshot: "alpha",
		SSHKey:   "key123",
		IPV6:     true,
	}
	s2, err := client.CreateBareMetalServer("test2", 4, 5, 6, options)
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

func Test_BareMetalServers_RenameBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.RenameBareMetalServer("123456789", "new-name")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_RenameBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.RenameBareMetalServer("123456789", "new-name"))
}

func Test_BareMetalServers_TagBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.TagBareMetalServer("123456789", "new-tag")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_TagBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.TagBareMetalServer("123456789", "new-tag"))
}

func Test_BareMetalServers_HaltBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.HaltBareMetalServer("123456789")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_HaltBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.HaltBareMetalServer("123456789"))
}

func Test_BareMetalServers_RebootBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.RebootBareMetalServer("123456789")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_RebootBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.RebootBareMetalServer("123456789"))
}

func Test_BareMetalServers_ReinstallBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.ReinstallBareMetalServer("123456789")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_ReinstallBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.ReinstallBareMetalServer("123456789"))
}

func Test_BareMetalServers_DeleteBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.DeleteBareMetalServer("123456789")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_DeleteBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.DeleteBareMetalServer("123456789"))
}

func Test_BareMetalServers_SetFirewallGroup_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.SetFirewallGroup("123456789", "123456789")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_SetFirewallGroup_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.SetFirewallGroup("123456789", "123456789"))
}

func Test_BareMetalServers_UnsetFirewallGroup_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.UnsetFirewallGroup("123456789")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_UnsetFirewallGroup_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.UnsetFirewallGroup("123456789"))
}

func Test_BareMetalServers_ChangeOSofBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.ChangeOSofBareMetalServer("123456789", 160)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_ChangeOSofBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.ChangeOSofBareMetalServer("123456789", 160))
}

func Test_BareMetalServers_ListOSforBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	os, err := client.ListOSforBareMetalServer("123456789")
	assert.Nil(t, os)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_ListOSforBareMetalServer_NoOS(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	os, err := client.ListOSforBareMetalServer("123456789")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, os)
}

func Test_BareMetalServers_ListOSforBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{
"179":{"OSID":179,"name":"CoreOS Stable","arch":"x64","family":"coreos","windows":false,"surcharge":"1.25"},
"127":{"OSID":127,"name":"CentOS 6 x64","arch":"x64","family":"centos","windows":false,"surcharge":"0.00"},
"124":{"OSID":124,"name":"Windows 2012 R2 x64","arch":"x64","family":"windows","windows":true,"surcharge":"5.99"}}`)
	defer server.Close()

	os, err := client.ListOSforBareMetalServer("123456789")
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, os) {
		assert.Equal(t, 3, len(os))

		assert.Equal(t, 127, os[0].ID)
		assert.Equal(t, "CentOS 6 x64", os[0].Name)
		assert.Equal(t, "x64", os[0].Arch)
		assert.Equal(t, "centos", os[0].Family)
		assert.Equal(t, "0.00", os[0].Surcharge)

		assert.Equal(t, 179, os[1].ID)
		assert.Equal(t, "coreos", os[1].Family)
		assert.Equal(t, "CoreOS Stable", os[1].Name)
		assert.Equal(t, false, os[1].Windows)
		assert.Equal(t, "1.25", os[1].Surcharge)

		assert.Equal(t, 124, os[2].ID)
		assert.Equal(t, "windows", os[2].Family)
		assert.Equal(t, "Windows 2012 R2 x64", os[2].Name)
		assert.Equal(t, true, os[2].Windows)
		assert.Equal(t, "5.99", os[2].Surcharge)
	}
}

func Test_BareMetalServers_BandwidthOfBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	bandwidth, err := client.BandwidthOfBareMetalServer("123456789")
	assert.Nil(t, bandwidth)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_BandwidthOfBareMetalServer_NoOS(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	bandwidth, err := client.BandwidthOfBareMetalServer("123456789")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, bandwidth)
}

func Test_BareMetalServers_BandwidthOfBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{
    "incoming_bytes": [
        ["2014-06-10",81072581],["2014-06-11",222387466],
        ["2014-06-12",216885232],["2014-06-13",117262318]
    ],
    "outgoing_bytes": [
        ["2014-06-10",4059610],["2014-06-11",13432380],
        ["2014-06-12",2455005],["2014-06-13",1106963]
    ]}`)
	defer server.Close()

	bandwidth, err := client.BandwidthOfBareMetalServer("123456789")
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, bandwidth) {
		assert.Equal(t, 4, len(bandwidth))
		assert.Equal(t, "2014-06-10", bandwidth[0]["date"])
		assert.Equal(t, "81072581", bandwidth[0]["incoming"])
		assert.Equal(t, "4059610", bandwidth[0]["outgoing"])
		assert.Equal(t, "2014-06-12", bandwidth[2]["date"])
		assert.Equal(t, "216885232", bandwidth[2]["incoming"])
		assert.Equal(t, "2455005", bandwidth[2]["outgoing"])
	}
}

func Test_BareMetalServers_ChangeApplicationofBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	err := client.ChangeApplicationofBareMetalServer("123456789", "3")
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_ChangeApplicationofBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{no-response?!}`)
	defer server.Close()

	assert.Nil(t, client.ChangeApplicationofBareMetalServer("123456789", "3"))
}

func Test_BareMetalServers_ListApplicationsforBareMetalServer_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	apps, err := client.ListApplicationsforBareMetalServer("123456789")
	assert.Nil(t, apps)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalServers_ListApplicationsforBareMetalServer_NoOS(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	apps, err := client.ListApplicationsforBareMetalServer("123456789")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, apps)
}

func Test_BareMetalServers_ListApplicationsforBareMetalServer_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{
"2": {"APPID": "2","name": "WordPress","short_name": "wordpress","deploy_name": "WordPress on CentOS 6 x64","surcharge": 0},
"1": {"APPID": "1","name": "LEMP","short_name": "lemp","deploy_name": "LEMP on CentOS 6 x64","surcharge": 5}
}`)
	defer server.Close()

	apps, err := client.ListApplicationsforBareMetalServer("123456789")
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, apps) {
		assert.Equal(t, 2, len(apps))

		assert.Equal(t, "1", apps[0].ID)
		assert.Equal(t, "LEMP", apps[0].Name)
		assert.Equal(t, "lemp", apps[0].ShortName)
		assert.Equal(t, "LEMP on CentOS 6 x64", apps[0].DeployName)
		assert.Equal(t, float64(5), apps[0].Surcharge)

		assert.Equal(t, "2", apps[1].ID)
		assert.Equal(t, "WordPress", apps[1].Name)
		assert.Equal(t, "wordpress", apps[1].ShortName)
		assert.Equal(t, "WordPress on CentOS 6 x64", apps[1].DeployName)
	}
}
