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

## TODO 
- [x] refactor log
- [x] Add observability
- [x] Add conversational id to logs
- [x] Add CI/CD
- [x] Change Http layer for echo
- [ ] Add tests
- [ ] Add simple auth using package and an interface to able to setup your own auth or keyclock
