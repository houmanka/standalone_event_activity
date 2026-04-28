# Standalone Activity Sample (Temporal Go SDK)

This project demonstrates how to run a **standalone activity** in Temporal (an activity started directly by a client, not by a workflow), while also showing that the activity can start a workflow.

## What this sample does

1. `worker/main.go` starts a worker on task queue `standalone-activity-helloworld`.
2. The worker registers:
   - `Greeting.Activity` from `activity/activity.go`
   - `GreetingSample` workflow and workflow activities from `workflow/`
3. `starter/starter.go` executes `Greeting.Activity` directly with `client.ExecuteActivity(...)`.
4. Inside that standalone activity, `GreetingSample` is started via `client.ExecuteWorkflow(...)`.

The standalone activity returns `Hello Temporal!` and the started workflow runs its own activity chain (`GetGreeting`, `GetName`, `SayGreeting`).

## Project structure

```text
.
├── activity/
│   └── activity.go      # Standalone activity implementation
├── starter/
│   └── starter.go       # Starts standalone activity from the client
├── worker/
│   └── main.go          # Worker process and registrations
├── workflow/
│   ├── activity.go      # Activities used by GreetingSample workflow
│   └── sample.go        # GreetingSample workflow
├── Dockerfile           # Temporal CLI image used by Makefile
├── Makefile             # Helpers to start a local Temporal dev server
└── go.mod
```

## Prerequisites

- Go (project `go.mod` targets `go 1.24.12`)
- Docker (if using `make start-dev`)

## Start a local Temporal server

From the repository root:

```bash
make start-dev
```

This builds the `Dockerfile` image and runs Temporal dev server with ports:

- `7233` Temporal gRPC endpoint
- `8233` Web UI
- `40259` HTTP metrics

## Run the sample

Open a second terminal in the repo root and start the worker:

```bash
go run ./worker
```

Open a third terminal and start the standalone activity client:

```bash
go run ./starter
```

## Expected output

From `starter`, you should see logs similar to:

- `Started standalone activity ...`
- `Activity result: Hello Temporal!`
- Activity list/count output for task queue `standalone-activity-helloworld`

From `worker`, you should see activity/workflow execution logs, including the workflow started from inside the standalone activity.

## Troubleshooting

- If you get connection errors (`Unable to create client`), make sure the Temporal dev server is running on `localhost:7233`.
- If no task is picked up, make sure both starter and worker use task queue `standalone-activity-helloworld`.
- If `make start-dev` fails, verify Docker is running.
- Client options are loaded with `envconfig.MustLoadDefaultClientOptions()`, so any custom Temporal endpoint/environment settings should be applied consistently to both starter and worker.

