# Todo App with Go Fiber

A simple Todo application built with the [Go Fiber](https://gofiber.io/) framework. This application supports task creation, deletion, and completion, with all data persisted in a SQLite database. The app is designed with modularity in mind, separating concerns across models, controllers, and routers.

---

## Features

- Add new tasks
- Delete existing tasks
- Mark tasks as completed
- View all tasks in a single HTML page
- Uses Go Fiber's HTML template engine for rendering
- SQLite database for data persistence
- Modular code structure for scalability

---

## Prerequisites

Before running the project, ensure you have the following installed:

- [Go](https://go.dev/) (version 1.18 or later)
- SQLite
- `git` (optional, for cloning the repository)

---

## Getting Started

### Step 1: Clone the Repository
```bash
git clone https://github.com/parshavshah/todo-app-with-go.git
cd todo-app
