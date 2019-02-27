package lib

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Backups_GetBackups_OK(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusOK, `{
    "543d340f6dbce": {
        "BACKUPID": "543d340f6dbce",
        "date_created": "2014-10-13 16:11:46",
        "description": "a",
        "size": "10000000",
        "status": "complete"
    },
    "543d34149403a": {
        "BACKUPID": "543d34149403a",
        "date_created": "2014-10-14 12:40:40",
        "description": "Automatic server backup",
        "size": "42949672960",
        "status": "complete"
    }
}`)
	defer server.Close()

	backups, err := client.GetBackups("123456789", "asdf")
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, backups) {
		assert.Equal(t, 2, len(backups))

		assert.Equal(t, "543d34149403a", backups[0].ID)
		assert.Equal(t, "2014-10-14 12:40:40", backups[0].Created)
		assert.Equal(t, "Automatic server backup", backups[0].Description)
		assert.Equal(t, "42949672960", backups[0].Size)
		assert.Equal(t, "complete", backups[0].Status)

		assert.Equal(t, "543d340f6dbce", backups[1].ID)
		assert.Equal(t, "2014-10-13 16:11:46", backups[1].Created)
		assert.Equal(t, "a", backups[1].Description)
		assert.Equal(t, "10000000", backups[1].Size)
		assert.Equal(t, "complete", backups[1].Status)
	}
}

func Test_Servers_Backups_GetBackups_Error(t *testing.T) {
	server, client := getTestServerAndClient(http.StatusNotAcceptable, `{error}`)
	defer server.Close()

	backups, err := client.GetBackups("123456789", "asdf")
	assert.Nil(t, backups)
	if assert.NotNil(t, err) {
		assert.Equal(t, `{error}`, err.Error())
	}
}
