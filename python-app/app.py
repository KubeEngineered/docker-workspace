from flask import Flask, jsonify, request
import os
import socket
import datetime

app = Flask(__name__)

# Fetch environment variables with robust defaults
APP_VERSION = os.getenv("APP_VERSION", "v2.0.0")
ENVIRONMENT = os.getenv("ENVIRONMENT", "lab-testing")
HOSTNAME = socket.gethostname()

@app.route("/")
def home():
    """Root endpoint showing app identity and container hostname."""
    return jsonify({
        "service": "rilly-docker-lab-service",
        "status": "online",
        "message": "Welcome to the fully overhauled Python Docker lab application!",
        "version": APP_VERSION,
        "environment": ENVIRONMENT,
        "container_hostname": HOSTNAME,
        "timestamp": datetime.datetime.utcnow().isoformat()
    })

@app.route("/health")
def health():
    """Standard health check endpoint for container probes."""
    return jsonify({
        "status": "healthy",
        "uptime_check": "passed",
        "timestamp": datetime.datetime.utcnow().isoformat()
    }), 200

@app.route("/info")
def info():
    """System and runtime diagnostic info."""
    return jsonify({
        "python_version": os.sys.version,
        "platform": os.sys.platform,
        "server_hostname": HOSTNAME,
        "active_version": APP_VERSION
    })

@app.route("/echo", methods=["POST"])
def echo_payload():
    """Allows testing POST requests and JSON payload handling inside containers."""
    incoming_data = request.get_json(silent=True) or {}
    return jsonify({
        "received_data": incoming_data,
        "processed_by_container": HOSTNAME,
        "status": "success"
    }), 201

if __name__ == "__main__":
    print(f"Starting Rilly Lab Service [{APP_VERSION}] on host {HOSTNAME}...")
    app.run(host="0.0.0.0", port=5000)