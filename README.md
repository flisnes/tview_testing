# Tview Demo

This is a simple Golang familiarize myself with the use of the [Tview](https://github.com/rivo/tview) module to create a terminal-based user interface (UI).

## Prerequisites

- Go 1.21.1 or later

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/flisnes/tview_testing.git
    cd tview_testing
    ```

2. **Initialize the Go module:**

    ```sh
    go mod init github.com/flisnes/tview_testing
    ```

3. **Install dependencies:**

    ```sh
    go get github.com/rivo/tview
    go get github.com/gdamore/tcell/v2
    ```

## Usage

Run the application:

```sh
go run main.go
```

Interact with the UI:

- Use the arrow keys to navigate through the list
- Press the letter associated with each item to select it
- Press q to quit the application

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgements

- Tview - The library used to create the terminal UI.
- Tcell - The underlying library for handling terminal input and output.
