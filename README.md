# git-hooks-template
My Handy Custom [Git Hooks](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks) Template. I've also written a tiny blog for Git Hooks focusing on the specific samples I've uploaded here. You can check this tiny blog from [here]().

## How to Enable the Hooks
The default path for Git hooks is `.git/hooks` and Git looks for special-named scripts in there, namely looks for a script named as pre-push in our case. If such a script exist in there then git pre-push hook is enabled and Git is going to run that script before each push.

### Share Your Hooks also in Remote
As can be seen this folder is under `.git` so the hook will not be visible to the others if you work on any repository with hooks collaboratively. If you want to get your hook on the remote repository as well then you need to configure the `core.hooksPath` of your local git repository as follows:

```
git config core.hooksPath hooks
```
Above command sets the hooks folder for Git to a folder named `hooks` in the current directory, you can set to any path you wish. Now Git is going to look for hooks in this folder instead of `.git/hooks`.

## How to Bypass after enabling
If you don't want your push to be reviewed by pre-push hook then you can use:

```git push --no-verify```

This is going to push your new commits to remote directly.

### Who is this template for?
Anybody who likes to utilize handy automation constructs provided by Git to automate some additional checks and running errands before applying a certain git command. 

### What does it do?
Currently there are only two custom git hooks but I can enrich later if I happen to use any other git hook for a super useful requirement.

#### [pre-push](https://github.com/ysyesilyurt/git-hooks-template/blob/main/hooks/pre-push)
This custom template currently has a generic `task` and `task queue` structure implementation in itself. Using the structure here, _any_ `pre-push` task can be converted to a `task` implementation just by adding a new `bash function` and then can be used in _any_ custom `task queue` in _any_ order.

You can also check [this](https://github.com/ysyesilyurt/git-hooks-template/blob/main/go-git-hooks-sample/hooks/pre-push) sample `pre-push` hook for a mono repository that has both backend (with Go) and frontend codes for more insight. Tasks defined in this sample hook are:
* `check_basename` Checks whether working directory is set correctly by Git, Common Task.
* `check_upstream` Checks whether current `HEAD` is behind upstream or not, for any push to go through, the local `HEAD` should be ahead of the upstream branch, Common Task.
* `auto_generate_backend_mocks` Checks the files that has committed changes since last push, if detects `interfaces` with an existing mock and has changes, regenerates the mock interfaces for them, Backend Task.
* `run_tests` Runs all implemented tests of the backend repository
* `commit_backend_mock_changes` Commits all the regenerated mock interfaces (if any), Backend Task.
* `run_cypress_tests` Runs all cypress tests defined in frontend repository

#### [pre-commit](https://github.com/ysyesilyurt/git-hooks-template/blob/main/hooks/pre-commit)
This custom template currently handles minor errands that can be important before applying each commit.

You can also check [this](https://github.com/ysyesilyurt/git-hooks-template/blob/main/go-git-hooks-sample/hooks/pre-commit) sample `pre-commit` hook. Currently it only formats (`gofmt`) files before committing.
