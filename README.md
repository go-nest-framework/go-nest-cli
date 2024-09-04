# GoNest CLI

Welcome to the GoNest CLI! This CLI tool is designed to help developers quickly scaffold and manage projects using the GoNest framework, a Go-based web framework inspired by NestJS. 

## Features

- **Project Creation**: Quickly create new GoNest projects with the `gn create-project` command.
- **Module Generation**: Add new modules to your GoNest project effortlessly.
- **Clean Architecture**: The CLI scaffolds projects and modules that follow clean architecture principles.
- **Easy to Extend**: You can add more commands or customize the CLI for your specific needs.

## Installation

### Prerequisites

- **Go 1.19+**: Make sure you have Go installed on your system. If not, download it from [here](https://golang.org/dl/).
  
### Install the CLI

You can install the GoNest CLI by running the following command:

```bash
go install github.com/go-nest-framework/go-nest-cli/cmd@latest
```

After installation, you can run the CLI commands directly from your terminal using the gn command.

## Usage

### Create a New Project

To create a new GoNest project, run:

```bash
gn new <project-name>
```

This will create a new project folder with the following structure:

```bash
<project-name>/
├── main.go
├── go.mod
├── common/
├── domain/
└── service/
```

### Generate a New Module

To add a new module to your GoNest project, use the following command:

```bash
gn generate <module-name>
```

This will add a module scaffold inside your project.

## Contributing
We welcome contributions to the GoNest CLI! Here’s how you can help:

- ***Fork the repository***: You can fork this repository on GitHub.
- ***Create a feature branch***: Create a new branch for your feature.
- ***Open a pull request***: Once your feature is ready, open a pull request.

### Issues
If you find any bugs or have suggestions for new features, feel free to open an issue.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.


