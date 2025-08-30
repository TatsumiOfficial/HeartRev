package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Config struct {
	APIKey string `json:"apikey"`
}

type Domain struct {
	Domain string `json:"domain"`
}

type Response struct {
	Domains []Domain `json:"domains"`
}

func loadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, fmt.Errorf("gagal membuka config.json: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("gagal decode config.json: %v", err)
	}

	return &config, nil
}

func makeAPIRequest(ipAddress, apiKey string) (*Response, error) {
	url := fmt.Sprintf("https://api.tatsumi-crew.net/index.php?ip_address=%s&apikey=%s", ipAddress, apiKey)
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("gagal melakukan request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error: status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca response body: %v", err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("gagal decode JSON response: %v", err)
	}

	return &response, nil
}

func saveDomains(response *Response, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("gagal membuat file %s: %v", filename, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(response); err != nil {
		return fmt.Errorf("gagal encode JSON ke file: %v", err)
	}

	return nil
}

func main() {

	var ipAddress string
	if len(os.Args) > 1 {
		ipAddress = os.Args[1]
	} else {
		fmt.Print("Masukkan IP address: ")
		fmt.Scanln(&ipAddress)
	}

	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Printf("Melakukan request ke API dengan IP: %s\n", ipAddress)

	response, err := makeAPIRequest(ipAddress, config.APIKey)
	if err != nil {
		log.Fatalf("Error making API request: %v", err)
	}

	fmt.Printf("Berhasil mendapatkan %d domain(s):\n", len(response.Domains))
	for _, domain := range response.Domains {
		fmt.Printf("- %s\n", domain.Domain)
	}

	outputFile := "domains_result.json"
	if err := saveDomains(response, outputFile); err != nil {
		log.Fatalf("Error saving domains: %v", err)
	}

	fmt.Printf("Hasil telah disimpan ke: %s\n", outputFile)
}