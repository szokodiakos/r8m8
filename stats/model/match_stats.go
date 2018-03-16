package model

// MatchStats struct
type MatchStats struct {
	ReporterDisplayName    string
	WinnerMatchPlayerStats []MatchPlayerStats
	LoserMatchPlayerStats  []MatchPlayerStats
}
