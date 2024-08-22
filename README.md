# WebWatch: Real-Time Website Analytics 

=============================================

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
12. [License](#license)
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