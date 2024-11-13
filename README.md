# Embed-Trivy

The go program embeds the trivy binary into the go program. And creates a "trivy-test-isolation" binary and runs the --version of trivy CLI. Commands:

Install trivy:
```bash
curl -L -o trivy.tar.gz https://github.com/aquasecurity/trivy/releases/download/v0.52.0/trivy_0.52.0_Linux-64bit.tar.gz \
    && tar xvzf trivy.tar.gz \
    && chmod +x trivy && cp trivy /usr/local/bin/trivy
```

Copy the trivy binary to the current directory's asset:
```bash
mkdir assets
sudo cp /usr/local/bin/trivy ./assets
```
(if trivy is not in the /usr/local/bin, please change the path accordingly. use `which trivy` to find the path)

Build the go program which is embedding the trivy binary:
```bash
go build .
```

Run the go program:
```bash
./embed-trivy
```