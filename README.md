# Content of githooks

This repository includes a Git hook to prevent commits without a Jira issue key in the ***first line*** of a commit message.

The Git generated merge commit messages that start with **Merge** are always accepted.

If you are working on a branch that contains an issue key in its name, for example `feature/DS-17`, the commit message will be enhanced with that key at the beginning. Of course, only if the first line does not already contain that key.

The file `commit-msg` must be an executable file in the folder `.git/hooks` of a repository.
If you changed this default folder via the `core.hooksPath` git configuration variable, you need to install it there.

For more information see https://git-scm.com/docs/githooks

# Installation for single Repositories

For a simple installation of the Git hook in a repository you can use the script `install-jira-git/hook`.

If you do not pass a folder name to the script, the hook is installed in the repository where the current directory is located.

You can install the script in multiple repositories, by passing the repository folders to the script, e.g.
```shell
$ install-jira-git-hook repo1 repo2 repo3
```

# Restriction to specific Jira projects

In case the issue key must belong to certain Jira projects, you can specify the Jira project key list. Just use the option `-p` or `--projects`.

# Usage
```
Usage: install-jira-git-hook [ OPTIONS ] [ dir ... ]

Install a Git hook to search for a Jira issue key
in the commit message or branch name.

Options:
  -y, --yes       Override existing commit-msg files
  -p, --projects  Let the hook only accept keys for these Jira projects
                  e.g. --projects=DS,MYJIRA,MARS
  -h, --help      Show this help
```

# Global Installation

Sometimes you have multiple repositories in a folder, and they all should use the same hooks. This can be done via the Git configuration variable `core.hooksPath`.

Let's assume you work for two projects. Project **Alpha** uses a Jira project with the key `ALPHA`. The second project is **Beta** and uses a Jira project with the key `BETA`.

Create the following file / folder structure

```
~
├── .gitconfig
├── .gitconfig-alpha
├── .gitconfig-beta
└── .githooks
    ├── alpha
    │   └── commit-msg
    └── beta
        └── commit-msg
```

in the file `.gitconfig` you include further configuration settings depending on the location of the repositories. In the example below, the configuration file `.gitconfig-alpha` is read for all Git repositories that have `~/work/project-alpha/` as parent folder.

```shell
[includeIf "gitdir:~/work/project-alpha/"]
    path = .gitconfig-alpha
[includeIf "gitdir:~/work/project-beta/"]
    path = .gitconfig-beta
```

The file `.gitconfig-alpha` is then used to set the Git hooks folder to a place with the Git hooks for the project **Alpha**.

```
[core]
    hooksPath=~/.githooks/alpha
```

The folder `~/.githooks/alpha` finally contains the `commit-msg` file of this repository with the entry `PROJECTS="(ALPHA)"`.

The settings for the second project are analog.
