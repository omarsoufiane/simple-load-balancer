package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// LoadBalancer struct holds information about the backend servers
type LoadBalancer struct {
	servers []string
	current int
}

// NewLoadBalancer creates a new instance of LoadBalancer
func NewLoadBalancer(servers []string) *LoadBalancer {
	return &LoadBalancer{
		servers: servers,
		current: 0,
	}
}

// NextServer selects the next server in a round-robin fashion
func (lb *LoadBalancer) NextServer() string {
	server := lb.servers[lb.current]
	lb.current = (lb.current + 1) % len(lb.servers)
	return server
}

// Handler function to handle incoming requests
func (lb *LoadBalancer) Handler(w http.ResponseWriter, r *http.Request) {
	selectedServer := lb.NextServer()
	fmt.Fprintf(w, "Request routed to server: %s\n", selectedServer)
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Read servers from environment variable
	serversString := os.Getenv("SERVERS")
	servers := strings.Split(serversString, ",")

	// Create a new instance of the load balancer
	loadBalancer := NewLoadBalancer(servers)

	// Set up the HTTP server to handle incoming requests
	http.HandleFunc("/", loadBalancer.Handler)
	fmt.Println("Load balancer running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
