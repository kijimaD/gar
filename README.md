# gar

gar is easy review reply tool.

## Install

```shell
go install github.com/kijimaD/gar@main
```

## Workflow

1. Get a code review
  + <img src="https://user-images.githubusercontent.com/11595790/205493035-6a0be592-3d0c-4ca2-ac02-43dc5b1e1417.png" width="40%">
2. Fix problem and commit. Include correspond review comment URL in commit message.<br> e.g. `https://github.com/kijimaD/gar/pull/1#discussion_r1037682054`<br>
3. `git push`
4. Run gar. Auto reply all review comment.
  + e.g. `GH_TOKEN=xxxxx gar -n 1` -n is PR number
  + <img src="https://user-images.githubusercontent.com/11595790/205493043-97d7b855-94fb-487e-b5e9-be9039d3918c.png" width="40%">

## Development

```shell
make help
```

## TODO

- dry run
- check duplicate
- pretty UI
