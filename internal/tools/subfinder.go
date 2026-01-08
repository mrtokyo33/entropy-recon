package tools

import "context"

type SubfinderResult struct {
	Host   string
	Source string
}

type SubfinderTool interface {
	Run(ctx context.Context, domain string) ([]SubfinderResult, error)
}
