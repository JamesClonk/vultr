package lib

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BareMetalPlans_GetBareMetalPlans_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	bareMetalPlans, err := client.GetBareMetalPlans()
	assert.Nil(t, bareMetalPlans)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalPlans_GetBareMetalPlans_NoBareMetalPlans(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	bareMetalPlans, err := client.GetBareMetalPlans()
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, bareMetalPlans)
}

func Test_BareMetalPlans_GetBareMetalPlans_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{
"30":{"METALPLANID":"30","name":"1024 MB RAM,20 GB SSD,2.00 TB BW","cpu_count":2,"ram":1024,"disk":"20","bandwidth_tb":2,"price_per_month":7},
"29":{"METALPLANID":"29","name":"768 MB RAM,15 GB SSD,1.00 TB BW","cpu_count":1,"ram":768,"disk":"15","bandwidth_tb":1,"price_per_month":5,"available_locations":[1,2,3]}}`)
	defer server.Close()

	bareMetalPlans, err := client.GetBareMetalPlans()
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, bareMetalPlans) {
		assert.Equal(t, 2, len(bareMetalPlans))

		assert.Equal(t, 29, bareMetalPlans[0].ID)
		assert.Equal(t, "768 MB RAM,15 GB SSD,1.00 TB BW", bareMetalPlans[0].Name)
		assert.Equal(t, 1, bareMetalPlans[0].CPUs)
		assert.Equal(t, 768, bareMetalPlans[0].RAM)
		assert.Equal(t, 5, bareMetalPlans[0].Price)
		assert.Equal(t, 1, bareMetalPlans[0].Regions[0])
		assert.Equal(t, 3, bareMetalPlans[0].Regions[2])

		assert.Equal(t, 30, bareMetalPlans[1].ID)
		assert.Equal(t, "1024 MB RAM,20 GB SSD,2.00 TB BW", bareMetalPlans[1].Name)
		assert.Equal(t, 2, bareMetalPlans[1].CPUs)
		assert.Equal(t, "20", bareMetalPlans[1].Disk)
		assert.Equal(t, 2, bareMetalPlans[1].Bandwidth)
		assert.Equal(t, 0, len(bareMetalPlans[1].Regions))
	}
}

func Test_BareMetalPlans_GetAvailableBareMetalPlansForRegion_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	bareMetalPlans, err := client.GetAvailableBareMetalPlansForRegion(1)
	assert.Nil(t, bareMetalPlans)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}

func Test_BareMetalPlans_GetAvailableBareMetalPlansForRegion_NoBareMetalPlans(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[]`)
	defer server.Close()

	bareMetalPlans, err := client.GetAvailableBareMetalPlansForRegion(2)
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, bareMetalPlans)
}

func Test_BareMetalPlans_GetAvailableBareMetalPlansForRegion_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `[29,30,3,27,28,11,13,81]`)
	defer server.Close()

	bareMetalPlans, err := client.GetAvailableBareMetalPlansForRegion(3)
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, bareMetalPlans) {
		assert.Equal(t, 8, len(bareMetalPlans))
	}
}
