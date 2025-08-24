# Testing Envelope Stateful Sessions

This demo shows how to test Envoy's envelope stateful session functionality with a Python Flask backend.

## Quick Start

1. **Navigate to the demo directory**:
   ```bash
   cd .vscode/envelope-session-demo
   ```

2. **Start the services with Docker Compose**:
   ```bash
   docker-compose up --build
   ```

2. **Test envelope sessions**:
   ```bash
   # Basic envelope session request
   curl 'http://localhost:8080/?envelope=true'
   
   # Make multiple requests to see session persistence
   curl 'http://localhost:8080/?envelope=true'
   curl 'http://localhost:8080/?envelope=true'
   
   # Post data with session
   curl -X POST -H 'Content-Type: application/json' \
     -d '{"test": "data"}' 'http://localhost:8080/data?envelope=true'
   
   # View all sessions
   curl http://localhost:8080/sessions
   
   # Check backend health
   curl http://localhost:8000/health
   ```

3. **Stop services** (from the demo directory):
   ```bash
   docker-compose down
   ```

## Expected Behavior

- Each request with `?envelope=true` creates or updates a session
- Session data is encoded as base64 JSON in the `x-session-id` header
- Request count increments with each request from the same session
- Posted data is stored in the session

## Architecture

```
Client → Envoy Proxy (Port 8080) → Flask Backend (Port 8000)
```

- **Envoy**: Configured with envelope stateful session support
- **Backend**: Python Flask application that handles envelope session parsing
- **Session Format**: `base64(json({"session_id": "...", "request_count": N, "server_timestamp": "..."}))`

## Configuration Files

- `docker-compose.yml` - Docker Compose configuration
- `Dockerfile.backend` - Backend container definition
- `envoy.yaml` - Envoy proxy configuration with envelope session support
- `backend.py` - Python Flask backend with envelope session handling
