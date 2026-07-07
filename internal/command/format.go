// Copyright IBM Corp. 2026
// SPDX-License-Identifier: CC0-1.0

package command

import (
	"strings"

	"github.com/hashicorp/terraform/internal/tfdiags"
)

type FormatCommand struct {
	Meta
}

func (c *FormatCommand) Run(args []string) int {
	c.showDiagnostics(tfdiags.Sourceless(
		tfdiags.Error,
		"Command \"terraform format\" does not exist",
		"You likely want to use \"terraform fmt\".",
	))
	return 1
}

func (c *FormatCommand) Help() string {
	helpText := `
This command is to help users that are not used to the right format command name.
`
	return strings.TrimSpace(helpText)
}

func (c *FormatCommand) Synopsis() string {
	return "Helper command for to indicate the right format command name."
}
