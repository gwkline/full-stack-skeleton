[![Backend CI/CD](https://github.com/gwkline/full-stack-skeleton/actions/workflows/backend.yml/badge.svg)](https://github.com/gwkline/full-stack-skeleton/actions/workflows/backend.yml)
[![Frontend CI/CD](https://github.com/gwkline/full-stack-skeleton/actions/workflows/frontend.yml/badge.svg)](https://github.com/gwkline/full-stack-skeleton/actions/workflows/frontend.yml)
[![Backend Coverage](https://codecov.io/gh/gwkline/full-stack-skeleton/branch/main/graph/badge.svg?token=FQGXXYYJT1)](https://codecov.io/gh/gwkline/full-stack-skeleton)

To start all containers (development):
`docker-compose -f docker-compose.dev.yaml up -d --build`

To build just the BE + DB:
`docker-compose -f docker-compose.dev.yaml up -d --build backend`

To stop all containers:
`docker-compose down`

To turn on file watching:
`docker-compose alpha watch`


<div align="center">
<h1 align="center">
<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" />
<br>full-stack-skeleton
</h1>
<h3>◦ Building bridges, powering dreams.</h3>
<h3>◦ Developed with the software and tools listed below.</h3>

<p align="center">
<img src="https://img.shields.io/badge/GNU%20Bash-4EAA25.svg?style&logo=GNU-Bash&logoColor=white" alt="GNU%20Bash" />
<img src="https://img.shields.io/badge/Svelte-FF3E00.svg?style&logo=Svelte&logoColor=white" alt="Svelte" />
<img src="https://img.shields.io/badge/JavaScript-F7DF1E.svg?style&logo=JavaScript&logoColor=black" alt="JavaScript" />
<img src="https://img.shields.io/badge/Prettier-F7B93E.svg?style&logo=Prettier&logoColor=black" alt="Prettier" />
<img src="https://img.shields.io/badge/HTML5-E34F26.svg?style&logo=HTML5&logoColor=white" alt="HTML5" />
<img src="https://img.shields.io/badge/GraphQL-E10098.svg?style&logo=GraphQL&logoColor=white" alt="GraphQL" />
<img src="https://img.shields.io/badge/Pug-A86454.svg?style&logo=Pug&logoColor=white" alt="Pug" />
<img src="https://img.shields.io/badge/Vite-646CFF.svg?style&logo=Vite&logoColor=white" alt="Vite" />

<img src="https://img.shields.io/badge/ESLint-4B32C3.svg?style&logo=ESLint&logoColor=white" alt="ESLint" />
<img src="https://img.shields.io/badge/SemVer-3F4551.svg?style&logo=SemVer&logoColor=white" alt="SemVer" />
<img src="https://img.shields.io/badge/TypeScript-3178C6.svg?style&logo=TypeScript&logoColor=white" alt="TypeScript" />
<img src="https://img.shields.io/badge/Docker-2496ED.svg?style&logo=Docker&logoColor=white" alt="Docker" />
<img src="https://img.shields.io/badge/GitHub%20Actions-2088FF.svg?style&logo=GitHub-Actions&logoColor=white" alt="GitHub%20Actions" />
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style&logo=Go&logoColor=white" alt="Go" />
<img src="https://img.shields.io/badge/JSON-000000.svg?style&logo=JSON&logoColor=white" alt="JSON" />
<img src="https://img.shields.io/badge/Markdown-000000.svg?style&logo=Markdown&logoColor=white" alt="Markdown" />
</p>
<img src="https://img.shields.io/github/languages/top/gwkline/full-stack-skeleton?style&color=5D6D7E" alt="GitHub top language" />
<img src="https://img.shields.io/github/languages/code-size/gwkline/full-stack-skeleton?style&color=5D6D7E" alt="GitHub code size in bytes" />
<img src="https://img.shields.io/github/commit-activity/m/gwkline/full-stack-skeleton?style&color=5D6D7E" alt="GitHub commit activity" />
<img src="https://img.shields.io/github/license/gwkline/full-stack-skeleton?style&color=5D6D7E" alt="GitHub license" />
</div>

---

## 📒 Table of Contents
- [📒 Table of Contents](#-table-of-contents)
- [📍 Overview](#-overview)
- [⚙️ Features](#-features)
- [📂 Project Structure](#project-structure)
- [🧩 Modules](#modules)
- [🚀 Getting Started](#-getting-started)
- [🗺 Roadmap](#-roadmap)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)
- [👏 Acknowledgments](#-acknowledgments)

---


## 📍 Overview

The project is a full-stack skeleton application that combines a frontend built with SvelteKit and a backend implemented with the Go and Gin frameworks. It provides a web server with a GraphQL API that integrates with a PostgreSQL database for creating and managing todo items. The core functionalities include user authentication, CRUD operations on todo items, error tracking with Sentry, and Docker support for easy deployment. Overall, the project offers a scalable and easily customizable foundation for building web applications with modern frameworks and technologies.

---

## ⚙️ Features

| Feature                | Description                           |
| ---------------------- | ------------------------------------- |
| **⚙️ Architecture**     | The codebase follows a client-server architecture, with the frontend built using SvelteKit and the backend built with Go and Gin. The server implements a GraphQL API and integrates with PostgreSQL database and Sentry for error tracking. The frontend and backend are separate components communicating via API calls. Overall, the system is structured as a modern web application following best practices. |
| **📖 Documentation**   | The codebase lacks comprehensive documentation. While comments exist in some files, they are limited and mostly confined to individual functions or interfaces. Developers need to rely on code exploration and self-discovery to understand the project. A more comprehensive documentation effort would greatly benefit the project. |
| **🔗 Dependencies**    | The project has a significant number of dependencies on both the frontend and backend. Notable dependencies include SvelteKit, Gin, PostgreSQL driver, and various utility packages for tasks like validation, error tracking, and tooling. Managing these dependencies requires the use of package managers like PNPM, NPM, or YARN for frontend, and Go modules for the backend. |
| **🧩 Modularity**      | The project is reasonably modular, with clear separation between frontend and backend components. The frontend codebase follows a component-driven approach, making it easier to isolate and reuse functionality. The backend is split into separate files for server logic, GraphQL schema definition, database interaction, middleware, and utilities. This modularity allows for easier maintenance and future extensibility. |
| **✔️ Testing**          | The codebase includes tests for the backend code (server and database), but no specific tests are implemented for the frontend. The backend tests cover various aspects, including environment variable initialization, HTTP request handling, configuration loading, and API endpoint testing. The project could benefit from introducing frontend tests, such as component testing and integration testing using tools like testing libraries or frameworks for SvelteKit. |
| **⚡️ Performance**      | Based on the code analysis, it's not possible to make a definitive assessment of the system's performance. However, the use of Go for the backend and SvelteKit for the frontend indicates a focus on building efficient and performant applications. In terms of resource usage, the frontend and backend build processes may consume significant resources due to the number of dependencies and compilation steps involved. Proper optimization will be required to ensure optimal runtime performance. |
| **🔐 Security**        | The codebase demonstrates some security practices, including validation of email format using the checkmail library and configuring Sentry for error tracking, which can aid in catching potential security-related issues. However, the analysis is based on reviewing the code and does not cover potential security aspects that may exist at the deployment level, such as secure communication (HTTPS) or server hardening. Proper security practices, vulnerability assessments, and other security measures should be taken to ensure a secure application. |
| **🔀 Version Control** | The codebase utilizes Git for version control, as evidenced by the presence of the repository itself. However, the analysis of the codebase does not reveal specific version control strategies or tools being used. To ensure proper version management, it is recommended to follow standard best practices for

---


## 📂 Project Structure


```bash
repo
├── README.md
├── backend
│   ├── Dockerfile
│   ├── Dockerfile.dev
│   ├── Dockerfile.test
│   ├── database
│   │   ├── Dockerfile
│   │   ├── database.go
│   │   └── schema.sql
│   ├── go.mod
│   ├── go.sum
│   ├── gqlgen.yml
│   ├── graph
│   │   ├── generated.go
│   │   ├── model
│   │   │   ├── models_gen.go
│   │   │   └── todo.go
│   │   ├── resolver.go
│   │   ├── schema.graphqls
│   │   └── schema.resolvers.go
│   ├── helpers
│   │   └── helpers.go
│   ├── server.go
│   ├── server_test.go
│   ├── tools.go
│   └── wait.sh
├── docker-compose.dev.yaml
├── docker-compose.prod.yaml
├── docker-compose.test.yaml
└── frontend
    ├── Dockerfile
    ├── Dockerfile.dev
    ├── Dockerfile.test
    ├── README.md
    ├── codegen.ts
    ├── package-lock.json
    ├── package.json
    ├── src
    │   ├── app.d.ts
    │   ├── app.html
    │   ├── hooks.client.ts
    │   ├── hooks.server.ts
    │   ├── lib
    │   │   └── graphql
    │   │       └── generated.ts
    │   └── routes
    │       └── +page.svelte
    ├── static
    │   └── favicon.png
    ├── svelte.config.js
    ├── tsconfig.json
    └── vite.config.ts

12 directories, 41 files
```

---

## 🧩 Modules

<details closed><summary>Frontend</summary>

| File                                                                                                   | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| ---                                                                                                    | ---                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| [svelte.config.js](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/svelte.config.js) | This code snippet is a SvelteKit configuration file that specifies the adapter to be used for the app deployment. It utilizes the'@sveltejs/adapter-node' package for adapter configuration, and the'@sveltejs/kit/vite' package for preprocessors. The'vitePreprocess()' function is applied as a preprocessor, enabling additional processing of the code. The'adapter()' function is used to configure the chosen adapter for the SvelteKit app deployment. Overall, this configuration file sets up the necessary tools for building and deploying a SvelteKit application. |
| [codegen.ts](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/codegen.ts)             | This code snippet defines a configuration object for code generation using a GraphQL codegen library. It specifies the GraphQL schema location, where the generated code will be saved, and the plugins to be used for code generation.                                                                                                                                                                                                                                                                                                                                         |
| [Dockerfile.test](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/Dockerfile.test)   | This code snippet builds a Node.js project in a multi-stage Docker build. The first stage installs dependencies, builds the project, and prunes unnecessary packages. Then, in the second stage, the built artifacts and required dependencies are copied into a production-ready image. The image exposes port 3000 and sets the production environment before starting the application.                                                                                                                                                                                       |
| [Dockerfile](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/Dockerfile)             | This code snippet is a Docker multi-stage build that creates a lightweight Node.js container for running a production-ready application. It copies the necessary files, installs dependencies, builds the application, prunes unnecessary dependencies, and exposes port 3000. The final container runs the built application in production mode.                                                                                                                                                                                                                               |
| [.eslintrc.cjs](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/.eslintrc.cjs)       | This code snippet configures ESLint for a project using TypeScript and Svelte. It sets up recommended rules for ESLint and TypeScript, as well as prettier integration. It also sets the parser options for TypeScript and specifies browser, ES2017, and Node environments. It overrides the parser and parser options specifically for Svelte files, using the svelte-eslint-parser and the TypeScript parser.                                                                                                                                                                |
| [.npmrc](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/.npmrc)                     | The code snippet sets properties for the engine in strict mode and specifies the highest resolution mode.                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| [.prettierignore](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/.prettierignore)   | The code snippet provides a set of rules for file and directory exclusion. It excludes commonly ignored files or directories like ".DS_Store", "node_modules", and "/build". It also ignores certain environment files while including the ".env.example" file. Additionally, it ignores lock files generated by package managers like PNPM, NPM, and YARN.                                                                                                                                                                                                                     |
| [Dockerfile.dev](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/Dockerfile.dev)     | The code snippet sets up a Node.js environment, installs dependencies, copies code files, exposes port 3000, sets environment variables to "development mode," and executes a development run command.                                                                                                                                                                                                                                                                                                                                                                          |
| [.eslintignore](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/.eslintignore)       | The code snippet contains a set of rules for files and directories that should be ignored when managing the project with PNPM, NPM, or YARN package managers. It includes commonly ignored files like ".DS_Store", "node_modules", "build", and various lock files.                                                                                                                                                                                                                                                                                                             |
| [vite.config.ts](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/vite.config.ts)     | This code snippet sets up a Svelte Kit application with Vite. It includes plugins for Sentry, which helps monitor and manage errors in the frontend. The configuration allows for the uploading of source maps to Sentry, enabling easy debugging. The dotenv library is used to load environment variables. The application runs on port 3000.                                                                                                                                                                                                                                 |

</details>

<details closed><summary>Src</summary>

| File                                                                                                     | Summary                                                                                                                                                                                                                                                                                                                                  |
| ---                                                                                                      | ---                                                                                                                                                                                                                                                                                                                                      |
| [app.d.ts](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/src/app.d.ts)               | The code snippet includes commented interfaces for Error, Locals, PageData, and Platform inside the global App namespace. It then exports an empty object to indicate the end of the code.                                                                                                                                               |
| [hooks.client.ts](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/src/hooks.client.ts) | The code snippet initializes and configures the Sentry error reporting tool for a SvelteKit application. It sets up the Sentry instance with the necessary options like the DSN, sample rates for tracing and session replays, and enables the Replay integration. It also provides a custom error handler to handle errors with Sentry. |
| [app.html](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/src/app.html)               | This code snippet represents a basic HTML template file. It includes the necessary metadata, viewport settings, and references to assets. Additionally, it has placeholders for SvelteKit specific functionality related to head and body components. Overall, it sets up the structure for a web page.                                  |
| [hooks.server.ts](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/src/hooks.server.ts) | The code snippet initializes and configures Sentry for error handling and monitoring in a SvelteKit app. It sets up the Sentry client with a provided DSN and a traces sample rate of 1.0. It also defines a custom error handling sequence and handler that integrates with Sentry.                                                     |

</details>

<details closed><summary>Graphql</summary>

| File                                                                                                           | Summary                                                                                                                                                                                                                                                                      |
| ---                                                                                                            | ---                                                                                                                                                                                                                                                                          |
| [generated.ts](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/src/lib/graphql/generated.ts) | The provided code snippet is a collection of type definitions and utility functions related to a GraphQL schema. It includes definitions for mutations, queries, and custom scalars. There are also functions related to initializing the KitQL client and resetting caches. |

</details>

<details closed><summary>Routes</summary>

| File                                                                                                      | Summary                                                                                                                                                                                 |
| ---                                                                                                       | ---                                                                                                                                                                                     |
| [+page.svelte](https://github.com/gwkline/full-stack-skeleton/blob/main/frontend/src/routes/+page.svelte) | The provided code snippet is a basic HTML template that displays a welcome message and a link. The link directs the user to the documentation page of the "porg" application framework. |

</details>

<details closed><summary>Backend</summary>

| File                                                                                                | Summary                                                                                                                                                                                                                                                                                                                                                                                                      |
| ---                                                                                                 | ---                                                                                                                                                                                                                                                                                                                                                                                                          |
| [go.mod](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/go.mod)                   | The code snippet includes a list of Go packages required by the module. These packages are necessary for implementing various functionalities such as GraphQL support, email validation, CORS handling, database connectivity, error logging, and MIME type identification, among others.                                                                                                                    |
| [server.go](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/server.go)             | This code snippet sets up a web server using the Gin framework and implements a GraphQL API. It also integrates Sentry for error tracking and PostgreSQL for database support. CORS middleware is applied based on the environment. The server listens on a specified port.                                                                                                                                  |
| [Dockerfile.test](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/Dockerfile.test) | The provided code snippet is for building a Docker image for a Go application. It consists of two stages-the build stage and the final stage. In the build stage, Go dependencies are downloaded and the source code is copied. In the final stage, the source code and dependencies are copied from the build stage. Finally, it runs a command to test the application and generate a coverage report.     |
| [Dockerfile](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/Dockerfile)           | This code snippet is a Dockerfile for building and running a Go application. It has two stages-building the application and running it. The builder stage downloads dependencies, copies the code, and builds the application. The final stage sets up the runtime environment, copies the binary and a script, makes the script executable, exposes a port, and executes the script to run the application. |
| [tools.go](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/tools.go)               | The code snippet is for managing third-party tools using Go modules. It imports the gqlgen package, which is a popular tool for creating GraphQL servers, but doesn't use it directly. This package is specifically used for managing development, build, and tooling dependencies.                                                                                                                          |
| [Dockerfile.dev](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/Dockerfile.dev)   | The provided code snippet sets up a Docker container based on Alpine Linux with necessary dependencies. It copies the application code, downloads Go modules, and exposes port 8888. It then runs a shell script to wait for the database to be ready before executing the server application.                                                                                                               |
| [wait.sh](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/wait.sh)                 | This code snippet is a bash script used to test the availability of a TCP host and port. It includes options for setting a timeout, strict mode, and executing additional commands after the test completes. The script uses netcat or built-in bash functionality to check the availability of the host and port.                                                                                           |
| [server_test.go](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/server_test.go)   | The code snippet includes multiple test functions that cover different aspects of the application, such as testing environment variable retrieval and initialization, handling HTTP requests, loading configuration, and testing specific API endpoints.                                                                                                                                                     |

</details>

<details closed><summary>Database</summary>

| File                                                                                                 | Summary                                                                                                                                                                                                       |
| ---                                                                                                  | ---                                                                                                                                                                                                           |
| [schema.sql](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/database/schema.sql)   | This code snippet creates a PostgreSQL database named "postgresql_main" and a table named "contacts" with columns for contact ID, first name, last name, email, phone number, and creation/update timestamps. |
| [Dockerfile](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/database/Dockerfile)   | This code snippet sets up a PostgreSQL Docker container by adding an SQL script called "schema.sql" to the "docker-entrypoint-initdb.d" directory.                                                            |
| [database.go](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/database/database.go) | This code snippet initializes a connection with a PostgreSQL database and provides functions for inserting, retrieving, updating, and deleting contact data from the database.                                |

</details>

<details closed><summary>Graph</summary>

| File                                                                                                              | Summary                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ---                                                                                                               | ---                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| [schema.graphqls](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/graph/schema.graphqls)         | This GraphQL schema defines two types, Todo and User. Todo has fields for id, text, done, and user, while User has fields for id and name. The Query type has a todos field which returns a list of Todo objects. The Mutation type has a createTodo field which receives a NewTodo input and creates a new Todo object. The NewTodo input has fields for text and userId. This schema allows querying for todos and creating new todos. |
| [generated.go](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/graph/generated.go)               | Prompt exceeds max token limit: 12339.                                                                                                                                                                                                                                                                                                                                                                                                   |
| [resolver.go](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/graph/resolver.go)                 | The provided code snippet contains a custom Resolver struct that holds a slice of Todo pointers from the model package. It serves as a dependency injection and is used to manage and retrieve Todos in the application.                                                                                                                                                                                                                 |
| [schema.resolvers.go](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/graph/schema.resolvers.go) | This code snippet is a part of a GraphQL resolver implementation for a todo app. It includes functionality for creating a new todo, retrieving todos, and resolving the user associated with a todo. The code utilizes some generated models and functions for random number generation and string formatting.                                                                                                                           |

</details>

<details closed><summary>Model</summary>

| File                                                                                                        | Summary                                                                                                                                                                                                                             |
| ---                                                                                                         | ---                                                                                                                                                                                                                                 |
| [models_gen.go](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/graph/model/models_gen.go) | The code snippet generates models for a GraphQL schema. It defines a data structure for a new todo item, containing text and a user ID. It also defines a user model with an ID and a name. The structs have tags to map JSON keys. |
| [todo.go](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/graph/model/todo.go)             | The code provides a "Todo" structure that represents a to-do item. It includes fields for ID, text, completion status, userID, and a pointer to a User structure.                                                                   |

</details>

<details closed><summary>Helpers</summary>

| File                                                                                              | Summary                                                                                                                                                                                                                                                                                               |
| ---                                                                                               | ---                                                                                                                                                                                                                                                                                                   |
| [helpers.go](https://github.com/gwkline/full-stack-skeleton/blob/main/backend/helpers/helpers.go) | This code snippet in the "helpers" package provides a function called "ValidateEmail" which takes an email string as input. It uses the "checkmail" library to validate the email format and the email host. It returns an error if either validation fails, or it returns nil if the email is valid. |

</details>

---

## 🚀 Getting Started

### ✔️ Prerequisites

Before you begin, ensure that you have the following prerequisites installed:
> - `ℹ️ Requirement 1`
> - `ℹ️ Requirement 2`
> - `ℹ️ ...`

### 📦 Installation

1. Clone the full-stack-skeleton repository:
```sh
git clone https://github.com/gwkline/full-stack-skeleton
```

2. Change to the project directory:
```sh
cd full-stack-skeleton
```

3. Install the dependencies:
```sh
go build -o myapp
```

### 🎮 Using full-stack-skeleton

```sh
./myapp
```

### 🧪 Running Tests
```sh
go test
```

---


## 🗺 Roadmap

> - [X] `ℹ️  Task 1: Implement X`
> - [ ] `ℹ️  Task 2: Refactor Y`
> - [ ] `ℹ️ ...`


---

## 🤝 Contributing

Contributions are always welcome! Please follow these steps:
1. Fork the project repository. This creates a copy of the project on your account that you can modify without affecting the original project.
2. Clone the forked repository to your local machine using a Git client like Git or GitHub Desktop.
3. Create a new branch with a descriptive name (e.g., `new-feature-branch` or `bugfix-issue-123`).
```sh
git checkout -b new-feature-branch
```
4. Make changes to the project's codebase.
5. Commit your changes to your local branch with a clear commit message that explains the changes you've made.
```sh
git commit -m 'Implemented new feature.'
```
6. Push your changes to your forked repository on GitHub using the following command
```sh
git push origin new-feature-branch
```
7. Create a new pull request to the original project repository. In the pull request, describe the changes you've made and why they're necessary.
The project maintainers will review your changes and provide feedback or merge them into the main branch.

---

## 📄 License

This project is licensed under the `ℹ️  INSERT-LICENSE-TYPE` License. See the [LICENSE](https://docs.github.com/en/communities/setting-up-your-project-for-healthy-contributions/adding-a-license-to-a-repository) file for additional info.

---

## 👏 Acknowledgments

> - `ℹ️  List any resources, contributors, inspiration, etc.`

---
