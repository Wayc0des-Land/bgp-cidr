# ASN Lookup Tool

This tool retrieves ASNs (Autonomous System Numbers) and associated CIDR blocks from the BGP routing information site [bgp.he.net](https://bgp.he.net/).

This project is a Go adaptation of the Python script originally created by [defparam](https://gist.githubusercontent.com/defparam/ab1e6b249ceaed51ae60ff6479b67869/raw/b8e3051890963ec1e90eddc48b8aa92cb17a5ddc/cidr_info.py).

## Features

- **lookup_asn_from_org**: Looks up ASNs associated with a specified organization.
- **get_cidrs_from_asn**: Retrieves CIDR blocks associated with a specific ASN.

## Requirements

- Go (Golang)
- `github.com/PuerkitoBio/goquery` package for HTML parsing.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Wayc0des-Land/bgp-cidr.git
   ```

2. Install dependencies:

   ```bash
   go get github.com/PuerkitoBio/goquery
   ```

3. Build the executable:

   ```bash
   go build
   ```

## Usage

```bash
./bgp-cidr <organization>
```

Replace `<organization>` with the name of the organization you want to look up.

## Example

```bash
./bgp-cidr Google
```

This will fetch ASNs associated with "Google" and their corresponding CIDR blocks.

## Contributing

Contributions are welcome! Please feel free to fork the repository and submit pull requests.

## Credits

- Original Python script by [defparam](https://gist.githubusercontent.com/defparam/ab1e6b249ceaed51ae60ff6479b67869/raw/b8e3051890963ec1e90eddc48b8aa92cb17a5ddc/cidr_info.py).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### Notes:
- Replace `<organization>` with the actual organization name you want to look up.
- Ensure that you have Go installed and `goquery` package is fetched before building the executable.
- Customize the `LICENSE` section with the appropriate license information for your project.

Save this content into a file named `README.md` in the root directory of your GitHub repository. This README now includes a credits section acknowledging the original author of the Python script. Adjust URLs, commands, and descriptions as needed to fit your specific project details.
