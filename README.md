
# `konvertu` - unit converter webapp

## Overview
This project is a simple web application designed to convert between different 
physical units. It allows users to input a value, select the source and target 
units, and instantly view the converted result. The application supports a wide 
range of unit categories including length, weight, volume, area, temperature, 
and more.

## Features
- User-friendly web interface.
- Conversion between multiple unit categories:
  - **Length**: millimeter, centimeter, meter, kilometer, inch, foot, yard, mile.
  - **Weight**: milligram, gram, kilogram, ounce, pound.
  - **Temperature**: Celsius, Fahrenheit, Kelvin.
- Form-based input: users enter a value and select units.
- Server-side processing: the form is submitted to the server, which performs 
the conversion and returns the result.
- Results are displayed directly on the web page.

## Getting started

- Clone the repository and navigate into the project directory:

```bash
git clone https://github.com/Mirsait/konvertu.git
cd konvertu
```

- Build the project

```bash
go build -o konvertu
```

Start the server and open the application in your browser:

```bash
./konvertu
```

## License

[MIT License](LICENSE) — feel free to use, modify, and distribute.
