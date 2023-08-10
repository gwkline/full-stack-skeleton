<div align="center">
<h1 align="center">
<img src="https://upload.wikimedia.org/wikipedia/commons/3/31/Devil_Skull_Icon.svg" width="100" />
<p align="center">
[![Backend CI/CD](https://github.com/gwkline/full-stack-skeleton/actions/workflows/backend.yml/badge.svg)](https://github.com/gwkline/full-stack-skeleton/actions/workflows/backend.yml)
[![Frontend CI/CD](https://github.com/gwkline/full-stack-skeleton/actions/workflows/frontend.yml/badge.svg)](https://github.com/gwkline/full-stack-skeleton/actions/workflows/frontend.yml)
[![Backend Coverage](https://codecov.io/gh/gwkline/full-stack-skeleton/branch/main/graph/badge.svg?token=FQGXXYYJT1)](https://codecov.io/gh/gwkline/full-stack-skeleton)
</p>
<br>full-stack-skeleton
</h1>
<h3>◦ Spin up a bleeding edge web-app in minutes</h3>
<h3>◦ Developed with the software and tools listed below:</h3>

<p align="center">
<img src="https://img.shields.io/badge/GraphQL-E10098.svg?style&logo=GraphQL&logoColor=white" alt="GraphQL" />
<img src="https://img.shields.io/badge/Svelte-FF3E00.svg?style&logo=Svelte&logoColor=white" alt="Svelte" />
<img src="https://img.shields.io/badge/Prettier-F7B93E.svg?style&logo=Prettier&logoColor=black" alt="Prettier" />
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style&logo=Go&logoColor=white" alt="Go" />
<img src="https://img.shields.io/badge/Docker-2496ED.svg?style&logo=Docker&logoColor=white" alt="Docker" />
<img src="https://img.shields.io/badge/GitHub%20Actions-2088FF.svg?style&logo=GitHub-Actions&logoColor=white" alt="GitHub%20Actions" />
<img src="https://img.shields.io/badge/TypeScript-3178C6.svg?style&logo=TypeScript&logoColor=white" alt="TypeScript" />
<img src="https://img.shields.io/badge/postgres-%23316192.svg?style&logo=postgresql&logoColor=white" alt="PostgreSQL" />
<img src="https://img.shields.io/badge/Vite-646CFF.svg?style&logo=Vite&logoColor=white" alt="Vite" />
<img src="https://img.shields.io/badge/ESLint-4B32C3.svg?style&logo=ESLint&logoColor=white" alt="ESLint" />
</p>
<img src="https://img.shields.io/github/languages/top/gwkline/full-stack-skeleton?style&color=5D6D7E" alt="GitHub top language" />
<img src="https://img.shields.io/github/languages/code-size/gwkline/full-stack-skeleton?style&color=5D6D7E" alt="GitHub code size in bytes" />
<img src="https://img.shields.io/github/commit-activity/m/gwkline/full-stack-skeleton?style&color=5D6D7E" alt="GitHub commit activity" />
<img src="https://img.shields.io/github/license/gwkline/full-stack-skeleton?style&color=5D6D7E" alt="GitHub license" />
</div>

---

## 📒 Table of Contents

- [📍 Overview](#-overview)
- [🚀 Getting Started](#-getting-started)
- [🎮 Using Full Stack Skeleton](#-using-full-stack-skeleton)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)

---

## 📍 Overview

The project is a full-stack skeleton application that combines a frontend built with SvelteKit and a backend implemented with Go. It provides a Gin web server with a GraphQL API that integrates with a PostgreSQL database for creating and managing todo items. The core functionalities include user authentication, CRUD operations on todo items, error tracking with Sentry, and Docker support for easy deployment. Overall, the project offers a scalable and easily customizable foundation for building web applications with modern frameworks and technologies.

---

## 🚀 Getting Started

### ✔️ Prerequisites

Before you begin, ensure that you have the following prerequisites installed:

> - `🐳 Docker`
> - `🐿️ Go (including Toolchain)`
> - `🟩 Node.JS`

### 📦 Installation

1. Clone the full-stack-skeleton repository:

```bash
git clone https://github.com/gwkline/full-stack-skeleton
```

2. Change to the project directory:

```bash
cd full-stack-skeleton
```

3. Start the development container:

```bash
docker-compose -f docker-compose.dev.yaml up -d --build
```

4. (Optional) Turn on watch mode

```bash
docker-compose alpha watch
```

To stop all containers:

```bash
docker-compose down
```

---

## 🎮 Using Full Stack Skeleton

By default, the frontend is exposed on [port 3000](http://localhost:3000), and the backend is exposed on [port 8888](http://localhost:8888).
All configuration happens in your `.env` file. Rename the `.env.example` and fill with your own secret variables.

This codebase was built to make starting your next app as quick and easy as possible.
Don't waste your time configuring a brand new environment when you can fork this repo, and start building right away.

### Example new feature workflow:

1. Run this command in your terminal to start the development environment:

```bash
docker-compose -f docker-compose.dev.yaml up -d --build && docker-compose -f docker-compose.dev.yaml alpha watch
```

_Backend_

2. Create/update a service or resource in `./backend`, including any business logic
3. In `./backend/graph`, edit your `schema.graphqls` to define this service/resource using GraphQL
4. Run the following command in `./backend` to regenerate any dependent models, types, or resolvers:

```bash
go run github.com/99designs/gqlgen generate
```

5. Back in `./backend/graph`, update your `schema.resolvers.go` to expose the finished service/resource
6. Write unit test(s) to verify your logic and attempt to prevent future bugs

_Frontend_

7. Create/update the necessary `*.gql` files associated with your pages in `./frontend/src/routes` to use your new service/resource
8. Create/update the necessary `*.svelte`, `*.ts`, or `*.js` files to show off this new GraphQL mutation, query, or field
9. Dogfood on http://localhost:3000

_Finishing Touches_

10. Run this command in your terminal to stop the development environment:

```bash
docker-compose -f docker-compose.dev.yaml down
```

11. Push your code

```bash
# plain ol git
git add .
git commit -m "[YOUR MESSAGE HERE]"
git push origin $(current_branch)

# or you can use Oh My Zsh
gcam "[YOUR MESSAGE HERE]"
ggpush
```

12. Create a PR by navigating to Github - if you have Github CLI, you can use this command in your terminal:

```bash
gh pr create
```

13. Review your PR, make sure it passes CI/CD, then request your peers for review
14. Once approved and merged, let the CI/CD pipeline handle the rest - welcome to production!
15. Keep an eye on [Sentry](https://sentry.io) for any bugs

---

### 🧪 Running Tests

Right now testing is only supported through unit tests on the Go backend. There are two ways to handle this testing currently. If you'd prefer to test outside of the Docker container:

```bash
(cd ./backend && go test)
```

Or you can use the pre-configured testing Docker flow

```bash
docker build -f ./backend/Dockerfile.test -t backend-tester ./backend && docker run --name backend-tester backend-tester && docker rm backend-tester
```

---

## 🤝 Contributing

Contributions are always welcome! Please follow these steps:

1. Fork the project repository. This creates a copy of the project on your account that you can modify without affecting the original project.
2. Clone the forked repository to your local machine using a Git client like Git or GitHub Desktop.
3. Create a new branch with a descriptive name (e.g., `new-feature-branch` or `bugfix-issue-123`).

```bash
git checkout -b new-feature-branch
```

4. Make changes to the project's codebase.
5. Commit your changes to your local branch with a clear commit message that explains the changes you've made.

```bash
git commit -m 'Implemented new feature.'
```

6. Push your changes to your forked repository on GitHub using the following command

```bash
git push origin new-feature-branch
```

7. Create a new pull request to the original project repository. In the pull request, describe the changes you've made and why they're necessary.
   The project maintainers will review your changes and provide feedback or merge them into the main branch.

---

## 📄 License

This project is licensed under the [open-source MIT license](https://github.com/gwkline/full-stack-skeleton/blob/main/LICENSE)

---
