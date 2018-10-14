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

// GetBackups retrieves a list of all backups on Vultr account
func (c *Client) GetBackups() (backupList []BackupSchedule, err error) {
	var backupSchedulesMap map[string]BackupSchedule
	if err := c.get(`backup/list`, &backupSchedulesMap); err != nil {
		return nil, err
	}

	for _, backup := range backupSchedulesMap {
		backupList = append(backupList, backup)
	}
	sort.Sort(backups(backupList))
	return backupList, nil
}
