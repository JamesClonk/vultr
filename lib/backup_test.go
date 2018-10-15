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

	snapshots, err := client.GetBackups("123456789", "1")
	if err != nil {
		t.Error(err)
	}
	if assert.NotNil(t, snapshots) {
		assert.Equal(t, 2, len(snapshots))

		assert.Equal(t, "543d34149403a", snapshots[0].ID)
		assert.Equal(t, "2014-10-14 12:40:40", snapshots[0].Created)
		assert.Equal(t, "Automatic server backup", snapshots[0].Description)
		assert.Equal(t, "42949672960", snapshots[0].Size)
		assert.Equal(t, "complete", snapshots[0].Status)

		assert.Equal(t, "543d340f6dbce", snapshots[1].ID)
		assert.Equal(t, "2014-10-13 16:11:46", snapshots[1].Created)
		assert.Equal(t, "a", snapshots[1].Description)
		assert.Equal(t, "10000000", snapshots[1].Size)
		assert.Equal(t, "complete", snapshots[1].Status)
	}
}
