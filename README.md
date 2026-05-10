# go-task-cli

A Lightweight CLI task manager built with GO

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
