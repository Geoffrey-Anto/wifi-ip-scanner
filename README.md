# WiFi IP Scanner

WiFi IP Scanner is a GoLang application designed to scan local network IP addresses to identify active hosts. It utilizes Go routines for concurrency, enabling faster ping responses and efficient network scanning.

## Features

- **Fast Scanning**: Utilizes Go routines for concurrent pinging, enabling quicker scanning of local IP addresses.
- **Customizable Parameters**: Allows customization of ping parameters such as number of retries and timeout duration.
- **Cross-Platform**: Supports building for multiple platforms including Linux, macOS, and Windows.
- **Easy Installation**: Provides build scripts for easy compilation and installation.

## Installation

To install the WiFi IP Scanner, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/geoffrey-anto/wifi-ip-scanner.git
    ```

2. Navigate to the project directory:

    ```bash
    cd wifi-ip-scanner
    ```

3. Build the executable (suffix build with -linux, -macos, -windows or leave it as such for builds for multiple platforms):

    ```bash
    make build
    ```

4. Make the executable file executable:

    ```bash
    make make-executable
    ```

5. Optionally, install the executable to your system:

    ```bash
    sudo make install
    ```

## Usage

After installation, you can run the WiFi IP Scanner by executing the generated executable file. Here are some example commands:

- Run the scanner:

    ```bash
    ./bin/scan
    ```

## Contributions

Contributions to the WiFi IP Scanner project are welcome! If you have any ideas for improvements or find any issues, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
