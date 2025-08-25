# Envelope Session Demo

This folder contains a complete demo for testing Envoy's envelope stateful session functionality.

## Files

- `docker-compose.yml` - Docker Compose configuration for running the demo
- `Dockerfile.backend` - Docker configuration for the Python Flask backend
- `envoy.yaml` - Envoy proxy configuration with envelope session support
- `backend.py` - Python Flask backend that handles envelope sessions
- `stateful-session.md` - Detailed testing instructions

## Usage

See [stateful-session.md](stateful-session.md) for complete testing instructions.

## Setup

```bash
cd .vscode/envelope-session-demo
docker-compose up -d
```

## Quick Test

Then in another terminal, dump the response header `x-session-id` into a environment variable:
```bash
export SESSION_ID=$(curl 'http://localhost:8080/?envelope=true' -v | grep 'x-session-id' | sed 's/x-session-id: //')
```

send another request with header`x-session-id` and value from previous response:
```bash
curl 'http://localhost:8080/?envelope=true' -H "x-session-id: $SESSION_ID" -v
```