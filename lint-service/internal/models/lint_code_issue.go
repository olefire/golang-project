package models

type LintCodeIssue struct {
	Message string `json:"message"`
	Line    int    `json:"line"`
}
