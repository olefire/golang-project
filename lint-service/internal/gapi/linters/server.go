package linters

import (
	"context"
	"lint-service/internal/models"
	"lint-service/internal/services"
	gen2 "lint-service/pkg/protos/gen"
)

type LintingService struct {
	services.LinterManagement
}

type Server struct {
	gen2.UnimplementedLintingServiceServer
	LintingService
}

func NewGrpcServer(ls LintingService) *Server {
	return &Server{
		LintingService: ls,
	}
}

func (s *Server) LintCode(_ context.Context, file *gen2.File) (*gen2.LintResults, error) {
	sourceFile := models.SourceFile{Code: file.GetCode(), Language: models.ProgrammingLanguage(file.GetLanguage())}
	lintCode, err := s.LintingService.LintCode(sourceFile)
	if err != nil {
		return nil, err
	}

	var lintResults gen2.LintResults

	for _, result := range lintCode {
		var lintResult gen2.LintResult
		for _, issue := range result.Issues {
			lintResult.Result = append(lintResult.Result, &gen2.LintCodeIssue{Line: int32(issue.Line), Message: issue.Message})
		}
		lintResults.Results = append(lintResults.Results, &lintResult)
		lintResults.Linter = string(result.Linter)
	}

	return &lintResults, nil
}
