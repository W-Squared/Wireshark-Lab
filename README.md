# HTTP Basic Authentication Demo for Wireshark Lab

## Overview

This application demonstrates the use of HTTP Basic Authentication, an insecure protocol often used in early web applications. The goal of this lab is to help you understand how credentials are transmitted over the network in plain text when not using HTTPS. By running this application and analyzing the network traffic with Wireshark, you will learn why secure protocols like HTTPS are essential.

## Requirements

Before starting, ensure your machine meets the following requirements:

1. **Go Programming Language**  
   - Install Go version 1.20 or later.  
   - [Download and install Go](https://go.dev/dl/).

2. **Wireshark**  
   - Install Wireshark to capture and analyze network traffic.  
   - [Download Wireshark](https://www.wireshark.org/download.html).

3. **Terminal or Command Prompt**  
   - Access to a terminal or command-line tool to run the application and make HTTP requests.



## Instructions

### Step 1: Clone the Repository

1. Download or clone this repository to your machine:
   ```bash
   git clone https://github.com/w-squared/wireshark-lab.git
   cd wireshark-lab
   ```

### Step 2: Run the Application

1. Open a terminal and navigate to the directory containing the application files.
2. Start the server by running the following command:
   ```bash
   go run main.go
3. The server will start on http://localhost:8080


### Step 3: Capture Traffic in Wireshark

1. Open Wireshark.
2. Start a packet capture on the network interface used for local traffic (e.g., lo0 on macOS/Linux or Loopback on Windows).
3. Filter traffic in Wireshark using the following filter:
4. Copy code
5. http

### Step 4: Trigger an HTTP Request

1. Open your web browser.
2. Navigate to `http://localhost:8080`.
3. When prompted, enter the following credentials:
   - **Username:** `student1`
   - **Password:** `password123`.
4. If the credentials are correct, you will see a welcome message.

### Step 5: Analyze Traffic in Wireshark

1. Open Wireshark and locate the captured HTTP request sent to `http://localhost:8080`.
2. Select the HTTP packet in the capture list to view its details.
3. In the packet details pane, expand the **HTML Form URL Encoded** section.
4. Note the Form data is in plantext. It contains the username and password submitted in the login.
