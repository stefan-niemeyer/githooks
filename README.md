# Content of githooks

This repository includes a Git hook to prevent commits without a Jira issue key in the commit message.  
Compliance with Conventional Commits can also be ensured.

Git-generated merge commit messages that start with **Merge** are always accepted.

If you are working on a branch that contains an issue key in its name, for example `feature/DS-17`, the commit message will be enhanced with that issue key. There are several format styles available to choose how the issue key will be inserted.

The insertion will only happen if the commit message does not already contain that key.

<a name="global-installation"></a>
# 1. Global Installation with githooks

Sometimes you have multiple repositories in a folder, and they all should use the same hooks.  
This can be done via the Git configuration variable `core.hooksPath`.

Let us assume you work on two projects. Project **Alpha** uses a Jira project with the key `ALPHA`.  
The second project is **Beta** and uses a Jira project with the key `BETA`.

To simplify the installation and creation of the folders and files needed, you can use the `githooks`
CLI tool. It is a command line application that is developed with Go.  
Once it is built, you can run it directly.

## Linux / macOS Installation
```shell
$ curl -sfL https://raw.githubusercontent.com/stefan-niemeyer/githooks/main/install.sh | sh
$ chmod +x githooks
$ mv githooks <folder-in-your-PATH>
```

## Windows Installation
```
irm https://raw.githubusercontent.com/stefan-niemeyer/githooks/main/install.ps1 | iex
```

## 1.1 Initialization

When you use `githooks` for the first time, you need to set up the usage by calling:

```shell
$ githooks init
```

The `init` command will create the Git hook command `commit-msg` under `$HOME/.githooks`.  
`commit-msg` is just a symbolic link to the `githooks` file. If you install a newer version of `githooks`, the actual Git hook command `commit-msg` will be updated as well, because it is just a link.

The behavior under Windows is different. `githooks init` will copy the file `githooks.exe` to `$HOME/.githooks/commit-msg.exe`.  
It is the same program with different names. The name defines the behavior when it is called.

Further configuration files will be created under `$HOME/.githooks/config`. This folder is called the **githooks config folder** in this documentation.  
Here you will find files such as `githooks.json`.  
If you are not familiar with `githooks`, please do not touch these files.

You can modify the content of this file via the `githooks` commands `add`, `edit`, and `delete`.

## 1.2 Updating the Git Hooks

By calling `githooks update` the Git hooks file `commit-msg` will be reinstalled.

If you move `githooks` to another folder under Linux, `githooks update` will recreate the symbolic link.

On a Windows machine, where `commit-msg.exe` is just a copy of `githooks.exe` and not a symbolic link, calling `githooks update` is necessary after installing a newer version of `githooks`.

## 1.3 Workspaces

A workspace for this application consists of the following elements:

- Name  
- ProjectKeyRE  
- CommitMessageStyle  
- AllowedTypes  
- Folder  

### Workspace Name

The `Name` is used to distinguish the workspace definitions.

### Project Key Regular Expression

`ProjectKeyRE` is a regular expression to determine if a commit message contains an issue key of the right project(s).  
If only one project is allowed, just type the project key, e.g. `ALPHA`.  
If you work with several Jira projects, you can use a regular expression like `(ALPHA|BETA|GAMMA)`.

