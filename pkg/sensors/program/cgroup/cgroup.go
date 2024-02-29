// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

package cgroup

import (
	"fmt"

	"github.com/cilium/tetragon/pkg/sensors/program"
	"golang.org/x/sys/unix"
)

const (
	fgsCgroupPath = "/run/tetragon/cgroup2"
)

var (
	fgsCgroupFD = -1
)

func LoadSockOpt(
	bpfDir string,
	load *program.Program, verbose int,
) error {
	return LoadCgroupProgram(bpfDir, load, verbose)
}

func LoadCgroupProgram(
	bpfDir string,
	load *program.Program, verbose int) error {
	if fgsCgroupFD < 0 {
		fd, err := unix.Open(fgsCgroupPath, unix.O_RDONLY, 0)
		if err != nil {
			return fmt.Errorf("failed to open '%s': %w", fgsCgroupPath, err)
		}
		fgsCgroupFD = fd
	}
	return program.LoadProgram(bpfDir, load, program.RawAttach(fgsCgroupFD), nil, nil, verbose)
}
