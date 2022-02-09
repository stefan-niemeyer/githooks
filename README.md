# Content of githooks

This repository includes a Git hook to prevent commits without a Jira issue key in the ***first line*** of a commit message.

If you are working on a branch that contains an issue key in its name, for example ``feature/DS-17``, the commit message will be enhanced with that key at the beginning. Of course, only if the first line does not already contain that key.

The file ``commit-msg`` must be an executable file in the folder ``.git/hooks`` of a repository.
If you changed this default folder via the ``core.hooksPath`` git configuration variable, you need to install it there.

For more information see https://git-scm.com/docs/githooks

# Installation

For a simple installation of the Git hook in a repository you can use the script ``install-jira-git/hook``.

If you do not pass a folder name to the script, the hook is installed in the repository where the current directory is located.

You can install the script in multiple repositories, by passing the repository folders to the script, e.g.
```shell
$ install-jira-git-hook repo1 repo2 repo3
```

# Restriction to specific Jira projects

In case the issue key must belong to certain Jira projects, you can specify the Jira project key list. Just use the option ``-p`` or ``--projects``.

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

