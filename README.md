<div align="center">
<h1 align="center">
<img src="https://upload.wikimedia.org/wikipedia/commons/3/31/Devil_Skull_Icon.svg" width="100" />
<br>full-stack-skeleton
</h1>
<h3>◦ Spin up a bleeding edge web-app in minutes</h3>
<h3>◦ Developed with the software and tools listed below:</h3>
<p align="center">
<img src="https://img.shields.io/badge/TypeScript-3178C6.svg?style&logo=TypeScript&logoColor=white" alt="TypeScript" />
<img src="https://img.shields.io/badge/React-61DAFB.svg?style&logo=React&logoColor=white" alt="React" />
<img src="https://img.shields.io/badge/Next.JS-000000.svg?style&logo=Next.JS&logoColor=white" alt="Next" />
<img src="https://img.shields.io/badge/Bun-000000.svg?style&logo=Bun&logoColor=white" alt="Bun" />
<img src="https://img.shields.io/badge/Biome-60A5FA.svg?style&logo=Biome&logoColor=white" alt="Biome" />
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style&logo=Go&logoColor=white" alt="Go" />
<img src="https://img.shields.io/badge/GraphQL-E10098.svg?style&logo=GraphQL&logoColor=white" alt="GraphQL" />
<img src="https://img.shields.io/badge/postgres-%23316192.svg?style&logo=postgresql&logoColor=white" alt="PostgreSQL" />
<img src="https://img.shields.io/badge/Docker-2496ED.svg?style&logo=Docker&logoColor=white" alt="Docker" />
<img src="https://img.shields.io/badge/GitHub%20Actions-2088FF.svg?style&logo=GitHub-Actions&logoColor=white" alt="GitHub%20Actions" />
</p>
<p align="center">
<img src="https://github.com/gwkline/full-stack-skeleton/actions/workflows/backend.yml/badge.svg" alt="Backend CI/CD">
<img src="https://img.shields.io/badge/github/deployments/gwkline/full-stack-skeleton/production?label=Frontend CI/CD&logo=vercel" alt="vercel" />
<img src="https://codecov.io/gh/gwkline/full-stack-skeleton/graph/badge.svg?token=5LE26Z9EQV" alt="Backend Coverage">
</p>
</div>

---

## 📒 Table of Contents

- [📍 Overview](#-overview)
- [🚀 Getting Started](#-getting-started)
- [🎮 Workflow](#-workflow)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)

---

## 📍 Overview

The project is a full-stack application that combines a frontend built with React + NextJS and a backend built with Go. The frontend utilizes the Next's App Router, with Tanstack Query handling data fetching and ShadcnUI as the UI component library. The backend uses Gin to serve a GraphQL API, which fetches data from a PostgreSQL database. The project is fully containerized with Docker and uses GitHub Actions for CI/CD.

---

## 🚀 Getting Started

### ✔️ Prerequisites

Before you begin, ensure that you have the following prerequisites installed:

> - `🐳 Docker`
> - `🐿️ Go`
> - `🟩 Node.JS`

### 📦 Installation

1. Clone the email-app repository:

```bash
git clone https://github.com/gwkline/full-stack-skeleton
```

2. Change to the project directory:

```bash
cd email-app
```

3. Start the development containers:

```bash
docker compose watch
```

To stop all containers:

```bash
docker compose down
```

---

## ✍️ Workflow

By default, the frontend is exposed on [port 3000](http://localhost:3000), and the backend is exposed on [port 8888](http://localhost:8888).
All configuration happens in the root `.env` file. Rename the `.env.example` and fill with your own secret variables.

### Example new feature workflow:

1. Pull the latest changes from the `main` branch:
```bash
#---------- git cli --------------
git checkout main
git pull origin main
#---------- oh-my-zsh --------------
gco main
ggpull 
```

2. Create a new branch with a descriptive name:

```bash
#---------- git cli --------------
git checkout -b "new-feature-branch"
#---------- oh-my-zsh --------------
gco -b "bugfix-issue-123"
```


3. Run this command in your terminal to start the development environment:

```bash
docker compose watch 
```

_Backend_

1. Create/update a service or resource in `./backend/pkg`, including any business logic
2. To expose this logic, navigate to `./backend/graph` and edit or create GraphQL types in`schema/*.gql`
3. Run the following command in the root directory to regenerate any dependent models, types, or resolvers:

```bash
(cd backend && go run github.com/99designs/gqlgen generate)
```

4. Back in `./backend/graph`, update the corresponding `resolvers/*_resolver.go` file(s) to expose the finished service/resource
5. Write unit test(s) to verify your logic and attempt to prevent future bugs

_Frontend_

1. Navigate to `./frontend/lib`, where you can create/update any needed GraphQL request functions.
2. We use a codegen to ensure all of the types match up with our GraphQL schema. Run the following command to generate the types:

```bash
(cd ./frontend && bun codegen)
```

3. Make sure the request you've written is properly typed, and then you can navigate to `./frontend/components` and/or `./frontend/app` to implement this feature
4. Create/update the necessary `*.tsx` or `*.ts` files to show off this new GraphQL mutation, query, or field
5. Dogfood on http://localhost:3000

_Finishing Touches_

1. Run this command in your terminal to stop the development environment:

```bash
docker compose down
```

2. Push your code and make a PR (see [🤝 Contributing](#-contributing))
3. Review your PR, make sure it passes CI/CD, then request your peers for review
4. Once approved and merged, let the CI/CD pipeline handle the rest - welcome to production!
5. Keep an eye on [Sentry](https://sentry.io) for any bugs

---

### 🧪 Running Tests

Right now testing is supported through unit tests on the Go backend:

```bash
(cd ./backend &&  go test -cover ./...)
```

To generate function mocks, we use [mockery](https://vektra.github.io/mockery/latest/installation/):
```bash
(cd backend && mockery)
```

---

## 🤝 Contributing

Contributions are always welcome! Please follow these steps:

1. Clone the project repository to your local machine using a Git client like Git or GitHub Desktop.
2. Create a new branch with a descriptive name (e.g., `new-feature-branch` or `bugfix-issue-123`).

```bash
#---------- git cli --------------
git checkout -b new-feature-branch

#---------- oh-my-zsh --------------\
gco -b new-feature-branch
```

4. Make changes to the project's codebase.
5. Commit your changes to your local branch with a clear commit message that explains the changes you've made.

```bash
#---------- git cli --------------
git add .
git commit -m "Implemented new feature."

#---------- oh-my-zsh --------------\
gcam "Implemented new feature."
```

6. Push your changes to the repository on GitHub using the following command

```bash
#---------- git cli --------------
git push origin new-feature-branch

#---------- oh-my-zsh --------------\
ggpush
```

7. Create a PR by navigating to Github - if you have Github CLI, you can use this command in your terminal:

```bash
gh pr create
```

In the pull request, describe the changes you've made and why they're necessary.
The project maintainers will review your changes and provide feedback or merge them into the main branch.

---

## 📄 License

This project is licensed under the [open-source MIT license](https://github.com/gwkline/full-stack-skeleton/blob/main/LICENSE)

---
