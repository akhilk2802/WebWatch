# WebWatch - Real Time Analytics Tool

## Features

1. **Real-Time Event Tracking**: Captures user interactions like page views, clicks, and session durations using a JavaScript tracker and processes them with a Go and Kafka-based backend.
2. **Data Processing and Aggregation**: Consumes events from Kafka, aggregates them, and stores the results in MongoDB for efficient analytics.

3. **Visualization and Monitoring**: Visualize data with Grafana dashboards and monitor system performance using Prometheus for alerting and health checks.

4. **Scalable and Extensible Architecture**: Easily scalable with additional Kafka brokers and MongoDB instances, and designed for extensibility to accommodate new event types and analytics features.

**Steps to start a kafka server**:

- bin/zookeeper-server-start.sh config/zookeeper.properties
- bin/kafka-server-start.sh config/server.properties

**To list all the kafka topics created**
kafka-topics.sh --list --bootstrap-server localhost:9092

**To Delete a kafka topic running on the system**
kafka-topics.sh --delete --topic <topicName> --bootstrap-server localhost:9092


