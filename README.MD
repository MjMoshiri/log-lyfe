<p align="center">
  <img src="icon.png" width="400">
</p>

# **Audit your logs**

Initially undertaken as a self-study project and a showcase of my abilities, gol is a high-performance log auditing microservice. Built with Go for server-side operations and Cassandra for efficient data storage, it emphasizes minimal external dependency, availability, scalability, and low latency.

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
curl -X POST -H "Authorization: canon" -H "Content-Type: application/json" -d @data.json http://localhost:8080/insert
```
#### ***Automated Data Insertion*** :

Use the `insert.sh` script to bulk insert JSON data objects from `insert.json` file.

- **Prerequisites** :
  - Install `jq` package 
- **Usage** :
```bash
  chmod +x insert.sh
  ./insert.sh
  ```
####  ***Advanced Interactions*** 
If you'd like to delve deeper into the application or perform tasks like running tests or generating documentation, there's a Makefile located in the `/gol` directory packed with useful commands.


## **Known Limitations & Areas for Improvement**

Even though I've paused development on this project for now, these are the areas I believe have immense potential for further enhancement and deeper learning:

### **Testing**:
- I'm still honing my skills in test implementation, so this version might not have exhaustive API-related testing.
- Keen to dive deeper into benchmarks, which are currently not in place.
- While I've addressed many errors there are still some that need to be handled for instance:
  - The HTTP body size isn't limited, posing potential blockage risks.
  - There's room to address potential injection vulnerabilities, especially with the gocql package, before rolling into production.

### **Code Structure**:
- I've made a conscious effort to segment packages. But, when it comes to API-related packages, there's potential for refining through more distinct interfaces and organization.
- Some parts of the codebase are hardcoded – this will need reconsideration before a production rollout.
- Comprehensive logging and control mechanisms are areas ripe for expansion.

### **Architecture & Optimization**:
- I've used advanced techniques like the json-iterator, which trumps the default JSON parser in speed, and have utilized sync.Pool for the insert handler given the high traffic it handles. Moreover, I've integrated a bucketing strategy for data distribution.
- Looking ahead:
  - Incorporating caching layers can yield better results.
  - Adjustments in database settings, including the possibility of a Token Aware policy, should be on the radar.
  - Splitting read/writer nodes can significantly refine optimization and offer a clearer distinction in roles.
  - The use of a message queue can help in decoupling the application and database, and can also help in handling spikes in traffic.
