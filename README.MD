# SEA Salon

## Overview

SEA Salon is a web application built with Go and HTML. This project is for Compfest SEA Academy, this project is the first level version.

## Features

- Home Page
- Services Overview
- Reservation System
- User Authentication

## Prerequisites

- [Go](https://golang.org/doc/install) 1.22.3 or later
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/onlydarel/sea-salon.git
    cd sea-salon
    ```

2. Set up the Go module:

   Ensure the `go.mod` file is correctly set up. If needed, you can initialize or update it:

    ```sh
    go mod init github.com/onlydarel/sea-salon
    go mod tidy
    ```

## Running the Application

1. Build and run the server:

    ```sh
    go run main.go
    ```

2. Open your web browser and navigate to `http://localhost:8080`.

## Project Structure

- `main.go`: The main entry point for the application.
- `pkg/handlers/`: Contains the handler functions for various routes.
- `assets/`: Contains static assets like images and JavaScript files.
- `styles/`: Contains CSS files for styling the web pages.

## Usage

- There is no usage yet, because the first level task is only to build the homepage.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature-name`).
3. Make your changes and commit them (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature-name`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

- **Darel Pnadey** - [GitHub](https://github.com/onlydarel)

