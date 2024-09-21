# Odin - Distributed HTTP Monitor for Kubernetes

Odin is a microservice-based HTTP monitor designed for distributed environments with built-in Kubernetes support. It performs health checks on specified endpoints in a CRON job format and integrates with popular cloud-native tools such as Grafana, Prometheus, and Helm for enhanced monitoring and scalability.

## Features

- **Health Checks**: Automatically checks the health of HTTP endpoints at specified intervals.
- **Microservice Architecture**: Designed to work seamlessly in distributed environments.
- **Kubernetes Support**: Full compatibility with Kubernetes for deployment and scaling.
- **CRON Job Format**: Schedule health checks using CRON expressions for flexibility.
- **Integration with Monitoring Tools**: Supports Grafana and Prometheus for monitoring and visualization.
- **Horizontal Pod Autoscaling**: Automatically scales the application based on demand.

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) (for containerization)
- [Kubernetes](https://kubernetes.io/) (for deployment)
- [Helm](https://helm.sh/) (for package management in Kubernetes)
- [Prometheus](https://prometheus.io/) (for monitoring)
- [Grafana](https://grafana.com/) (for visualization)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Manni-MinM/odin.git
   cd odin
   ```

2. Build the Docker image:

   ```bash
   docker build -t odin .
   ```

3. Deploy to Kubernetes using Helm:

   ```bash
   helm install odin ./charts/odin
   ```

### Configuration

You can configure the health checks and other settings by editing the configuration files in the `config` directory. 

### Usage

Once deployed, Odin will automatically start performing health checks on the specified endpoints. You can monitor the results using Grafana and Prometheus.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug reports.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
