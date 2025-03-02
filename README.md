# Simple Go Network Mapper

A simple and lightweight network mapper written in Go. This tool scans TCP ports on a given host (default ports 1 to 1024) and prints out which ports are open. It is intended for educational purposes.

## Features

- Scans specified TCP ports on a target host.
- Uses concurrent goroutines to speed up the scanning process.
- Configurable timeout for each port connection attempt.

## Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.13 or later recommended).

## Getting Started

1. **Clone the Repository or Download the File**

   You can clone this repository or simply download the `main.go` file.

2. **Build and Run**

   Open your terminal, navigate to the directory containing `main.go`, and run:

   ```bash
   go run main.go <target-host>
