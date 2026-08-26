# API Control Center

Create simulated API endpoints for testing client behaviour, network conditions, and failure handling. Each endpoint runs as a workload in the selected Kubernetes namespace.

### Response handlers

**Static response** returns the configured response body for every request.

**Echo request** returns the request body to the caller.

### Traffic controls

Each endpoint can introduce a delay range, limit request and response transfer rates, or return a selected HTTP error for a percentage of requests.

### Kubernetes resources

Each API endpoint is backed by:

- A Deployment that runs the endpoint
- A ConfigMap containing its configuration
- A Service that exposes it inside or outside the cluster
