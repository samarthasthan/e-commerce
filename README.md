# Multi-Vendor E-commerce Platform

Welcome to the Multi-Vendor E-commerce Platform project! This project aims to create a scalable, feature-rich e-commerce platform, including an admin portal for managing the platform, a vendor portal for sellers, and a public website for customers.

## Technologies Used

- **Backend**: Golang, gRPC, RESTful APIs.
- **Frontend**: React.js (Admin Portal), Next.js (Vendor Portal & Public Website).
- **Containerization**: Docker for easy deployment and scaling.
- **Databases**: MySQL, MongoDB, Redis.
- **Monitoring and Logging**: Grafana, Loki, Prometheus, Zipkin.
- **Messaging**: Kafka.


## Changelog

### [v1.0.0] - 2024-06-10
- Sign-up and OTP verification implemented in the authentication service.
- Mail service is completed; any service can send mail using Kafka.

## Upcoming Changes

- Complete all basic authentication functionality, such as login, forgot password, etc.
- Use Redis to cache authentication sessions.


## High-Level Design

![High-Level Design](./others/designs/multi-vendor-e-commerce.png)
_High-level design of the multi-vendor e-commerce platform._

The project employs various technologies to ensure efficiency, scalability, and maintainability:

- **Microservices Architecture**: Utilizing Golang and modern frameworks for seamless communication across the platform.
- **Frontend Development**: React.js for the admin portal and Next.js for vendor portals and the public website.
- **Containerization**: Docker for efficient deployment and scaling of microservices.
- **Monitoring and Logging**: Grafana, Loki, Prometheus, and Zipkin for performance monitoring and analysis.
- **Data Storage and Caching**: MySQL, MongoDB, and Redis for efficient data handling.

## Screenshot

### Monitoring and Logging

![Monitoring and Logging](./others/images/Screenshot%202024-04-20%20at%206.29.09%20PM.png)
_Monitoring and logging dashboard._

### Tracing

![Tracing](./others/images/Screenshot%202024-06-03%20at%202.46.48%20AM.png)
_Tracing dashboard._

### Admin Dashboard 

![Admin Dashboard](./others/images/Screenshot%202024-04-21%20at%205.14.16%20PM.png)
_Admin portal dashboard._

## Getting Started

To get started with the project, follow these steps:

1. Clone the repository.
2. Install Docker and Docker Compose.
3. Run `make up-dev` or `docker-compose -f ./build/compose/compose.dev.yaml` to start the services.
4. Access the respective portals and websites via the provided URLs.


## Port Configuration

Here's a list of ports used by the project:

| Name                 | Port  |
| -------------------- | ----- |
| Admin Portal         | 3000  |
| Vendor Portal        | 3001  |
| Website              | 3002  |
| Broker               | 7000  |
| Broker-GraphQL       | 7000  |
| Authentication       | 8000  |
| MySql                | 3306  |
| Redis                | 6379  |
| Kafka-internal       | 29092 |
| Kafka-external       | 9092  |
| Grafana Dashboard    | 15000 |
| Grafana Loki         | 3100  |
| Prometheus           | 15002 |
| Zipkins              | 9411  |

## Contribution Guidelines

Contributions to the project are encouraged! Whether you're interested in adding new features, fixing bugs, or improving documentation, please submit pull requests. Follow the contribution guidelines outlined in `CONTRIBUTING.md`.

## Contact

For any questions, suggestions, or further assistance, feel free to reach out:

- **Twitter**: [@samarthasthan](https://twitter.com/samarthasthan)
- **Email**: [samarthasthan27@gmail.com](mailto:samarthasthan27@gmail.com)
