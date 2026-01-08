package tools

import "context"

type MockSubfinder struct{}

func (m *MockSubfinder) Run(ctx context.Context, domain string) ([]string, error) {
	return []string{
		"www." + domain,
		"api." + domain,
		"admin." + domain,
	}, nil
}
