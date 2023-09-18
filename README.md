# Who Am I Application

The **Who Am I Application** is a command-line tool for searching information about IPs and domains. You can use it to find out details about a specific host, whether it's an IP address or a domain name.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Contributing](#contributing)
- [License](#license)

## Installation

To get started with the **Who Am I Application**, follow these steps:

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/whoami.git
   ```

2. **Build the Application**:

   ```bash
   cd whoami
   go build
   ```

3. **Run the Application**:

   ```bash
   ./whoami
   ```

Now you're ready to use the application.

## Usage

The **Who Am I Application** allows you to search for information about IPs and domains. Here's how you can use it:

- To search for information about an IP address, use the following command:

  ```bash
  ./whoami ips <host>
  ```

  Replace `<host>` with the IP address you want to look up.

- To search for information about a domain, use the following command:

  ```bash
  ./whoami domains <ip_address>
  ```

  Replace `<ip_address>` with the IP address associated with the domain you want to look up.

## Commands

The **Who Am I Application** provides the following commands:

- `ips` (alias: `i`): Search for information about an IP address.

- `domains` (alias: `d`): Search for information about a domain.

### Examples

- To search for information about an IP address:

  ```bash
  ./whoami ips 8.8.8.8
  ```

- To search for information about a domain:

  ```bash
  ./whoami domains google.com
  ```

## Contributing

Contributions to the **Who Am I Application** are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository on GitHub.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with clear commit messages.
4. Push your changes to your fork.
5. Submit a pull request to the main repository.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Feel free to customize this README to provide more specific information about your application or any additional details you'd like to include.