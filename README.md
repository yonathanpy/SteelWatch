<!-- Shield -->
<img src="https://img.icons8.com/ios-filled/500/000000/shield.png" width="200"/>

# SteelWatch

SteelWatch is an advanced HTTP honeypot and monitoring framework written in Go. Designed for security engineers, research labs, and organizations, it simulates web endpoints and monitors suspicious activity in real time. SteelWatch captures detailed request metadata, fingerprints clients, and scores potentially malicious behavior.

---

## Overview

SteelWatch operates by intercepting HTTP requests to simulated endpoints. It records the following for every request:

* IP address of the client
* HTTP method (GET, POST, etc.)
* Requested path or resource
* User-Agent string
* Computed threat score based on request characteristics
* Unique SHA1 fingerprint for client identification
* Timestamp of the interaction

Additional capabilities include:

* Header analysis (Accept, Encoding, Referrer anomalies)
* Detection of common scanning tools (e.g., curl, sqlmap, nmap scripts)
* Path probing detection (e.g., `/admin`, `/login`, `.env`, `/wp-admin`)
* Rate-based scoring for repeated requests
* Lightweight behavioral correlation across sessions

---

## Architecture

```text
SteelWatch/
├── main.go            # Entry point
├── core/
│   ├── scoring.go     # Threat scoring engine
│   ├── fingerprint.go # Client fingerprint generation
├── handlers/
│   ├── http.go        # HTTP request handler
│   ├── api.go         # REST API for events
├── storage/
│   ├── memory.go      # Thread-safe in-memory storage
├── types/
│   ├── event.go       # Event data structure
```

---

## Module Breakdown

* **core**
  Implements scoring heuristics and fingerprint generation.

  * Rule-based scoring (headers, paths, methods)
  * Weighted scoring system for anomaly detection
  * SHA1-based deterministic client fingerprinting

* **handlers**
  Entry point for all HTTP interactions.

  * Central request interception via `http.HandleFunc`
  * Request parsing and normalization
  * JSON REST API for retrieving captured events

* **storage**
  In-memory concurrent event store.

  * Mutex-protected data structures
  * Fast append and retrieval operations
  * Designed for low-latency ingestion

* **types**
  Strong typing for system-wide consistency.

  * Event schema definition
  * Extensible structures for future integrations

---

## Features

* High-performance HTTP listener using Go’s `net/http`
* Deterministic client fingerprinting (IP + headers → SHA1)
* Heuristic-based threat scoring engine
* Detection of reconnaissance and enumeration patterns
* Real-time event capture and in-memory indexing
* REST API for downstream processing or SIEM integration
* Minimal dependencies and easy deployment footprint
* Extensible architecture for custom detection rules

---

## Execution Flow

```text
Incoming HTTP Request
        │
http.HandleFunc("/") → HTTPHandler.Handle()
        │
Normalize & parse request
        │
Generate fingerprint (SHA1)
        │
Apply scoring heuristics
        │
Persist event (thread-safe store)
        │
Optional API exposure (/api/events)
        │
Return static response ("OK")
```

---

## Installation and Execution

1. Install Go (version 1.20 or higher)
2. Clone the repository:

   ```bash
   git clone https://github.com/yonathanpy/SteelWatch.git
   cd SteelWatch
   ```
3. Run the application:

   ```bash
   go run main.go
   ```

---

## Disclaimer

SteelWatch is intended for use in controlled and authorized environments only. Unauthorized use outside these contexts may violate local laws and regulations.
