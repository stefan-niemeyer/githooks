# Content of githooks

This repository includes a Git hook to prevent commits without a Jira issue key in the ***first line*** of a commit message.

The Git generated merge commit messages that start with **Merge** are always accepted.

If you are working on a branch that contains an issue key in its name, for example `feature/DS-17`,
the commit message will be enhanced with that key at the beginning.
Of course, only if the first line does not already contain that key.

<a name="global-installation"></a>
# 1. Global Installation with githooks (recommended)

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

## 1.1 Initialization
When you use `githooks` for the first time, you need to set up the usage by calling
```shell
$ githooks init
```

The `init` command will create the Git hook script `commit-msg` under `$HOME/.githooks`. Further configuration files
will be created under `$HOME/.githooks/config`. This folder is called the **githooks config folder** in this documentation.
Here you will find files such as `githooks.json`. 
If you are not familiar with `githooks`, please don't touch these files.

## 1.2 Adding a Workspace

```shell
$ githooks add
```
The **add** command will add a new workspace to the `githooks` workspace list by a given workspace name, 
the Jira project key, and it's workspace folder.
Once it's done, it will generate a config file under the githooks config folder with the following pattern: `gitconfig-<PROJECT_NAME>` 
Because `githooks` is basically a Git extension, it will append extra config at the bottom of `$HOME/.gitconfig`.

## 1.3 Listing the Workspaces
By calling **list** `githooks` prints a list with all workspaces it 'manages'.

## 1.4 Deleting a Workspace
**delete** will list all `githooks` workspaces. The user can select one of the workspaces to delete by moving the cursor.
If a workspace was removed successfully, it's config files will also be removed or updated by `githooks`.
Here again, if you are not familiar with `githooks`, just do not touch these config files,
otherwise `githooks` will lose the control of it. 

# 2 Under the Hood
Let's assume, you have called `githooks init` and added the workspaces **Alpha** and **Beta** with `githooks add`.
Then `githooks` will have created the following
file / folder structure

```
~
├── .gitconfig
└── .githooks
    ├── commit-msg
    └── config
        ├── gitconfig-alpha
        ├── gitconfig-beta
        └── githooks.json
    
```

The file `.gitconfig` will contain the configuration settings depending on the location of the
repositories. In the example below, the configuration file `gitconfig-alpha` is read for
all Git repositories that have `~/work/ws-alpha` as a parent folder.
The settings for the project **Beta** are analog.

```
# settings in .gitconfig

[includeIf "gitdir:~/work/ws-alpha/"]
    path = .githooks/config/gitconfig-alpha
[includeIf "gitdir:~/work/ws-beta/"]
    path = .githooks/config/gitconfig-beta
```

The files `gitconfig-alpha` and `gitconfig-beta` set the Git hooks folder and the allowed Jira project keys.
It might look like this.

```
# settings in gitconfig-alpha

[core]
    hooksPath=~/.githooks
[user]
    jiraProjects=ALPHA
```

The variable `core.hooksPath` is set to the folder with the shared hooks `~/.githooks`.

The Git variable `user.jiraProjects` is used to set different Jira Project keys for the workspaces.
This is a custom defined variable and not part of the Git core.

Configuring the allowed Jira project keys via `git config` provides the same flexibility as setting them using the
`PROJECTS` variable in the hook script. You can e.g. use regular expressions like `(GAMMA|DELTA)`.

# 3. Installation for a Single Repository

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

## 3.1 Restriction to specific Jira projects

In case the issue key must belong to certain Jira projects, you can specify the Jira project key list.
Just use the option `-p` or `--projects`.

## 3.2 Usage
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
