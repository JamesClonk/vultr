package lib

import (
	"net/url"
	"sort"
	"strings"
)

// BackupSchedule represents a scheduled backup on a server
// see: server/backup_set_schedule, server/backup_get_schedule
type BackupSchedule struct {
	Enabled              bool   `json:"enabled"`
	CronType             string `json:"cron_type"`
	NextScheduledTimeUtc string `json:"next_scheduled_time_utc"`
	Hour                 int    `json:"hour"`
	Dow                  int    `json:"dow"`
	Dom                  int    `json:"dom"`
}

// Backup of a virtual machine
type Backup struct {
	ID          string `json:"BACKUPID"`
	Created     string `json:"date_created"`
	Description string `json:"description"`
	Size        string `json:"size"`
	Status      string `json:"status"`
}

type backups []Backup

func (s backups) Len() int      { return len(s) }
func (s backups) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s backups) Less(i, j int) bool {
	return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name)
}

// GetBackups retrieves a list of all backups on Vultr account
func (c *Client) GetBackups(id string, backupid string) (backupList []BackupSchedule, err error) {
	var backupMap map[string]Backup
	values := url.Values{
		"SUBID":    {id},
		"BACKUPID": {backupid},
	}

	if err := c.post(`backup/list`, values, &backupMap); err != nil {
		return nil, err
	}

	for _, backup := range backupMap {
		backups = append(backups, backup)
	}
	sort.Sort(backups(backups))
	return backups, nil
}
