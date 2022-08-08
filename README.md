# ttynvt-runner

The ttynvt runner is used to observe _ttynvt._tcp mdns services of io4edge devices. It starts a new ttynvt instance when a new service shows up and terminates the corresponding instance again when the service disappears. The ttynvt instance creates for the _ttynvt._tcp mdns service a virtual tty `/dev/tty<mdns-instance-name>`.

For more information about ttynvt see https://gitlab.com/ci4rail/ttynvt.

# Usage
```
$ sudo ./ttynvt-runner [-m <major-number>] <ttynvt-program-path>
```

# Build

Example: Build for arm64:

```bash
cd cmd/ttynvt-runner
GOOS=linux GOARCH=arm64 go build -o ../bin/ttynvt-runner .
```
