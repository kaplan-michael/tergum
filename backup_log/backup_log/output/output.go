package output

import (
	"bytes"
	"io"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/sikalabs/tergum/backup_log"
)

func BackupLogTable(l backup_log.BackupLog, writer io.Writer) {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{
		"Success",
		"Backup", "Backup Time",
		"Target", "Upload Time",
		"Error",
		"Time Total",
	})

	for _, log := range l.Events {
		var strStatus, strError string

		if log.Success {
			strStatus = "OK"
		} else {
			strStatus = "ERROR"
		}

		if log.Error == nil {
			strError = ""
		} else {
			strError = log.Error.Error()
		}

		table.Append([]string{
			strStatus,
			log.SourceName + ": " + log.BackupID,
			strconv.Itoa(log.BackupDuration) + "s" +
				" (+" + strconv.Itoa(log.BackupMiddlewaresDuration) + "s)",
			log.TargetName + ": " + log.TargetID,
			strconv.Itoa(log.TargetDuration) + "s" +
				" (+" + strconv.Itoa(log.TargetMiddlewaresDuration) + "s)",
			strError,
			strconv.Itoa(log.TotalDuration()) + "s",
		})
	}
	table.Render()
}

func BackupLogToString(l backup_log.BackupLog) string {
	buf := new(bytes.Buffer)
	BackupLogTable(l, buf)
	return buf.String()
}

func BackupLogToOutput(l backup_log.BackupLog) {
	BackupLogTable(l, os.Stdout)
}
