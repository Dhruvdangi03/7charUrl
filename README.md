# 📌 7charUrl

**7charUrl** is a simple and fast URL shortening service written in **Go**. It generates **7-character short links** for long URLs, stores them, and redirects visitors to the original URL when the short link is accessed.

⚡ Built with scalability and simplicity in mind.

---

## 🚀 Features

* 🔗 Shorten URLs into **7-character codes**
* 🧠 Persistent storage (in memory or database – *you can configure*)
* 🏃‍♂️ Fast HTTP API written in Go
* 🧩 Modular internal packages for easy extension

---

## 🧱 Tech Stack

✔️ Go (Golang)
✔ HTTP server
✔ Modular internal packages (business logic & helpers)

---

## 📁 Project Structure

```
7charUrl/
├── cmd/server/            # Server entrypoint
├── internal/              # Core logic (handlers, store, utils...)
├── go.mod                 # Go modules
├── go.sum
└── LICENSE
```

---

## 🚀 Getting Started

### 💡 Prerequisites

Make sure you have **Go (1.18+)** installed.

```bash
go version
```

---

## 🛠️ Install & Run

Clone the repository:

```bash
git clone https://github.com/Dhruvdangi03/7charUrl.git
cd 7charUrl
```

Build the server:

```bash
go build ./cmd/server
```

Run the server:

```bash
./server
```

By default it should start on a port like `:8080` (configure if needed).

---

## 📡 Usage Examples

### ➕ Create Short URL

**Request:**

```
POST /shorten
Content-Type: application/json

{
  "url": "https://example.com/some/long/path"
}
```

**Response:**

```json
{
  "shortUrl": "http://localhost:8080/AbC123x"
}
```

(*Replace with your actual host / format from your API implementation*)

---

### 🔁 Redirect

Visiting:

```
GET http://localhost:8080/AbC123x
```

Will redirect the user to the original long URL.

---

## 🧠 How It Works

1. API receives a long URL
2. It generates a random **7-character code**
3. Saves short code → long URL mapping
4. When the short code is accessed, it redirects to the original URL

---

## 🧪 Tests

You can run tests (if available):

```bash
go test ./...
```

(*Add test instructions if you have tests implemented*)

---

## ♻️ Deployment

You can deploy this app to any server/platform that supports Go binaries (Heroku, DigitalOcean, Railway, Fly.io, etc.).

---

## 📦 Configuration

| Setting  | Description                                      |
| -------- | ------------------------------------------------ |
| `PORT`   | Port number to serve on                          |
| `DB_URL` | Persistence database connection (if implemented) |

(*Adjust based on your actual config system*)

---

## 🧾 License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details. ([GitHub][1])

---

## ✨ Contributing

Contributions are welcome!
Feel free to open issues and submit pull requests.

---
