// Copyright (c) 2021, SailPoint Technologies, Inc. All rights reserved.

package root

import (
	"bytes"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sailpoint-oss/sailpoint-cli/internal/mocks"
)

// Expected number of subcommands to `sp` root command
const numRootSubcommands = 6

func TestNewRootCmd_noArgs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cmd := NewRootCmd(mocks.NewMockClient(ctrl))
	if len(cmd.Commands()) != numRootSubcommands {
		t.Fatalf("expected: %d, actual: %d", len(cmd.Commands()), numRootSubcommands)
	}

	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("error execute cmd: %v", err)
	}

	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatalf("error read out: %v", err)
	}

	if string(out) != cmd.UsageString() {
		t.Errorf("expected: %s, actual: %s", cmd.UsageString(), string(out))
	}
}

func TestNewRootCmd_completionDisabled(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cmd := NewRootCmd(mocks.NewMockClient(ctrl))

	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"completion"})

	if err := cmd.Execute(); err == nil {
		t.Error("expected command to fail")
	}
}
