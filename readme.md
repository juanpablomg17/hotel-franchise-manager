# HotelApi - Backend

This project is a solution for hotel franchise management in ClubHub, built using Golang and following various methodologies and technologies. Below are the main features of the implementation:

## Technologies Used

- **Gin**: Gin was used as the web framework to build the REST API.
- **Hexagonal Architecture**: A hexagonal architecture was implemented to separate responsibilities and facilitate code maintenance and scalability.
- **Domain-Driven Design (DDD) Approach**: A domain-driven design approach was followed to model the system.
- **Concurrent Programming**: Concurrent programming was applied to improve efficiency and responsiveness of the system.
- **Retry Logic**: Retry logic was implemented to handle potential integration errors with libraries and external API calls.
- **MongoDB**: MongoDB was used as the NoSQL database engine.

## Running the Project

To run the project, follow the steps below:

1. Clone the repository from GitHub.

2. Make sure you have Docker and Docker Compose installed on your machine.

3. Navigate to the root directory of the project.

4. Run the following command in the terminal to build and start the Docker containers:

   ```
   docker-compose up -d --build
   ```

   This will create and run the necessary Docker containers to execute the application.

5. Once the containers are up and running, you can access the REST API at `http://localhost:8080`. You can test the different endpoints using a tool like Postman.

## Architecture Diagram and Solution Flow

The following diagram illustrates the architecture and flow of the solution:

![Architecture Diagram](architecture_diagram.png)

The REST API built with Gin receives client requests and directs them to the corresponding controller. The controller interacts with the application services, which implement the business logic and perform operations on the data repositories. The repositories communicate with the persistence layer, which consists of a SQL or NoSQL database, depending on the choice.

## Consuming Endpoints

To facilitate the visualization and testing of the endpoints, Swagger UI has been integrated into the application. Once the application is up and running, you can access the API documentation at `http://localhost:8080/docs/index.html`.

Here is a preview of how Swagger UI looks:

![Swagger UI](https://i.ibb.co/pwcLbz1/Screenshot-1.jpg)