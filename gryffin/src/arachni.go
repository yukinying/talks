// Copyright 2015, Yahoo Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arachni

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/yahoo/gryffin"
)

type Fuzzer struct{}

func (s *Fuzzer) Fuzz(g *gryffin.Scan) (count int, err error) {

	var cookies []string
	for _, c := range g.Cookies {
		cookies = append(cookies, c.String())
	}

	args := []string{
		"--checks", "xss*", "--output-only-positives",
		"--http-request-concurrency", "1",
		"--http-request-timeout", "10000",
		"--timeout", "00:03:00",
		"--scope-dom-depth-limit", "0",
		"--scope-directory-depth-limit", "0",
		"--scope-page-limit", "1",
		"--audit-with-both-methods",
		"--report-save-path", "/dev/null",
		"--snapshot-save-path", "/dev/null",
	}

	args = append(args, g.Request.URL.String())
	cmd := exec.Command("arachni", args...)
	//...
}

func (s *Fuzzer) extract(g *gryffin.Scan, output string) (count int) {

	for _, l := range strings.Split(output, "\n") {
		l = strings.TrimSpace(l)
		switch {
		case strings.HasPrefix(l, "[~] Affected page"):
			g.Logm("Arachni.Findings", l)
			count++
		}
	}

	return
}
