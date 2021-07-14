# Gitall

use git command for all git repositories in the directory.

## Usage

clone all repositories from group

```
gitall clone --scm=gitlab --base-url=http://your_gitlab_domain --group=your_group_name --token=your_token
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