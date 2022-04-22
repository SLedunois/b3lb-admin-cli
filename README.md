# b3lbctl

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/487a9c5102c8465ebbfd36ca1b62194e)](https://www.codacy.com/gh/SLedunois/b3lbctl/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=SLedunois/b3lbctl&amp;utm_campaign=Badge_Grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/487a9c5102c8465ebbfd36ca1b62194e)](https://www.codacy.com/gh/SLedunois/b3lbctl/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=SLedunois/b3lbctl&amp;utm_campaign=Badge_Coverage)
[![Code linting](https://github.com/SLedunois/b3lbctl/actions/workflows/lint.yml/badge.svg)](https://github.com/SLedunois/b3lbctl/actions/workflows/lint.yml)
[![Unit tests](https://github.com/SLedunois/b3lbctl/actions/workflows/unit_test.yml/badge.svg)](https://github.com/SLedunois/b3lbctl/actions/workflows/unit_test.yml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sledunois/b3lbctl)
![GitHub](https://img.shields.io/github/license/SLedunois/b3lbctl)

The b3lbctl command line tool lets you control [b3lb](https://github.com/SLedunois/b3lb) clusters.

For configuration, b3lbctl looks for a file named config in the `$HOME/.b3lb.yaml` directory. You can specify other b3blconfig files by setting the `--config` flag.

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
  apply        Apply a configuration to b3lb server using a file
  cluster-info Get overall cluster information
  completion   Generate the autocompletion script for the specified shell
  delete       Delete a specific resource
  describe     Show details of a specific resource or group of resources
  get          Display a resource
  help         Help about any command
  init         Initialize a resource

Flags:
      --config string   config file (default is $HOME/.b3lb/.b3lbctl.yml) (default "$HOME/.b3lb/.b3lbctl.yml")
  -h, --help            help for b3lbctl

Use "b3lbctl [command] --help" for more information about a command.
```

### Init configuration
Use `b3lbctl init config` to initialize `b3lbctl` configuration.
```bash
Create b3lbctl if not exists and initialize a basic configuration

Usage:
  b3lbctl init config [flags]

Flags:
  -b, --b3lb string   B3lb url
  -d, --dest string   Configuration file folder destination (default "$HOME/.b3lb")
  -h, --help          help for config
  -k, --key string    B3lb admin api key
```
[![asciicast](https://asciinema.org/a/489335.svg)](https://asciinema.org/a/489335)


### Manage instances
`b3lbctl` provide some command to manage your cluster instances:
- `b3lbctl init instances` that create instances file if it does not exists
```bash
Create instances list file if it does not exists

Usage:
  b3lbctl init instances [flags]

Flags:
  -d, --dest string   File folder destination (default "$HOME/.b3lb")
  -h, --help          help for instances
```
- `b3lbctl apply -f instances.yml`
```bash
Apply a configuration to b3lb server using a file

Usage:
  b3lbctl apply -f [filepath] [flags]

Flags:
  -f, --file string   resource file path
  -h, --help          help for apply

Global Flags:
      --config string   config file (default is $HOME/.b3lb/.b3lbctl.yml) (default "$HOME/.b3lb/.b3lbctl.yml")
```
- `b3lbctl get instances` display all BigBlueButton instances found in your B3lb cluster
```bash
Display all BigBlueButton instances available in your B3LB cluster

Usage:
  b3lbctl get instances [flags]

Flags:
  -c, --csv    csv output
  -h, --help   help for instances
  -j, --json   json output

Global Flags:
      --config string   config file (default is $HOME/.b3lb/.b3lbctl.yml) (default "$HOME/.b3lb/.b3lbctl.yml")
```
[![asciicast](https://asciinema.org/a/489340.svg)](https://asciinema.org/a/489340)

### Manage tenants
B3lb is a multi tenant load balancer and `b3lbctl` offer tools to manage tenants.
- `b3lbctl init tenant` initialize a tenant file
```bash
Initialize a new b3lb tenant configuration file if not exits

Usage:
  b3lbctl init tenant [flags]

Flags:
  -d, --dest string   File folder destination (default "$HOME/.b3lb")
  -h, --help          help for tenant
      --host string   Tenant hostname
```
- `b3lbctl apply -f tenant.yml`
```bash
Apply a configuration to b3lb server using a file

Usage:
  b3lbctl apply -f [filepath] [flags]

Flags:
  -f, --file string   resource file path
  -h, --help          help for apply

Global Flags:
      --config string   config file (default is $HOME/.b3lb/.b3lbctl.yml) (default "$HOME/.b3lb/.b3lbctl.yml")
```
- `b3lbctl get tenants` display all tenants found in B3lb cluster
```bash
Display all B3lb tenants available in your B3LB cluster

Usage:
  b3lbctl get tenants [flags]

Flags:
  -c, --csv    csv output
  -h, --help   help for tenants
  -j, --json   json output

Global Flags:
      --config string   config file (default is $HOME/.b3lb/.b3lbctl.yml) (default "$HOME/.b3lb/.b3lbctl.yml")
```
- `b3lbctl describe tenant` describe a tenant
```bash
Describe a given B3LB tenant.

Usage:
  b3lbctl describe tenant <hostname> [flags]

Flags:
  -h, --help   help for tenant

Global Flags:
      --config string   config file (default is $HOME/.b3lb/.b3lbctl.yml) (default "$HOME/.b3lb/.b3lbctl.yml")
```
[![asciicast](https://asciinema.org/a/489344.svg)](https://asciinema.org/a/489344)

### Check your cluster
`b3lbctl` let you monitor your cluster using `cluster-info` command.
```bash
Get overall cluster information. It display all instances with %CPU, %MEM, Active meetings, Active paricipants and API status

Usage:
  b3lbctl cluster-info [flags]

Flags:
  -h, --help   help for cluster-info

Global Flags:
      --config string   config file (default is $HOME/.b3lb/.b3lbctl.yml) (default "$HOME/.b3lb/.b3lbctl.yml")
``` 
[![asciicast](https://asciinema.org/a/489346.svg)](https://asciinema.org/a/489346)

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
[![asciicast](https://asciinema.org/a/489345.svg)](https://asciinema.org/a/489345)