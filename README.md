# WebWatch: Real-Time Website Analytics

## Table Of Contents

1. [Project Overview](#project-overview)
2. [Features](#features)
3. [Technology Stack](#technology-stack)
4. [Architecture](#architecture)
5. [Setup & Installation](#setup--installation)
6. [Usage](#usage)
7. [Kafka Topics](#kafka-topics)
8. [InfluxDB Schema](#influxdb-schema)
9. [Grafana Dashboards](#grafana-dashboards)
10. [Future Enhancements](#future-enhancements)
11. [Contributing](#contributing)
12. [How the visualisations look](#how-the-visualisations-look)
13. [Acknowledgments](#acknowledgments)

### 1. Project Overview

WebWatch is a real-time website analytics tool designed to track and visualize various user interactions on websites. It provides insights into user behavior by capturing data such as page views, clicks, scroll depth, mouse movements, video plays, and more.

### 2. Features

- **Real-time analytics**: Get instant insights into user behavior on your website.
- **Multi-platform support**: Track user interactions across various devices and browsers.
- **Customizable dashboards**: Create personalized dashboards to focus on key metrics.
- **Page Views**: Monitor which pages are being viewed and how often.
- **Clicks**: Capture click events on different elements to understand user interaction.
- **Session Duration**: Track how long users stay on your website.
- **Scroll Depth**: Measure how far users scroll down each page.
- **Mouse Movements**: Visualize mouse activity intensity across your pages.
- **Form Submissions**: Track form interactions to monitor user engagement.
- **Video Plays/Completions**: Analyze user interaction with video content.
- **Downloads**: Monitor file download activity from your site.
- **Image Views**: Track and analyze which images are being viewed and when.

### 3. Technology Stack

#### Frontend:

- **JavaScript**: Used for the Tracking.js to capture user interactions.
- **React.js**: For building the User Dashboard, I haven't built Dashboard yet.
- **React-Bootstrap**: Utilized for UI Components.

#### Backend:

- **Golang**: Handles data ingestion, processing, and Kafka integration.

#### Data Streaming:

- **Apache Kafka**: Used for event streaming.

#### Data Storage:

- **InfluxDB**: A time series database for real-time data storage.

#### Visualization:

- **Grafana**: Used for dashboard creation and visualization.

#### Other Tools:

- **Docker**: For containerization of the application.
- **Kubernetes** (Optional): For container orchestration.

### 4. Architecture

#### Frontend:

- **Embedded JavaScript**: `tracking.js` is embedded on the client website to capture user interactions.
- **React-based Dashboard**: Used for data visualization and user interaction. (Working on it)

#### Backend:

- **Golang Server**: Handles incoming data and integrates with Kafka for event streaming.
- **Kafka Producers**: Sends events to the appropriate Kafka topics based on event types.
- **Kafka Consumers**: Processes events from Kafka topics and stores the data in InfluxDB.

#### Data Storage:

- **InfluxDB**: Used for storing both raw and aggregated data for analysis.

#### Visualization:

- **Grafana**: Creates dashboards and visualizes the data stored in InfluxDB.

### 5. Setup and installation

#### Prerequisites:

- **Golang**: Install Go on your machine.
- **Docker**: Install Docker on your machine.
- **Kubernetes** (Optional): Install Kubernetes on your machine.
- **Kafka**: Install Kafka on your machine.
- **InfluxDB**: Install InfluxDB on your machine.
- **Grafana**: Install Grafana on your machine.
- **React**: Install React on your machine.

#### Installation Steps:

1. Clone the repository.

```
    git clone https://github.com/akhilk2802/WebWatch
    cd webwatch
```

2. Setup Kafka

   - Start Zookeeper

   ```
   bin/zookeeper-server-start.sh config/zookeeper.properties
   ```

   - Start Kafka

   ```
   bin/kafka-server-start.sh config/server.properties
   ```

   - Create topics (Optional, coz after running the project, it will create by itself)

3. Setup InfluxDB

   - Start InfluxDB and configure your database, organization, and buckets.
   - Generate an InfluxDB token for authentication.

4. Setup Grafana

   - Start Grafana and configure your InfluxDB connection.
   - Import predefined dashboards or create custom ones. (Working on to export existing dashboards)

5. Environment Variables

```
KAFKA_BROKER_URL=

KAFKA_TOPIC_PAGEVIEW=
KAFKA_TOPIC_CLICK=
KAFKA_TOPIC_DURATION=
KAFKA_TOPIC_SCROLL=
KAFKA_TOPIC_MOUSEMOVE=
KAFKA_TOPIC_HOVER=
KAFKA_TOPIC_FORM_SUBMISSION=
KAFKA_TOPIC_FIELD_FOCUS=
KAFKA_TOPIC_FIELD_BLUR=
KAFKA_TOPIC_IDLE_TIME=
KAFKA_TOPIC_VIDEO_PLAY=
KAFKA_TOPIC_VIDEO_COMPLETION=
KAFKA_TOPIC_AUDIO_PLAY=
KAFKA_TOPIC_DOWNLOAD=
KAFKA_TOPIC_IMAGE_VIEW=


SERVER_PORT=
KAFKA_GROUP_ID=
INFLUX_TOKEN=
INFLUX_BUCKET=
INFLUX_ORGANISATION=
```

6. Run the backend

```
cd backend
go run cmd/main.go
```

7. Embed tracking.js

   - Add the tracking.js script to your website to start capturing user data. (in index.html)

8. View Dashboards

   - Access Grafana at http://localhost:3000 to visualize the captured data.

### 6. Usage

- Embed the tracker

  - Add the following script to the <head> or <body> section of your HTML:

  ```
  <script src="path/to/file/tracking-js></script>
  ```

- Accessing the Dashboard
  - Access Grafana at http://localhost:3000 to visualize the captured data.
  - Explore predefined dashboards or create custom visualizations based on your data.

### 7. Kafka Topics

- **KAFKA_TOPIC_PAGEVIEW**: Tracks page views.
- **KAFKA_TOPIC_CLICK**: Tracks click events.
- **KAFKA_TOPIC_DURATION**: Tracks session duration.
- **KAFKA_TOPIC_IDLE_TIME**: Tracks idle time.
- **KAFKA_TOPIC_VIDEO_PLAY**: Tracks video plays.
- **KAFKA_TOPIC_VIDEO_COMPLETION**: Tracks video completions.
- **KAFKA_TOPIC_AUDIO_PLAY**: Tracks audio plays.
- **KAFKA_TOPIC_AUDIO_COMPLETION**: Tracks audio completions.
- **KAFKA_TOPIC_IMAGE_VIEW**: Tracks image views.

### 8. InfluxDB Schema

- **Measurement**: `pageviews`, `clicks`, `session_durations`, `scrolls`, etc.
- **Tags**: `url`, `target`, `data_type` (raw/aggregated).
- **Fields**: `_value`, `scroll_percentage`, `avg_duration`, etc.

### 9. Grafana Dashboards

- **Page Views**: Time Series, Heatmap.
- **Clicks**: Bar Gauge, Stat Panel.
- **Session Duration**: Gauge, Histogram.
- **Scroll Depth**: Time Series, Heatmap.
- **Mouse Movements**: Heatmap, Geospatial Panel.
- **Form Submissions**: Pie Chart, Time Series.
- **Video Plays/Completions**: Gauge, Stat Panel.
- **Downloads**: Bar Gauge, Pie Chart.

- **To Use existing Dasboard**: Import from the `grafana-dashboard` folder


### 10. Future Enhancements

- **User Authentication**: Allow multiple users to sign up and use the service.
- **Multi-tenant Support**: Allow users to create and manage their own data dashboards.
- **Advanced Analytics**: Implement machine learning algorithms for predictive analytics.
- **Scalability**: Deploy on AWS/GCP for scalability and reliability.

### 11. Contributing

- Fork the repository.
- Create a new branch for your feature.
- Submit a pull request for review.

### 12. How the Visualisations look

- ![Visualisation](/images/grafana.jpg)

### 13. Acknowledgments

- Special thanks to the creators of InfluxDB, Kafka, and Grafana.
- Inspiration from various open-source analytics projects.
