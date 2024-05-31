### Img2QR

#### Overview

The **Img2QR** is a command-line tool designed to convert image files into QR codes. This utility takes image files from the current directory, converts them into base64-encoded strings, and then encodes these base64 strings into QR codes. The resulting QR code images are saved in the same directory, providing a convenient way to encode and distribute image data via QR codes.

#### Features

- **Directory Scanning**: Automatically scans a specified directory for image files with extensions `.jpg`, `.jpeg`, and `.png`.
- **Base64 Encoding**: Converts each image file into a base64-encoded string.
- **QR Code Generation**: Encodes the base64 string into a QR code.
- **Customizable Output**: Allows customization of QR code dimensions (width and height) via command-line flags.
- **Batch Processing**: Processes all supported image files in the specified directory.

#### Usage

To use the utility, compile the Go program and run the executable with the desired flags.

##### Compiling the Program

```sh
go build -o img2qr main.go
```

##### Running the Executable

```sh
./img2qr -width 512 -height 512
```

##### Command-Line Flags

- `-width`: Specifies the width of the QR code (default: 256 pixels).
- `-height`: Specifies the height of the QR code (default: 256 pixels).

#### Example

```sh
./img2qr -width 512 -height 512
```

This command will scan the `./images` directory for image files, convert each image to a base64 string, generate a QR code of size 512x512 pixels for each base64 string, and save the QR codes in the same directory.
