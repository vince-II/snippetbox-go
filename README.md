# Let's Go — Web Application in Go

This project is a hands-on implementation of the concepts from  
[_Let's Go: Build a Professional Web Application in Go_ by Alex Edwards](https://lets-go.alexedwards.net/).

The goal is to learn how to design, build, and deploy a production-ready web application using Go.

---

## 📖 About

This project follows along with the book and covers topics such as:

- Project structure and organization
- Routing and middleware
- Templates and dynamic content
- Database integration (PostgreSQL/MySQL/SQLite)
- Authentication and session management
- Security best practices
- Deployment strategies

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.22+
- A database (e.g., PostgreSQL)
- `git`

### Clone the repository

```bash
git clone https://github.com/yourusername/lets-go-webapp.git
cd lets-go-webapp
```

Run the application

```bash
go run ./cmd/web
```

### The app should now be running on http://localhost:4000

🛠 Project Structure

```bash
├── cmd/  # Application entry points
│ └── web/ # Main web application
├── internal/ # Reusable packages
│ ├── models/ # Database models and queries
│ ├── validator/ # Form validation helpers
│ └── ...
├── ui/ # Templates, static files, etc.
├── go.mod
└── go.sum
```

📚 Progress Tracking

- [x] Chapter 1: Setting Up a Web Application
- [x] Chapter 2: Templates
- [ ] Chapter 3: Database Models
- [ ] Chapter 4: Middleware
- [ ] Chapter 5: Authentication
- [ ] Chapter 6: Security
- [ ] Chapter 7: Deployment

💡 Notes

This repository is for learning purposes only.
The code may closely follow the book examples, with small modifications or extensions as practice.