Go supports the RE2 format. The syntax is explained in the official [Go documentation](https://pkg.go.dev/regexp/syntax).  
However, the example above should fit for most use cases.

If you do not specify a particular project key, `githooks` just checks for the existence of a Jira issue key, regardless of the project key.

### Commit Message Style

You can place a Jira issue key in a commit message wherever you want. If the project key matches `ProjectKeyRE`, the commit message will not be changed by `githooks`.

The strength of `githooks` is extracting the Jira issue key from a Git branch and inserting that issue key into the commit message.  
This is convenient and prevents typos.

The insertion of the issue key can be influenced. `githooks` supports several format styles.  
The table below shows how a commit message will be modified.

| Format Style   | Original                  | Result                                          |
|----------------|---------------------------|-------------------------------------------------|
| Plain Space    | add feature               | JIRA-123 add feature                            |
| Plain Colon    | add feature               | JIRA-123: add feature                           |
| Plain Brackets | add feature               | [JIRA-123] add feature                          |
| Conventional   | fix(parser)!: changes API | fix(parser)!: changes API<br><br>Refs: JIRA-123 |

#### Issue Key Positioning for Conventional Commits

The position where `githooks` inserts an issue key from a branch is special for the format style `Conventional`.  
The application uses [`git interpret-trailers`](https://git-scm.com/docs/git-interpret-trailers) to add the issue key in the footer of a commit message.

By default, `githooks` uses the trailer key `Refs`. This will produce an output like the example above.  
You can go to a Git repository and call:

```shell
git config set trailer.refs.key "Reference"
```

Then Git will change the key `Refs` passed by `githooks` to `Reference` as the key in the footer of the commit message.

### Allowed Types

The format style `Conventional` ensures compliance of a commit message with the Conventional Commits requirements.

There is a specification of Conventional Commits: https://www.conventionalcommits.org

The Angular project demands this style, see [Commit Message Format](https://github.com/angular/angular/blob/main/contributing-docs/commit-message-guidelines.md).

Another project is [Karma](https://karma-runner.github.io/1.0/dev/git-commit-msg.html).

As you can see, the projects differ in the allowed types in the commit message.  
For that purpose, `githooks` accepts a list of types that the commit message can or must contain.

This is defined per workspace, so you can adapt it to your needs.

### Folder

The configuration of a workspace is valid for all Git repositories in this folder and its subfolders.

However, the definition of a workspace with the folder `~/projects/alpha` takes precedence over the configuration of a workspace for the folder `~/projects`.

## 1.4 Adding a Workspace

```shell
$ githooks add
```

The **add** command will add a new workspace to the `githooks` workspace list.

Once it is done, it will generate a config file under the githooks config folder with the following pattern: `gitconfig-<PROJECT_NAME>`.  
Because `githooks` is essentially a Git extension, it will append extra configuration at the bottom of `$HOME/.gitconfig`.

## 1.5 Listing the Workspaces

By using the command **list**, `githooks` prints all workspaces it manages.

## 1.6 Editing a Workspace

The configuration elements of a workspace can be changed by using the command **edit**.

## 1.7 Deleting a Workspace

The **delete** command will list all `githooks` workspaces. You can select one of the workspaces to delete by moving the cursor.  
If a workspace was removed successfully, its config files will also be removed or updated by `githooks`.  
Again, if you are not familiar with `githooks`, just do not touch these config files; otherwise `githooks` will lose control of them.

# 2. Under the Hood

Let us assume you have called `githooks init` and added the workspaces **Alpha** and **Beta** with `githooks add`.  
Then `githooks` will have created the following file and folder structure:

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

The file `.gitconfig` will contain the configuration settings depending on the location of the repositories.  
In the example below, the configuration file `gitconfig-alpha` is read for all Git repositories that have `~/work/ws-alpha` as a parent folder.  
The settings for the project **Beta** are analogous.

```
# settings in .gitconfig

[includeIf "gitdir:~/work/ws-alpha/"]
    path = .githooks/config/gitconfig-alpha
[includeIf "gitdir:~/work/ws-beta/"]
    path = .githooks/config/gitconfig-beta
```

The files `gitconfig-alpha` and `gitconfig-beta` set the Git hooks folder and the allowed Jira project keys.  
They might look like this:

```
# settings in gitconfig-alpha

[core]
    hooksPath=~/.githooks
[user]
    jiraProjects=ALPHA
    commitMessageStyle=conventional
    allowedTypes=build,chore,ci,docs,feat,fix,ops,perf,refactor,style,test
```

The variable `core.hooksPath` is set to the folder with the shared hooks `~/.githooks`.

The Git variable `user.jiraProjects` is used to set different Jira project keys for the workspaces.  
This is a custom-defined variable and not part of the Git core. The same applies to the entries `commitMessageStyle`
and `allowedTypes`.


# 3. Related Projects

If you are only interested in the conventional commits format and do not need the Jira issue key existence, have
a look at [Bengt Brodersen's](https://github.com/qoomon) project [Git Conventional Commits](https://github.com/qoomon/git-conventional-commits)
His code allows the creation of change logs as well.

And check out [Bai Xia's](https://github.com/xiabai84) project [githooks](https://github.com/xiabai84/githooks).

Most of the workspace configuration was contributed by him. Thank's again Bai.
