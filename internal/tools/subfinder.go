package tools

import "context"

type SubfinderTool interface {
	Run(ctx context.Context, domain string) ([]string, error)
}
