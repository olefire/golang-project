package models

type LintResult struct {
	Issues []LintCodeIssue `json:"issues"`
	Linter `json:"linter"`
}
