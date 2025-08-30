# API Domain Scraper Tool

A simple Go application to fetch domain information from Tatsumi Crew API based on IP address input.

## Features

- ✅ Configurable API key through JSON config file
- ✅ IP address input via command line or interactive prompt
- ✅ HTTP GET request to Tatsumi Crew API
- ✅ JSON response parsing and validation
- ✅ Results saved to output file
- ✅ Error handling and logging
- ✅ Clean and readable code structure

## Prerequisites

- Go 1.13 or higher
- Internet connection for API requests

## Installation

1. Clone this repository:
```bash
git clone https://github.com/TatsumiOfficial/HeartRev.git
cd HeartRev
```

2. Create the configuration file:
```bash
cp config.json.example config.json
```

3. Edit `config.json` with your API key:
```json
{
  "apikey": "YOUR_API_KEY_HERE"
}
```

## Configuration

Create a `config.json` file in the project root directory:

```json
{
  "apikey": "HDNS-02KBDMSGVB3S"
}
```

Replace `HDNS-02KBDMSGVB3S` with your actual API key from Tatsumi Crew.

## Usage

### Method 1: Interactive Input
```bash
go run rev.go
```
The program will prompt you to enter an IP address.

### Method 2: Command Line Argument
```bash
go run rev.go 34.117.176.22/16
```

### Method 3: Build and Run
```bash
go build -o domain-scraper rev.go
./domain-scraper 192.168.1.1/24
```

## API Response Format

The application expects the following JSON response format from the API:

```json
{
  "domains": [
    {"domain": "example.com"},
    {"domain": "test.org"}
  ]
}
```

## Output

### Console Output
```
Melakukan request ke API dengan IP: 34.117.176.22/16
Berhasil mendapatkan 2 domain(s):
- service.ng
- competitions.co
Hasil telah disimpan ke: domains_result.json
```

### File Output
Results are automatically saved to `domains_result.json`:

```json
{
  "domains": [
    {
      "domain": "service.ng"
    },
    {
      "domain": "competitions.co"
    }
  ]
}
```

## Project Structure

```
.
├── main.go              # Main application code
├── config.json          # API configuration file
├── domains_result.json  # Output file (generated)
└── README.md           # This file
```

## Error Handling

The application handles various error scenarios:

- Missing or invalid `config.json` file
- Network connection issues
- Invalid API responses
- HTTP error status codes
- JSON parsing errors
- File I/O errors

## API Endpoint

This tool uses the Tatsumi Crew API:
```
https://api.tatsumi-crew.net/index.php?ip_address={IP}&apikey={API_KEY}
```

## Dependencies

This project uses only Go standard library packages:
- `encoding/json` - JSON encoding/decoding
- `fmt` - Formatted I/O
- `io/ioutil` - I/O utility functions
- `log` - Logging utilities
- `net/http` - HTTP client
- `os` - Operating system interface

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions:

1. Check the [Issues](https://github.com/yourusername/api-domain-scraper/issues) page
2. Create a new issue with detailed information about the problem
3. Include error logs and system information when reporting bugs

## Changelog

### v1.0.0
- Initial release
- Basic API request functionality
- JSON configuration support
- Domain scraping and saving features

---

**Note:** Make sure to keep your API key secure and never commit it to version control. Use environment variables or secure configuration management in production environments.
