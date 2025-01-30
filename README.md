# GoExploreData

GoExploreData is a simple Command-Line Interface (CLI) tool built in Go that transforms CSV files into dataframes and provides basic statistical insights. This tool allows users to perform quick data analysis directly from the terminal.

## Features

- DataFrame Creation: Convert CSV files into manageable dataframes
- Statistical Analysis: Calculate basic statistics like mean, median, and standard deviation
- Easy to Use: Simple commands for data exploration

## Installation

### Building from Source

1. Clone the repository:
```bash
git clone https://github.com/kylejryan/goexploredata
cd goexploredata
```

2. Build the binary:
```bash
go build -o goexploredata
```

3. Make it available system-wide (Linux/macOS):
```bash
sudo mv goexploredata /usr/local/bin/
```

Or on Windows, move the `goexploredata.exe` to a directory in your PATH.

### Using go install

If you just want to install directly, run:
```bash
go install github.com/kylejryan/goexploredata@latest
```

This will automatically compile and install the binary in your `$GOPATH/bin` directory. Make sure this directory is in your PATH.

### Using go get

Alternatively, you can use:
```bash
go get github.com/kylejryan/goexploredata
```

### Verifying Installation

To verify the installation, run:
```bash
goexploredata --version
```

## Usage

After installation, you can use GoExploreData via the terminal:

```bash
goexploredata [command] [flags] [arguments]
```

### Get Statistics

Generate summary statistics for numeric columns in a CSV file:

```bash
goexploredata stats sample.csv
```

Optional flags:
- `--column, -c`: Specify a column name for focused statistics

### Example

Create a `sample.csv` file with this content:

```csv
name,age,salary
Alice,30,70000
Bob,25,50000
Charlie,35,80000
```

Then run:

```bash
goexploredata stats sample.csv
```

Expected output:

```
Summary Statistics:
Column          Mean       Median     StdDev     Count     
age             30.0000    30.0000    5.0000     3        
salary          66666.6667 70000.0000 15275.2526 3        
```

## Contributing

1. Fork the Repository
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.