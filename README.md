# Learning Go

This repository contains a collection of Go programming exercises, small projects, and workshops designed to help learn and practice Go (Golang).

## Project Structure

The repository is organized into several directories, each focusing on different aspects of Go programming:

- `arrays-practice`: Exercises for working with arrays in Go
- `arrays-slices-maps`: Examples and practice with arrays, slices, and maps
- `bedrock-go`: A small Go project (details to be added)
- `calculator`: A more complex project demonstrating various Go concepts
- `control-structures`: Examples of control structures in Go
- `custom-types`: Exercises for creating and using custom types
- `firstgo`: Your first Go program
- `functions-divedeep`: In-depth exploration of Go functions
- `go-cdk-workshop`: An AWS CDK workshop using Go
- `goppgm`: A Go project (details to be added)
- `pointers`: Examples and exercises with Go pointers
- `profit-calculator`: A simple profit calculation program
- `structs`: Examples and exercises with Go structs
- `structs-practice`: More practice with Go structs

## Setup

To run the projects in this repository, you need to have Go installed on your system. You can download and install Go from the [official Go website](https://golang.org/dl/).

For the `go-cdk-workshop` project, you'll also need to have the AWS CDK installed. You can find installation instructions in the [AWS CDK documentation](https://docs.aws.amazon.com/cdk/latest/guide/getting_started.html).

## Usage

To run any of the Go programs:

1. Navigate to the directory containing the program you want to run.
2. Use the `go run` command followed by the name of the Go file. For example:

   ```
   cd firstgo
   go run app.go
   ```

For the `go-cdk-workshop` project, refer to the `cdk.json` file in that directory for specific run commands.

## Main Components

### Calculator

The `calculator` directory contains a more complex project with multiple components:

- `cmdmanager`: Manages command-line interactions
- `conversion`: Handles data type conversions
- `filemanager`: Manages file operations
- `IOmanager`: Handles input/output operations
- `prices`: Deals with price calculations

To run the calculator, navigate to the `calculator` directory and run:

```
go run main.go
```

### Go CDK Workshop

The `go-cdk-workshop` directory contains an AWS CDK project written in Go. It includes Lambda functions and CDK stacks for deploying a simple web application with a hit counter.

To run this project, make sure you have the AWS CDK installed and configured, then navigate to the `go-cdk-workshop` directory and run:

```
cdk deploy
```