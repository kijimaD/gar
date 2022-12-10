────● █▀▀ ▄▀█ █▀█ <br>
●──── █▄█ █▀█ █▀▄ is easy review reply tool.

<img src="https://user-images.githubusercontent.com/11595790/205955931-31b88633-e3ba-4aa3-96f1-660c41c114ee.png" width="30%" align=right>

## Install

```shell
go install github.com/kijimaD/gar@main
```

To use gar, you need to get a GitHub API token with an account which has enough permissions to comment.For a private repository you need `repo` scope and for a public repository you need `public_repo` scope.

```shell
export GH_TOKEN="....."
```

Or set it in github.token in gitconfig:

```shell
git config --global github.token "....."
```

## Workflow

1. Receive a code review.
  + <img src="https://user-images.githubusercontent.com/11595790/205493035-6a0be592-3d0c-4ca2-ac02-43dc5b1e1417.png" width="40%">
2. Fix problem and commit. Include correspond review comment URL in commit message.<br> e.g. `https://github.com/kijimaD/gar/pull/1#discussion_r1037682054`<br>
3. `git push`
4. Run gar on working directory. Show dry run result and yes/no prompt.
  + e.g. `gar 1` (1 is PR number)
```
$ gar 1
+-----+-------------------+----------------+------+
| IDX |      COMMIT       | LINKED COMMENT | SEND |
+-----+-------------------+----------------+------+
|  00 | 6eed27d test: thi |                | No   |
|  01 | 369a79d feat: thi | this is review | Yes  |
+-----+-------------------+----------------+------+
? Send reply?[yes/no]: 
  ▸ yes
    no
```
5. Answer prompt and send reply
6. Check GitHub
  + <img src="https://user-images.githubusercontent.com/11595790/205493043-97d7b855-94fb-487e-b5e9-be9039d3918c.png" width="40%">

## Development

```shell
make help
```
