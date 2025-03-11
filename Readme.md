This project is a vulnerability scanner built in Go that scans files from GitHub repositories, processes vulnerability data, and stores it in a JSON file. It allows users to scan, query, and analyze security vulnerabilities in various open-source packages.

Table of Contents
Features
Prerequisites
Installation
Running Locally
Running with Docker
Usage
Scan Endpoint
Query Endpoint
Contributing
License
Features
Fetch vulnerability data from GitHub repositories.
Unmarshal vulnerability data into Go structures.
Store vulnerabilities in a local JSON file.
Query vulnerabilities based on filters such as severity.
Prerequisites
Ensure you have the following installed before using this project:

Go (version 1.18 or above) - for running the code locally.
Docker - for running the application in a containerized environment.
Installation
Running Locally
Clone the repository:

bash
Copy
git clone https://github.com/your-username/vns.git
cd vns
Install dependencies:

bash
Copy
go mod download
Build the application:

bash
Copy
go build -o main ./cmd/server
Run the application:

bash
Copy
./main
This will start a local server on port 8080.

Running with Docker
Build the Docker image:

bash
Copy
docker build -t vns .
Run the Docker container:

bash
Copy
docker run -p 8080:8080 vns
This will start the application in a Docker container, accessible at http://localhost:8080.

Usage
Scan Endpoint
To trigger a scan, send a POST request to the /scan endpoint with a JSON body containing the repository and files you want to scan.

Example Request:

json
Copy
{
  "repo": "velancio/vulnerability_scans",
  "files": ["vulnscan16.json"]
}
Example Command:

bash
Copy
curl -X POST http://localhost:8080/scan -d '{
    "repo": "velancio/vulnerability_scans",
    "files": ["vulnscan16.json"]
}' -H "Content-Type: application/json"
This will trigger the scan, process the vulnerabilities, and store them in payloads.json.

Query Endpoint
To query stored vulnerabilities, send a POST request to the /query endpoint with the desired filters (e.g., by severity).

Example Request:

json
Copy
{
  "filters": {
    "severity": "HIGH"
  }
}
Example Command:

bash
Copy
curl -X POST http://localhost:8080/query -d '{
    "filters": {
        "severity": "HIGH"
    }
}' -H "Content-Type: application/json"
This will return the stored vulnerabilities with the specified severity.

Contributing
We welcome contributions to this project! Feel free to fork the repository, make changes, and create pull requests.

Steps to contribute:
Fork the repository
Create a new branch (git checkout -b feature-name)
Make changes and commit (git commit -m 'Add new feature')
Push to the branch (git push origin feature-name)
Create a pull request
