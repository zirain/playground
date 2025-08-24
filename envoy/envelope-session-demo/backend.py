#!/usr/bin/env python3

from flask import Flask, request, jsonify, make_response
import uuid
import json
import base64

app = Flask(__name__)

# Simple in-memory session store
sessions = {}

def parse_envelope_session(session_header):
    """Parse envelope session header format: base64(json({session_id, ...}))"""
    try:
        decoded = base64.b64decode(session_header).decode('utf-8')
        envelope = json.loads(decoded)
        return envelope
    except Exception as e:
        print(f"Failed to parse envelope session: {e}")
        return None

def create_envelope_session(session_id, additional_data=None):
    """Create envelope session header format: base64(json({session_id, ...}))"""
    envelope = {"session_id": session_id}
    if additional_data:
        envelope.update(additional_data)
    
    json_str = json.dumps(envelope)
    encoded = base64.b64encode(json_str.encode('utf-8')).decode('utf-8')
    return encoded

def is_envelope_format(session_header):
    """Check if session header is in envelope format (base64 encoded JSON)"""
    try:
        decoded = base64.b64decode(session_header).decode('utf-8')
        json.loads(decoded)
        return True
    except:
        return False

@app.route('/')
def home():
    session_header = request.headers.get('x-session-id')
    
    if session_header:
        session_id = None
        
        # Check if it's envelope format or plain header format
        if is_envelope_format(session_header):
            # Envelope format
            envelope = parse_envelope_session(session_header)
            if envelope and 'session_id' in envelope:
                session_id = envelope['session_id']
                
                if session_id in sessions:
                    sessions[session_id]['request_count'] += 1
                    response_data = {
                        "message": f"Hello! Continuing envelope session {session_id}",
                        "session_id": session_id,
                        "format": "envelope",
                        "envelope": envelope,
                        "request_count": sessions[session_id]['request_count'],
                        "session_data": sessions[session_id]
                    }
                    
                    # Update envelope with new request count
                    new_envelope = create_envelope_session(
                        session_id, 
                        {"request_count": sessions[session_id]['request_count']}
                    )
                    
                    response = make_response(jsonify(response_data))
                    response.headers['x-session-id'] = new_envelope
                    return response
        else:
            # Plain header format
            session_id = session_header
            if session_id in sessions:
                sessions[session_id]['request_count'] += 1
                response_data = {
                    "message": f"Hello! Continuing header session {session_id}",
                    "session_id": session_id,
                    "format": "header",
                    "request_count": sessions[session_id]['request_count'],
                    "session_data": sessions[session_id]
                }
                
                response = make_response(jsonify(response_data))
                response.headers['x-session-id'] = session_id
                return response
    
    # Create new session
    session_id = str(uuid.uuid4())
    sessions[session_id] = {
        "created_at": "now",
        "request_count": 1,
        "user_data": {}
    }
    
    # Check if we should use envelope format (for testing)
    use_envelope = request.args.get('envelope', 'false').lower() == 'true'
    
    if use_envelope:
        envelope_session = create_envelope_session(session_id, {"request_count": 1})
        response_data = {
            "message": f"Hello! New envelope session created: {session_id}",
            "session_id": session_id,
            "format": "envelope",
            "request_count": 1
        }
        response = make_response(jsonify(response_data))
        response.headers['x-session-id'] = envelope_session
    else:
        response_data = {
            "message": f"Hello! New header session created: {session_id}",
            "session_id": session_id,
            "format": "header",
            "request_count": 1
        }
        response = make_response(jsonify(response_data))
        response.headers['x-session-id'] = session_id
    
    return response

@app.route('/status')
def status():
    session_header = request.headers.get('x-session-id')
    
    response_data = {
        "endpoint": "/status",
        "session_header_present": bool(session_header),
        "session_header": session_header,
        "total_sessions": len(sessions)
    }
    
    if session_header:
        if is_envelope_format(session_header):
            envelope = parse_envelope_session(session_header)
            response_data["format"] = "envelope"
            response_data["envelope"] = envelope
            if envelope and 'session_id' in envelope:
                response_data["session_id"] = envelope['session_id']
                response_data["session_valid"] = envelope['session_id'] in sessions
        else:
            response_data["format"] = "header"
            response_data["session_id"] = session_header
            response_data["session_valid"] = session_header in sessions
    
    return jsonify(response_data)

@app.route('/data', methods=['GET', 'POST'])
def data():
    session_header = request.headers.get('x-session-id')
    session_id = None
    
    # Extract session ID regardless of format
    if session_header:
        if is_envelope_format(session_header):
            envelope = parse_envelope_session(session_header)
            if envelope and 'session_id' in envelope:
                session_id = envelope['session_id']
        else:
            session_id = session_header
    
    if request.method == 'POST':
        data = request.get_json() or {}
        
        # Store data in session if session exists
        if session_id and session_id in sessions:
            sessions[session_id]['user_data'].update(data)
            sessions[session_id]['request_count'] += 1
            
        response_data = {
            "method": "POST",
            "received_data": data,
            "session_id": session_id,
            "session_stored": session_id in sessions if session_id else False,
            "message": f"Data received for session {session_id}" if session_id else "Data received (no session)"
        }
    else:
        response_data = {
            "method": "GET",
            "session_id": session_id,
            "sample_data": {"key": "value", "number": 42},
            "session_data": sessions.get(session_id, {}) if session_id else {},
            "message": f"Data for session {session_id}" if session_id else "Data (no session)"
        }
        
        # Increment request count for valid sessions
        if session_id and session_id in sessions:
            sessions[session_id]['request_count'] += 1
    
    response = make_response(jsonify(response_data))
    
    # Pass through session header
    if session_header:
        response.headers['x-session-id'] = session_header
    
    return response

@app.route('/sessions')
def list_sessions():
    """List all active sessions"""
    return jsonify({
        "active_sessions": len(sessions),
        "sessions": {k: v for k, v in sessions.items()}
    })

if __name__ == '__main__':
    print("Starting Flask backend server...")
    print("Supports both header-based and envelope session formats:")
    print("  Header format: x-session-id: <session-uuid>")
    print("  Envelope format: x-session-id: base64(json({session_id: <uuid>}))")
    print("")
    print("Test commands:")
    print("  # Header-based sessions (default):")
    print("  curl http://localhost:8080/")
    print("  curl http://localhost:8080/status")
    print("  curl http://localhost:8080/sessions")
    print("")
    print("  # Envelope sessions (add ?envelope=true):")
    print("  curl 'http://localhost:8080/?envelope=true'")
    print("  curl -X POST -H 'Content-Type: application/json' -d '{\"test\": \"data\"}' http://localhost:8080/data")
    print()

@app.route('/health', methods=['GET'])
def health():
    """Health check endpoint for Docker"""
    return jsonify({"status": "healthy", "service": "backend"}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8000, debug=True)
