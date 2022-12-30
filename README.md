<p align="center">
<img src="assets/bbsctl.png" alt="logo" />
</p>

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/1958ff373bdb47659cb97888992b322b)](https://www.codacy.com/gh/bigblueswarm/bbsctl/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=bigblueswarm/bbsctl&amp;utm_campaign=Badge_Grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/1958ff373bdb47659cb97888992b322b)](https://www.codacy.com/gh/bigblueswarm/bbsctl/dashboard?utm_source=github.com&utm_medium=referral&utm_content=bigblueswarm/bbsctl&utm_campaign=Badge_Coverage)
[![Code linting](https://github.com/bigblueswarm/bbsctl/actions/workflows/lint.yml/badge.svg)](https://github.com/SLedunois/b3lbctl/actions/workflows/lint.yml)
[![Unit tests](https://github.com/bigblueswarm/bbsctl/actions/workflows/unit_test.yml/badge.svg)](https://github.com/SLedunois/b3lbctl/actions/workflows/unit_test.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/bigblueswarm/bbsctl)
![GitHub](https://img.shields.io/github/license/bigblueswarm/bbsctl)

The bbsctl command line tool lets you control [BigBlueSwarm](https://github.com/bigblueswarm/bigblueswarm) clusters.

For configuration, bbsctl looks for a file named config in the `$HOME/.bigblueswarm/.bbsctl.yml` directory. You can specify other bbsctl files by setting the `--config` flag.

## Installation

Download last release from [release page](https://github.com/bigblueswarm/bbsctl/releases).

Copy the binary into `/usr/local/bin`.

## Usage

```bash
Manage your BigBlueSwarm cluster from the command line

Usage:
  bbsctl <command> [flags]
  bbsctl [command]

Available Commands:
  apply        Apply a configuration to bigblueswarm server using a file
  cluster-info Get overall cluster information
  completion   Generate the autocompletion script for the specified shell
  delete       Delete a specific resource
  describe     Show details of a specific resource or group of resources
  get          Display a resource
  help         Help about any command
  init         Initialize a resource

Flags:
      --config string   config file (default is $HOME/.bigblueswarm/.bbsctl.yml) (default "$HOME/.bigblueswarm/.bbsctl.yml")
  -h, --help            help for bbsctl

Use "bbsctl [command] --help" for more information about a command.
```

See the [documentation](docs/bbsctl.md) for more details.

### Init configuration
Use `b3lbctl init config` to initialize `b3lbctl` configuration.
```bash
Create bbsctl if not exists and initialize a basic configuration

Usage:
  bbsctl init config [flags]

Flags:
  -b, --bbs string    BigBlueSwarm url
  -d, --dest string   Configuration file folder destination (default "$HOME/.bigblueswarm")
  -h, --help          help for config
  -k, --key string    BigBlueSwarm admin api key
```
<a href="https://asciinema.org/a/2qz4H250QCzMCbioMqEsnkuVE" target="_blank"><img src="https://asciinema.org/a/2qz4H250QCzMCbioMqEsnkuVE.svg" height="300" /></a>


### Manage instances
`bbsctl` provide some command to manage your cluster instances:
- `bbsctl init instances` that create instances file if it does not exists
```bash
Create instances list file if it does not exists

Usage:
  bbsctl init instances [flags]

Flags:
  -d, --dest string   File folder destination (default "$HOME/.bigblueswarm")
  -h, --help          help for instances
```
- `b3lbctl apply -f instances.yml`
```bash
Apply a configuration to bigblueswarm server using a file

Usage:
  bbsctl apply -f [filepath] [flags]

Flags:
  -f, --file string   resource file path
  -h, --help          help for apply

Global Flags:
      --config string   config file (default is $HOME/.bigblueswarm/.bbsctl.yml) (default "$HOME/.bigblueswarm/.bbsctl.yml")
```
- `b3lbctl get instances` display all BigBlueButton instances found in your B3lb cluster
```bash
Display all BigBlueButton instances available in your BigBlueSwarm cluster

Usage:
  bbsctl get instances [flags]

Flags:
  -c, --csv    csv output
  -h, --help   help for instances
  -j, --json   json output

Global Flags:
      --config string   config file (default is $HOME/.bigblueswarm/.bbsctl.yml) (default "$HOME/.bigblueswarm/.bbsctl.yml")
```
<a href="https://asciinema.org/a/nlueByPpY8E9g7DElk85jRcqH" target="_blank"><img src="https://asciinema.org/a/nlueByPpY8E9g7DElk85jRcqH.svg" height="300" /></a>

### Manage tenants
B3lb is a multi tenant load balancer and `b3lbctl` offer tools to manage tenants.
- `b3lbctl init tenant` initialize a tenant file
```bash
Initialize a new bigblueswarm tenant configuration file if not exits

Usage:
  bbsctl init tenant [flags]

Flags:
  -d, --dest string   File folder destination (default "$HOME/.bigblueswarm")
  -h, --help          help for tenant
      --host string   Tenant hostname
```
- `b3lbctl apply -f tenant.yml`
```bash
Apply a configuration to bigblueswarm server using a file

Usage:
  bbsctl apply -f [filepath] [flags]

Flags:
  -f, --file string   resource file path
  -h, --help          help for apply

Global Flags:
      --config string   config file (default is $HOME/.bigblueswarm/.bbsctl.yml) (default "$HOME/.bigblueswarm/.bbsctl.yml")
```
- `b3lbctl get tenants` display all tenants found in B3lb cluster
```bash
Display all BigBlueSwarm tenants available in your BigBlueSwarm cluster

Usage:
  bbsctl get tenants [flags]

Flags:
  -c, --csv    csv output
  -h, --help   help for tenants
  -j, --json   json output

Global Flags:
      --config string   config file (default is $HOME/.bigblueswarm/.bbsctl.yml) (default "$HOME/.bigblueswarm/.bbsctl.yml")
```
- `b3lbctl describe tenant` describe a tenant
```bash
Describe a given BigBlueSwarm tenant.

Usage:
  bbsctl describe tenant <hostname> [flags]

Flags:
  -h, --help   help for tenant

Global Flags:
      --config string   config file (default is $HOME/.bigblueswarm/.bbsctl.yml) (default "$HOME/.bigblueswarm/.bbsctl.yml")
```
<a href="https://asciinema.org/a/gB55DijzRQCl40or7bgpa3sv5" target="_blank"><img src="https://asciinema.org/a/gB55DijzRQCl40or7bgpa3sv5.svg" height="300" /></a>

### Check your cluster
`b3lbctl` let you monitor your cluster using `cluster-info` command.
```bash
Get overall cluster information. It display all instances with %CPU, %MEM, Active meetings, Active paricipants and API status

Usage:
  bbsctl cluster-info [flags]

Flags:
  -h, --help   help for cluster-info

Global Flags:
      --config string   config file (default is $HOME/.bigblueswarm/.bbsctl.yml) (default "$HOME/.bigblueswarm/.bbsctl.yml")
``` 
<a href="https://asciinema.org/a/Nqec46FDprZpUzbP43oa940Xb" target="_blank"><img src="https://asciinema.org/a/Nqec46FDprZpUzbP43oa940Xb.svg" height="300" /></a>

### Check B3lb configuration
`b3lbctl` let you check your b3lb application configuration remotely. Even if you use a configuration file or consul provider, `b3lbctl` display your application configuration using `describe config` command.
```bash
describe B3LB configuration.

Usage:
  b3lbctl describe config [flags]

Flags:
  -h, --help   help for config

Global Flags:
      --config string   config file (default is $HOME/.b3lb/.b3lbctl.yml) (default "$HOME/.b3lb/.b3lbctl.yml")
```
<a href="https://asciinema.org/a/idAmn6AjWZIh77x1oWcaK8Yiv" target="_blank"><img src="https://asciinema.org/a/idAmn6AjWZIh77x1oWcaK8Yiv.svg" height="300" /></a>