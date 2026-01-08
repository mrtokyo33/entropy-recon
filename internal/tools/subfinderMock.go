package tools

import "context"

type MockSubfinder struct{}

func (m *MockSubfinder) Run(ctx context.Context, domain string) ([]SubfinderResult, error) {
	return []SubfinderResult{
		{Host: "api." + domain, Source: "crtsh"},
		{Host: "www." + domain, Source: "virustotal"},
	}, nil
}
