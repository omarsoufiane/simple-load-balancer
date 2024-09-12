[![Bounties on BountyHub](https://img.shields.io/badge/Bounties-on%20BountyHub-yellow)]( https://bountyhub.dev?repo=omarsoufiane/simple-load-balancer)
# Go Load Balancer

This is a simple round-robin load balancer implemented in Go.

## Description

The load balancer distributes incoming HTTP requests across a list of backend servers in a round-robin fashion.

## Prerequisites

- Go installed
- `.env` file with server addresses

SERVERS=Server1,Server2,Server3


## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/omarsoufiane/go-load-balancer.git
   cd go-load-balancer

## Usage

Running Locally

go run main.go
The load balancer will start running on port 8080.

Running with Docker
Build the Docker image:


docker build -t go-load-balancer .
Run the Docker container:

docker run -p 8080:8080 go-load-balancer
The load balancer will be accessible on port 8080 within the Docker container.

## Configuration
To change the server addresses, modify the .env file and update the SERVERS variable.
SERVERS=Server1,Server2,Server3

no kubernetes yet


## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
