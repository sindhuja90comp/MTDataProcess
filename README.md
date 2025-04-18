## About the Project
This project is a **Data Processing System** designed to simulate multiple worker threads processing data in parallel. The system consists of:
- **Worker Threads**: Retrieve data from a shared queue, process it, and save the results to a shared resource.
- **Concurrency Models**: Implemented in both **Java** and **Go**, leveraging their respective concurrency mechanisms.
- **Key Features**:
  - Efficient management of shared resources.
  - Mechanisms to avoid deadlocks.
  - Robust error handling for concurrent operations.

This project is implemented in **two programming languages**:
1. **Java**
2. **Go**

The goal of this project is to showcase how concurrency and parallel processing can be implemented in two different programming languages, leveraging their unique features.

---

## Prerequisites

### General Requirements
- Ensure you have **Git** installed to clone the repository.
- Clone the repository:
  ```bash
  git clone https://github.com/sindhuja90comp/MTDataProcess.git
  cd MTDataProcess
  ```

### Java Requirements
- **Java Development Kit (JDK)**: Version 8 or higher.
- **Apache Maven**: For dependency management.

### Go Requirements
- **Go**: Version 1.18 or higher.
- Properly configure `GOROOT` and `GOPATH` environment variables.

---

## How to Run

### Running Java Files
1. Navigate to the Java source directory:
   ```bash
   cd path/to/java
   ```
2. Compile using Maven:
   ```bash
   mvn compile
   ```
3. Run the program:
   ```bash
   mvn exec:java -Dexec.mainClass="com.example.MainClass"
   ```
   Replace `com.example.MainClass` with your main class name.

### Running Go Files
1. Navigate to the Go source directory:
   ```bash
   cd path/to/go
   ```
2. Run the Go program:
   ```bash
   go run main.go
   ```
   Replace `main.go` with the entry point file of your Go application if named differently.

---

## Notes
- Update the `path/to/java` and `path/to/go` with the respective paths to your Java and Go source directories within the repository.
- Make sure all dependencies are resolved before running the programs.

