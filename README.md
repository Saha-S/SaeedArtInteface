# art-interface (Go)

This README contains two versions:

1. **Version 1 (Summary)**: quick reviewer orientation
2. **Version 2 (Complete Manual)**: full setup, usage, testing, and review guide

---

## Version 1 — Summary

### What it is

`art-interface` is a Go web server that provides a browser interface for text-art decode/encode.

### Required behavior

- `GET /` returns the main page with HTTP `200`
- `POST /decoder` accepts form input and returns:
  - HTTP `202` for valid input
  - HTTP `400` for malformed input
- The page reloads after submit and displays the generated output
- The interface displays the latest HTTP response code

### Extra behavior implemented

- Server-side encoder mode (`decode` / `encode`)
- Multiline mode using literal `\n`
- Dedicated CSS styling for a clean UI

### Quick run

```bash
cd art-interface
go run ./cmd/server
```

Open: `http://localhost:8080`

---

## Version 2 — Complete Manual

### 1) Project overview

This project is a server-rendered interface for artists who want to use text-art without command-line knowledge.

The server owns all transformation logic (decode/encode), and HTML forms submit requests to backend endpoints.

### 2) Endpoints

| Method | Route      | Purpose                                 | Success | Error |
| ------ | ---------- | --------------------------------------- | ------- | ----- |
| GET    | `/`        | Render the main web interface           | `200`   | `404` for unknown route |
| POST   | `/decoder` | Decode/encode submitted form input      | `202`   | `400` for malformed input |

Any additional unsupported request receives an appropriate HTTP status.

### 3) Project structure

```text
art-interface/
├── cmd/server/main.go
├── internal/art/service.go
├── internal/http/handler.go
├── internal/art/service_test.go
├── internal/http/handler_test.go
├── web/templates/index.html
├── web/static/style.css
└── README.md
```

### 4) Prerequisites

- Go `1.22+`

Check version:

```bash
go version
```

### 5) Build and run

From workspace root:

```bash
cd art-interface
go mod tidy
go build ./cmd/server
go run ./cmd/server
```

Default URL: `http://localhost:8080`

Optional custom address:

```bash
ADDR=:9090 go run ./cmd/server
```

### 6) Usage

1. Open `http://localhost:8080`
2. Enter your input in the textarea
3. Choose mode:
   - `Decode` for encoded input (example: `[3 A][2 B]`)
   - `Encode` for plain text input (example: `AAABB`)
4. (Optional) enable `Multiline (literal \n)`
5. Click `Generate`
6. Read the output and shown HTTP status code

### 7) Examples

#### Decode

Input:

```text
[5 #][5 -_]-[5 #]
```

Output:

```text
#####-_-_-_-_-_-#####
```

#### Encode

Input:

```text
AAABB
```

Output:

```text
[3 A][2 B]
```

#### Multiline decode

Input (literal `\n`):

```text
[3 A]\n[2 B]
```

Output:

```text
AAA
BB
```

### 8) Testing

Automated tests:

```bash
cd art-interface
go test ./...
```

Manual endpoint checks:

```bash
# GET / => 200
curl -i -s http://localhost:8080/ | head -n 1

# POST /decoder valid => 202
curl -i -s -X POST http://localhost:8080/decoder \
  -d 'input=[3 A][2 B]&mode=decode' | head -n 1

# POST /decoder malformed => 400
curl -i -s -X POST http://localhost:8080/decoder \
  -d 'input=[x #]&mode=decode' | head -n 1
```

### 9) Additional features (implemented)

- Encoder mode is computed by server logic
- Multiline support with literal `\n`
- Styled interface with dedicated CSS

### 10) Troubleshooting

- `400` on decode: check format is `[count value]` with a space
- Wrong multiline output: ensure you typed literal `\n`
- Port busy: run with `ADDR=:9090`

