package models

type LintResult struct {
	Issues []LintCodeIssue `json:"issues"`
}
