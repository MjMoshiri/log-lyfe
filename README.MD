<p align="center">
  <img src="icon.png" width="400">
</p>

# **Audit your logs**

gol is a state-of-the-art log service written in Go for its server-side operations and Cassandra for backend data storage.

## **Understanding the Choice: Go & Cassandra**

### **Service Needs: A Non-functional Analysis**

Before delving into the technology stack, it's imperative to recognize the non-functional needs of the service:

- **Availability**: The service should be operational and accessible round-the-clock.

- **Consistency**: The service should be able to handle a large number of concurrent requests without compromising on the quality of responses.

- **Low Latency**: The service should be able to respond to requests in a timely manner.

- **Scalability**: As the demand grows, the service should gracefully handle an increased load.

- **Distributed Nature**: A distributed approach ensures enhanced resilience and potentially improved performance.

### **Go (Golang)**

- **Concurrency**: Born with concurrency in its DNA, Go deftly handles multiple simultaneous tasks.

- **Static Typing**: The rigorous type system ensures that a majority of potential pitfalls are identified during compilation, enhancing application stability.

- **Efficiency & Scalability**: Go's goroutines empower it to execute thousands of concurrent operations, cementing its reputation for scalability. Being a compiled language, it additionally promises superior performance.

- **Clean Syntax**: The clarity in Go's syntax assures ease of maintenance. Coupled with its comprehensive standard library, it simplifies development endeavors.

### **Cassandra**

- **Fast Writes**: In environments where continuous data inflow is the norm, like log services, Cassandra's write operation speed is invaluable.

- **Resilience**: Crafted with no single point of failure, Cassandra offers exceptional uptime guarantees.

- **Ready for Growth**: Cassandra's support for horizontal scalability ensures that growth is managed seamlessly.

## **Architecture & Structure**

### **Repository Structure**

The repository is organized into two primary folders:

- **gol/**: Contains the main project files.

- **cassandra/**: Holds the data-related components.

#### Inside `gol/`

- **`/cmd`**: This folder houses the main file that serves as the initial starting point of the app.

- **`/config`**: It stores configuration files essential for the app's operations, like database connection details.

- **`/internal`**: A dedicated space for the app's core components:
    - **`/api`**: Includes server application, and middleware.
    - **`/handlers`**:  Endpoint handlers functions.
    - **`/models`**: Hosts globally used models within the application, such as the event type.
    - **`/pkg`**: Contains utilities that might be needed by other services, like a JSON decoder.

- **`/storage`**: Here you'll find the in-app database schema and the database interface implementation.

### **Application Architecture**


#### **Endpoints** :

All endpoints require an `Authorization` header. The default password for this header is `canon`.

- **POST /insert**: Accepts data in JSON format. Takes an event and stores it into the database.

- **GET /query**: Requires an additional `Authorization` header named `Query-Key` with the default value `shooter`. Accepts JSON input, searching the database for matching Key/Value pairs. An optional header `Fetch-Size` can be set to determine the number of results needed, capped at 10,000 results.

- **GET /ok**: Acts as a health check endpoint.

- **GET /info**: Can be accessed to fetch app-related information.

The server operates on a port defined in `/gol/config/server.yaml` (default is set to 8080). Once a valid database connection is established, it starts listening for incoming requests.

## **Run & Use**

### **Docker Setup**

The application is dockerized for ease of deployment and consistency across different environments. To get the application up and running:

```bash
docker-compose up
```
This command will set up and start the entire application. Server configurations can be found and modified in `gol/config/server.yaml`.

### **Making Requests**

You can interact with the server using tools like `curl` and ensure you're providing the appropriate headers (refer to the [Endpoints](#endpoints-) section for details)

#### ***Example Request*** :

```bash
curl -X POST -H "Authorization: canon" -H "Content-Type: application/json" -d @insert.json http://localhost:8080/insert
```

####  ***Advanced Interactions*** 
If you'd like to delve deeper into the application or perform tasks like running tests or generating documentation, there's a Makefile located in the `/gol` directory packed with useful commands.
