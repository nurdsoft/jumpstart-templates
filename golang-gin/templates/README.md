# {{name}}

## Setup

1. Build the binary.
    ```bash
    make releases/{{name}}
    ```
2. Builds a Docker image.
    ```bash
    make docker
    ```
3. Build binary for different OS (e.g., linux, windows, darwin)
    ```bash
    make releases/{{name}}-linux
    ```
4. Remove all build artifacts
    ```bash
    make clean
    ```
