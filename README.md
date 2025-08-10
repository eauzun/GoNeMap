# GoNeMap — Lightweight TCP Port Scanner in Go

GoNeMap is a **minimalist, high-performance TCP port scanner** written in Go.  
It follows a lightweight approach while keeping the code **fully extensible** for network professionals, penetration testers, and learners.

## 🔹 Features
- Scans the most common TCP ports by default.
- **Easily extendable** — just add new ports to the `commonPorts` map in the source code.
- Adjustable connection timeout (in milliseconds).
- Uses **TCP Connect() scanning** (`net.DialTimeout`) to detect open ports.
- Minimal dependencies — only standard Go libraries.

## 🔹 How It Works
GoNeMap uses the **TCP Connect scan** method:
1. Attempts to establish a TCP connection to each target port.
2. If the connection is successful → the port is reported as **OPEN**.
3. If it fails or times out → the port is reported as **closed**.

This is similar in principle to Nmap's `-sT` scan mode, but implemented from scratch.

## 🔹 Installation & Usage

### 1. Clone the Repository
```bash
git clone https://github.com/eauzun/GoNeMap.git
cd GoNeMap
```

### 2. Run Without Building
```bash
go run main.go -host <target-host> -timeout <ms>
```

Example:
```bash
go run main.go -host 192.168.1.1 -timeout 500
```

### 3. Build & Run
```bash
go build -o gonemap main.go
./gonemap -host example.com -timeout 300
```

## 🔹 Adding Custom Ports
All ports to be scanned are stored in the `commonPorts` map inside `main.go`:
```go
var commonPorts = map[int]string{
    21: "FTP",
    22: "SSH",
    30: "Telnet-Alt", // Example custom port
    80: "HTTP",
    443: "HTTPS",
}
```
Simply add your desired port and service name to this list — the program will automatically include it in the scan.

## 🔹 Example Output
```
Scanning 192.168.1.1
-------------------------
[+] Port 22 (SSH): OPEN
[-] Port 23 (Telnet): closed
[+] Port 80 (HTTP): OPEN
Scan Process Completed
```

## 📜 License
MIT License — feel free to use and modify.

