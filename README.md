# mdpdf

**mdpdf** is a command-line tool that converts Markdown files into PDF documents using [Pandoc](https://pandoc.org/) and [Typst](https://typst.app/).

## Prerequisites

Before using mdpdf, ensure the following tools are installed and available in your system's `PATH`:

* [Pandoc](https://pandoc.org/installing.html)
* [Typst](https://typst.app/install)


## Usage

```bash
mdpdf <input.md> <output.pdf>
```

* `<input.md>` — Path to the Markdown file you want to convert.
* `<output.pdf>` — Desired output PDF file path (optional, if empty, pdf has the same name)

### Example

```bash
mdpdf example.md example.pdf
mdpdf example.md
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
