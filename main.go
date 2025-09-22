package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

// LoadBalancer struct holds information about the backend servers
type LoadBalancer struct {
	servers []string
	current int
	mu      sync.Mutex
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
	lb.mu.Lock()
	defer lb.mu.Unlock()
	server := lb.servers[lb.current]
	lb.current = (lb.current + 1) % len(lb.servers)
	return server
}

// Handler function to handle incoming requests
func (lb *LoadBalancer) Handler(w http.ResponseWriter, r *http.Request) {
	selectedServer := lb.NextServer()
	fmt.Fprintf(w, "Request routed to server: %s\n", selectedServer)
}

// HealthHandler provides health check endpoint for Kubernetes
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Read servers from environment variable
	serversString := os.Getenv("SERVERS")
	if serversString == "" {
		fmt.Println("Error: SERVERS environment variable is not set or empty")
		os.Exit(1)
	}
	servers := strings.Split(serversString, ",")

	// Remove any empty strings and trim spaces
	var validServers []string
	for _, s := range servers {
		s = strings.TrimSpace(s)
		if s != "" {
			validServers = append(validServers, s)
		}
	}
	if len(validServers) == 0 {
		fmt.Println("Error: No valid servers found in SERVERS environment variable")
		os.Exit(1)
	}

	// Create a new instance of the load balancer
	loadBalancer := NewLoadBalancer(validServers)

	// Set up the HTTP server to handle incoming requests
	http.HandleFunc("/", loadBalancer.Handler)
	http.HandleFunc("/health", HealthHandler)

	// Support configurable port via environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Load balancer running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
	}
}
