# How to Build Locally

If you want to build the installers on your own machine instead of using GitHub Actions, follow these steps.

## 1. Prerequisites
You need **Go 1.21+** installed.

## 2. Build the Binaries first
Run this in the project root:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o bin/phpvm-linux-amd64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o bin/phpvm-windows-amd64.exe .
```

---

## 3. Build Windows Installer (setup.exe)
**Requirement**: [Inno Setup 6](https://jrsoftware.org/isdl.php) (Windows Only).

1.  Install Inno Setup.
2.  Right-click `build/windows/installer.iss`.
3.  Choose **Compile**.
4.  The output file `phpvm-setup.exe` will appear in `build/windows/Output`.

---

## 4. Build Linux Packages (.deb / .rpm)
**Requirement**: [NFPM](https://nfpm.goreleaser.com/).

1.  Install NFPM:
    ```bash
    go install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest
    ```
2.  Run the packager:
    ```bash
    # Create .deb
    nfpm pkg --packager deb --target bin/ --config build/linux/nfpm.yaml

    # Create .rpm
    nfpm pkg --packager rpm --target bin/ --config build/linux/nfpm.yaml
    ```
3.  The `.deb` and `.rpm` files will be in the `bin/` directory.
