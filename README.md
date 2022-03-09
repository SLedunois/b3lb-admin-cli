# b3lbctl

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/487a9c5102c8465ebbfd36ca1b62194e)](https://www.codacy.com/gh/SLedunois/b3lbctl/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=SLedunois/b3lbctl&amp;utm_campaign=Badge_Grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/487a9c5102c8465ebbfd36ca1b62194e)](https://www.codacy.com/gh/SLedunois/b3lbctl/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=SLedunois/b3lbctl&amp;utm_campaign=Badge_Coverage)
[![Code linting](https://github.com/SLedunois/b3lbctl/actions/workflows/lint.yml/badge.svg)](https://github.com/SLedunois/b3lbctl/actions/workflows/lint.yml)
[![Unit tests](https://github.com/SLedunois/b3lbctl/actions/workflows/unit_test.yml/badge.svg)](https://github.com/SLedunois/b3lbctl/actions/workflows/unit_test.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sledunois/b3lbctl)
![GitHub](https://img.shields.io/github/license/SLedunois/b3lbctl)

The b3lbctl command line tool lets you control [b3lb](https://github.com/SLedunois/b3lb) clusters.

For configuration, b3lbctl looks for a file named config in the $HOME/.b3lb.yaml directory. You can specify other b3blconfig files by setting the --config flag.

## Installation

Download last release from [release page](https://github.com/SLedunois/b3lbctl/releases).

Copy the binary into `/usr/local/bin`.

## Usage

```bash
Manage your B3LB cluster from the command line

Usage:
  b3lbctl <command> [flags]
  b3lbctl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  instances   Manage B3LB instances

Flags:
      --config string   config file (default is $HOME/.b3lb.yaml) (default "$HOME/.b3lb.yaml")
  -h, --help            help for b3lbctl

Use "b3lbctl [command] --help" for more information about a command.
```