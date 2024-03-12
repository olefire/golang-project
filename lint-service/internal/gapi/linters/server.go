package linters

import (
	"context"
	"lint-service/internal/models"
	"lint-service/internal/services"
	"lint-service/pkg/protos/gen"
)

type LintingService struct {
	services.LinterManagement
}

type Server struct {
	gen.UnimplementedLintingServiceServer
	LintingService
}

func NewGrpcServer(ls LintingService) *Server {
	return &Server{
		LintingService: ls,
	}
}

func (s *Server) LintCode(_ context.Context, file *gen.File) (*gen.LintResults, error) {
	sourceFile := models.SourceFile{Code: file.GetCode(), Language: models.ProgrammingLanguage(file.GetLanguage())}
	lintCode, err := s.LintingService.LintCode(sourceFile)
	if err != nil {
		return nil, err
	}

	var lintResults gen.LintResults

	for _, result := range lintCode {
		var lintResult gen.LintResult
		for _, issue := range result.Issues {
			lintResult.Result = append(lintResult.Result, &gen.LintCodeIssue{Line: int32(issue.Line), Message: issue.Message})
		}
		lintResult.Linter = string(result.Linter)
		lintResults.Results = append(lintResults.Results, &lintResult)
	}

	return &lintResults, nil
}
