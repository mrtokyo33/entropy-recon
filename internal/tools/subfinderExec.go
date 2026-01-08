package tools

import (
	"bufio"
	"context"
	"encoding/json"
	"os/exec"
)

type Subfinder struct {
	Path string
}

type subfinderJSON struct {
	Host   string `json:"host"`
	Source string `json:"source"`
}

func (s *Subfinder) Run(ctx context.Context, domain string) ([]SubfinderResult, error) {
	cmd := exec.CommandContext(
		ctx,
		s.Path,
		"-d", domain,
		"-silent",
		"-json",
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	results := []SubfinderResult{}
	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		var row subfinderJSON
		if err := json.Unmarshal(scanner.Bytes(), &row); err != nil {
			continue
		}

		results = append(results, SubfinderResult{
			Host:   row.Host,
			Source: row.Source,
		})
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}
