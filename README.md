# go-task-cli

A Lightweight CLI task manager built with GO

## Add a Task

```bash

 go run main.go --task-cli=add -add "buy banana"
```

### Update a Task by ID

```bash
go run main.go --task-cli=update -update 10 "buy car BMW"
```

## Delete a Task by ID

```bash

 go run main.go --task-cli=delete -delete 6

```

[commit](https://github.com/Prakash-Ravichandran/go-task-cli/commit/c77f8f034df1d8e2a4c4c6302e2b6895cbe02ae0)

## Mark a Task in progress

```bash

go run main.go --task-cli=mark-in-progress -mark-in-progress 7

```

updates in tasks.json:

```json
{
  "id": 7,
  "description": "todo",
  "status": "in=progres",
  "createdat": "",
  "updatedat": "2026-05-10 20:02:04.72738 +0530 IST m=+0.002758601"
}
```

## Mark a Task as "done"

```bash
go run main.go --task-cli=mark-done --mark-done 8
```

updates in tasks.json:

```json
{
  "id": 8,
  "description": "todo",
  "status": "done",
  "createdat": "",
  "updatedat": "2026-05-11 10:54:09.7293626 +0530 IST m=+0.001138601"
}
```

## List a Task by status

```bash

 go run main.go --task-cli=list -list all
 go run main.go --task-cli=list -list done
 go run main.go --task-cli=list -lis todo
 go run main.go --task-cli=list -list in-progress

```
