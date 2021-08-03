# Gitall

Run git command for all git repositories in the directory.

## Installation

1. Download the release
2. Unzip the release file
3. Copy binary file to `/usr/local/bin`

## Usage

### Clone
Clone all repositories from group

```bash
gitall clone \
   --scm=gitlab \
   --base-url=http://your_gitlab_domain \
   --group=your_group_name \
   --token=your_token
```

### Git Command
Command for git

```bash
gitall [git subcommands]
```

e.g
```bash
gitall checkout -b foo
gitall branch -D foo
gitall add .
gitall commit -m 'foo bar'
```

### Merge request
Create merge request

```bash
# Creates MR(develop -> master) for all projects in the current folder
gitall mr develop master \
   --base-url=http://gitlab.your_domain \
   --token=your_token
```

## Subcommands & Flags

```bash
$ gitall -h
NAME:
   GitAll - Git command for multiple repositories

USAGE:
   gitall [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   mr       Multiple create gitlab merge request
   clone    Multiple clone
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

Clone

```
$ gitall clone -h
NAME:
   gitall clone - Multiple clone

USAGE:
   gitall clone [command options] [arguments...]

OPTIONS:
   --scm value       SCM type
   --base-url value  SCM base url
   --group value     SCM group
   --token value     SCM token
   --help, -h        show help (default: false)
```

Merge request

```
$ gitall mr -h
NAME:
   gitall mr - Multiple create gitlab merge request

USAGE:
   gitall mr [command options] [arguments...]

OPTIONS:
   --base-url value  SCM base url
   --token value     SCM token
   --help, -h        show help (default: false)
```