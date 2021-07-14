# Gitall

Run git command for all git repositories in the directory.

## Installation

1. Download the release
2. Unzip the release file
3. Copy binary file to `/usr/local/bin`

## Usage

clone all repositories from group

```
gitall clone \
--scm=gitlab \
--base-url=http://your_gitlab_domain \
--group=your_group_name \
--token=your_token
```

command for git

```
gitall [git subcommands]
```

e.g
```
gitall checkout -b foo
gitall branch -D foo
gitall add .
gitall commit -m 'foo bar'
```

## Subcommands & Flags

```
$ gitall -h
NAME:
   GitAll - Git command for multiple repositories

USAGE:
   gitall [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   clone    Multiple clone
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

clone

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