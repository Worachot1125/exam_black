package enum

type EmergencyReportStatus string

const (
	StatusPending    EmergencyReportStatus = "pending"
	StatusInProcess  EmergencyReportStatus = "in_process"
	StatusFinished   EmergencyReportStatus = "finished"
)