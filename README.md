# Github Workflow Simulation

A Go program that simulates the Workflow of a Github Repo done by **6 developer goroutines** pushing
to **3 distinct branch channels** which is monitored and managed by a branch listener.

## Architecture Overview

1. **Developers:** Six distinct goroutines (`Dev-1` through `Dev-6`) map to specific commits and pushed to their designated branch channels (`featureA`, `featureB` and `main`)

2. **Branch Listeners:** Three separate goroutines monitor the branches. Once a branch is drained and closed, the listeners triggers an automated `AUTO-MERGE` signal into the shared `mergeCh`-_unless_ it is the `main` branch.

3. **Merge Handler:** A final goroutine drains the `mergeCh`, outputs the merge log to console, and signals `main` via `done` channel when the workspace is synchronized

## Getting Started

### Prerequisites

- Go 1.22 or higher

### Installation and Execution

1. Clone the workspace directory or save the code to `main.go`.
2. Run the application directly from the terminal:

```bash
go run main.go
```

### Example Output

Since this is a concurrency program, the output may differ but it is expected

```text
[feature-B] Dev-3 pushing: feat: dark mode toggle
[feature-B] Dev-4 pushing: refactor: auth module
[main] Dev-5 pushing: chore: update deps
[main] Received commit: chore: update deps
[featureB] Received commit: feat: dark mode toggle
[feature-A] Dev-1 pushing: fix: login bug
[featureA] Received commit: fix: login bug
[feature-A] Dev-2 pushing: feat: add dashboard
[featureA] Received commit: feat: add dashboard
[main] Dev-6 pushing: docs: update README
[featureB] Received commit: refactor: auth module
AUTO-MERGE: featureA into main
AUTO-MERGE: featureB into main
[main] Received commit: docs: update README
```
