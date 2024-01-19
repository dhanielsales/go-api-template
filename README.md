## Getting Started

1. Download the needed packages `task install`
2. Create and populate the `app.env` file with the needed environment variables, can be found in `app.env.example`

## Dependencies, Tools and Infra

First, install the Taskfile to run the tasks and provide other tools: 

Taskfile:

https://taskfile.dev/installation/

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

Then, install the all other tools with the command:

```bash
task setup-tools
```

Now, create docker containers with dependencies:

```bash
task setup-infra
```

## Deployment

[PENDING]
- Checar GO_ENV para deploy

https://railway.app/pricing
