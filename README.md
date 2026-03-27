# Todo CLI — Go

A command-line todo manager with colored terminal output, backed by a local `.todos.json` file.

---

## 🔨 Build
```bash
go build ./cmd/todo
```

---

## 📋 Commands

| Action | Command |
|---|---|
| Add a todo | `./todo -add <your text here>` |
| Complete a todo | `./todo -complete=<index>` |
| List all todos | `./todo -list` |
| Delete a todo | `./todo -delete=<index>` |

> Indices start at **1** for better UX. A `.todos.json` file is created automatically on first add.

---

## 📸 Preview

<img width="561" alt="Todo CLI screenshot" src="https://github.com/christsantiris/cli-todo/assets/19711817/769e67ee-ed79-46dd-a188-59202e7bab79">
