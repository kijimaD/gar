# gar

gar is easy review reply tool.

## Install

```shell
go install github.com/kijimaD/gar@main
```

## Workflow

1. Get a code review
2. Fix and commit. Include correspond review comment URL in commit message.
  - ex. `https://github.com/kijimaD/gar/pull/1#discussion_r1037682054`
3. Run gar and reply all review comment.

```shell
GH_TOKEN=xxxxx gar -n 1

# `-n {PR number}`
```

## development

```shell
make help
```
