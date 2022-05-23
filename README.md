# Content of githooks

This repository includes a Git hook to prevent commits without a Jira issue key in the ***first line*** of a commit message.

The Git generated merge commit messages that start with **Merge** are always accepted.

If you are working on a branch that contains an issue key in its name, for example `feature/DS-17`,
the commit message will be enhanced with that key at the beginning.
Of course, only if the first line does not already contain that key.

<a name="global-installation"></a>
# Global Installation with githooks (recommended)

Sometimes you have multiple repositories in a folder, and they all should use the same hooks.
This can be done via the Git configuration variable `core.hooksPath`.

Let's assume you work for two projects. Project **Alpha** uses a Jira project with the key `ALPHA`.
The second project is **Beta** and uses a Jira project with the key `BETA`.

To simplify the installation and creation of the folders and files needed, you can use the `githooks`
CLI tool. It is a Linux command line application, which is developed with Go.
Once it is build, you can launch it directly.

```shell
$ curl -sfL https://raw.githubusercontent.com/stefan-niemeyer/githooks/main/install.sh | sh
$ chmod +x githooks
$ mv githooks <folder-in-your-PATH>
```

## Initialization
When you use `githooks` for the first time, you need to set up the usage by calling
```shell
$ githooks init
```

TODO: @Bai: Please describe what `githooks init` does. What is --init for ??? 

## Adding a Project
TODO: @Bai: Please describe what `githooks add` does. 

## Deleting a Project
TODO: @Bai: Please describe what `githooks delete` does. 

## Listing the Projects
TODO: @Bai: Please describe what `githooks list` does. 

## Under the Hood
After adding the projects **Alpha** and **Beta** `githooks` will have created the following
file / folder structure

```
~
├── .gitconfig
├── .gitconfig-alpha
├── .gitconfig-beta
└── .githooks
    ├── commit-msg
    └── githooks.log
    
```

The file `.gitconfig` will contain the configuration settings depending on the location of the
repositories. In the example below, the configuration file `.gitconfig-alpha` is read for
all Git repositories that have `~/work/project-alpha` as parent folder.
The setting for the project **Beta** are analog.

```
# settings in .gitconfig

[includeIf "gitdir:~/work/project-alpha/"]
    path = .gitconfig-alpha
[includeIf "gitdir:~/work/project-beta/"]
    path = .gitconfig-beta
```

The files `.gitconfig-alpha` and `.gitconfig-beta` set the Git hooks folder and the allowed Jira project keys.
It might look like this.

```
# settings in .gitconfig-alpha

[core]
    hooksPath=~/.githooks
[user]
    jiraProjects=ALPHA
```

The variable `core.hooksPath` is set to the folder with the shared hooks `~/.githooks`.

The Git variable `user.jiraProjects` is used to set different Jira Project keys for the projects.
This is a custom defined variable.

Configuring the allowed Jira project keys via `git config` provides the same flexibility as setting them using the
`PROJECTS` variable in the hook script. You can e.g. use regular expressions like `(GAMMA|DELTA)`.

# Installation for a Single Repository

The file `commit-msg` must be an executable file in the folder `.git/hooks` of a repository.
If you changed this default folder via the Git configuration variable `core.hooksPath`,
the script will be installed there.

For more information see [Global Installation](#global-installation) or https://git-scm.com/docs/githooks

For a simple installation of the Git hook directly in a repository you can use the script `install-jira-git-hook`.

If you do not pass a folder name to the script, the hook is installed in the repository where the current
directory is located.

You can install the script in multiple repositories, by passing the repository folders to the script, e.g.
```shell
$ ./install-jira-git-hook repo1 repo2 repo3
```

## Restriction to specific Jira projects

In case the issue key must belong to certain Jira projects, you can specify the Jira project key list.
Just use the option `-p` or `--projects`.

## Usage
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
