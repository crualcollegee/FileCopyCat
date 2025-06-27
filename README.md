# FileCopyCat Project

This project is a Go-based file management system that allows users to copy files with specific extensions from one directory to another. It uses the `github.com/spf13/cobra` package for command-line interface (CLI) management. This README also provides a script that checks dependencies, compiles the project, and allows interactive command execution.

## Prerequisites

Before running the project, ensure you have the following installed:

- Go (version 1.16 or later)
- Git
- Bash (for running the provided script)

### Install Go
You can download Go from the official website: [Go Downloads](https://golang.org/dl/).

### Install Git
Git can be installed from the official website: [Git Downloads](https://git-scm.com/downloads).

## Installation

### Step 1: Clone the repository

Clone the repository to your local machine:

```bash
git clone https://github.com/crualcollegee/FileCopyCat.git
cd FileCopyCat
```

Make the run.sh script executable:
```bash
chmod +x run.sh
```
Execute the script to compile the project and start the interactive command line interface:
```bash
./run.sh
```

The script will:

Check if the github.com/spf13/cobra package is installed. If not, it will install the package.

Compile the Go project.

Allow you to run commands interactively, such as copying files with specific extensions. The script will keep running until you type exit.

Example Command
When prompted by the script, enter a command in the format:

```bash
ExtCopy /path/to/source /path/to/target png
```
