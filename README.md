# GoForces

GoForces is a tool to easily set all you need to start a contest from CodeForces as fast as possible. For now it also makes some basic comparison for testing your code.

## Get GoForces

If you want to install GoForces all you need is to have [GoLang](https://golang.org) installed and it is [pretty easy](https://golang.org/doc/install).

Then you need to clone this repository and run the following commands:

```bash
go get && go install
```

## Usage

```bash
Usage:
  goforces [flags]
  goforces [command]

Available Commands:
  help        Help about any command
  scrape      Scrapes a contest or a problem
  test        Takes an executable and tests it with the input files

Flags:
  -h, --help   help for goforces

Use "goforces [command] --help" for more information about a command.
```

## Example

This command will scrape all the contest's problems and save each input and output in text files.

```bash
goforces scrape 1296
```

You will have now this folder structure:

```bash
1296/
├── A
│   ├── input-1.txt
│   └── output-1.txt
├── B
│   ├── input-1.txt
│   └── output-1.txt
├── C
│   ├── input-1.txt
│   └── output-1.txt
...
```

## Built With

* [GoLang](https://golang.org/) - The coding language used
* [Colly](http://go-colly.org/) - The scraper which I hated the documentation
* [Cobra](https://github.com/spf13/cobra) - The tool for creating the CLI

## Contributing

If you want to contribute to the project please read [CONTRIBUTING](CONTRIBUTING.md)

## Authors

* **Vitor Falcão** - *The only one who wants to do this job* - [vitorfhc](https://github.com/vitorfhc)

## License

This project is licensed under the GPL-3.0 - see the [LICENSE.md](LICENSE.md) file for details
