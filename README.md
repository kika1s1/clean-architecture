# Task Management API Documentation
  ## Overview
  
  This API provides functionalities to manage tasks, including creating, updating, deleting, and retrieving tasks. It uses an in-memory database for data storage.
  
  ## Requirements
  
  ### Endpoints

  - **POST /register**
    - Register a user
    -  **POST /Response ** 
       - status code: `201`
       - Body: 
            ```json
            [
               {
                "id": 1,
                "username":"admin",
                "password":"admin",
                "role": "admin",

              }

            ]  
  - **POST /login**
    - Login a user
    -  **POST /Response** 
       - status code: `201`
       - Body: 
            ```json
            [
               {
                "username":"admin",
                "password":"admin",

              }

            ]  
  
  - **GET /tasks**
    - Get a list of all tasks.
    - **Response**: 
      - Status: `200 OK`
      - Body: 
        ```json
        [
          {
            "id": "1",
            "title": "Task 1",
            "description": "Description for Task 1",
            "due_date": "2024-07-31T00:00:00Z",
            "status": "pending"
          },
        ]
        ```
  
  - **GET /tasks/:id**
    - Get the details of a specific task.
    - **Response**: 
      - Status: `200 OK`
      - Body: 
        ```json
        {
          "id": "1",
          "title": "Task 1",
          "description": "Description for Task 1",
          "due_date": "2024-07-31T00:00:00Z",
          "status": "pending"
        }
        ```
  
  - **POST /tasks**
    - Create a new task.
    - **Request**:
      - Body:
        ```json
        {
          "title": "New Task",
          "description": "Description for the new task",
          "due_date": "2024-08-01T00:00:00Z",
          "status": "pending"
        }
        ```
    - **Response**: 
      - Status: `201 Created`
      - Body: 
        ```json
        {
          "id": "2",
          "title": "New Task",
          "description": "Description for the new task",
          "due_date": "2024-08-01T00:00:00Z",
          "status": "pending"
        }
        ```
  
  - **PUT /tasks/:id**
    - Update a specific task.
    - **Request**:
      - Body:
        ```json
        {
          "title": "Updated Task",
          "description": "Updated description",
          "due_date": "2024-08-02T00:00:00Z",
          "status": "completed"
        }
        ```
    - **Response**: 
      - Status: `200 OK`
      - Body: 
        ```json
        {
          "id": "1",
          "title": "Updated Task",
          "description": "Updated description",
          "due_date": "2024-08-02T00:00:00Z",
          "status": "completed"
        }
        ```
  - **DELETE /tasks/:id**
    - Delete a specific task.
    - **Response**: 
      - Status: `204 No Content`
  
  ### Error Handling
  
  - **Invalid Request**: 
    - Status: `400 Bad Request`
    - Response: 
      ```json
      {
        "error": "Invalid input"
      }
      ```
  
  - **Resource Not Found**: 
    - Status: `404 Not Found`
    - Response: 
      ```json
      {
        "error": "Task not found"
      }
      ```
  
  ## Instructions
  
  ### Development Environment
  
  1. **Install Go**: Ensure Go is installed on your system.
  2. **Setup Project**:
     ```sh
     git clone https://github.com/kika1s1/A2SV-Go-Learning-Path.git
     mkdir task_manager
     cd task_manager
     go mod init task_manager
     go mod tidy
     touch .env
        in .env
          MONGO_URI = "mongodb://127.0.0.1:27017"
          JWT_SECRET = "YOUR_JWT_SECRET"

3. **Folder Structure**:
```     
    task-manager/
    ├── Delivery/
    │   ├── main.go
    │   ├── controllers/
    │   │   └── controller.go
    │   └── routers/
    │       └── router.go
    ├── Domain/
    │   └── domain.go
    ├── Infrastructure/
    │   ├── auth_middleWare.go
    │   ├── jwt_service.go
    │   └── password_service.go
    ├── Repositories/
    │   ├── task_repository.go
    │   └── user_repository.go
    └── Usecases/
        ├── task_usecases.go
        └── user_usecases.go 
        ```

```


  - **main.go:** Entry point of the application.
  - **controllers/task_controller.go:** Handles incoming HTTP requests.
  - **models/task.go:** Defines the Task struct.
  - ***data/task_service.go:*** Contains business logic and data manipulation functions.
  - ***router/router.go:*** Sets up the routes and initializes the Gin router.
  - ***docs/api_documentation.md:*** 
  
  API documentation
  1. **Run the Application**:
      To run the application, open a terminal or command prompt and navigate to the `task_manager` directory. Then, execute the following command:
  
``` go run main go ```
## Testing
Unit tests are essential to ensure that the application works as expected. The tests are organized in the `Tests/` directory, with each test file focusing on different parts of the application.

### Test Files:

- **Tests/mocks/task_repository_mock.go**: Contains mock implementations of the `TaskRepository` interface for use in tests.
- **Tests/mocks/user_repository_mock.go**: Contains mock implementations of the `UserRepository` interface for use in tests.
- **Tests/domain_test.go**: Tests for domain models, ensuring they behave as expected.
- **Tests/task_usecases_test.go**: Tests for task-related use cases, verifying business logic.
- **Tests/user_usecases_test.go**: Tests for user-related use cases, ensuring correct behavior for user management.
- **Tests/auth_middleWare_test.go**: Tests for the authentication middleware, checking for proper token validation and user authentication.
- **Tests/jwt_service_test.go**: Tests for the JWT service, ensuring token creation and validation work as expected.
- **Tests/password_service_test.go**: Tests for the password service, verifying hashing and validation of passwords.

### Running the Tests

To run the tests, use the following command in the root of the project:

`
go test ./... `
### Example Test Outputs

- **AuthMiddleware Test Example**:
  - **Valid Token**: 
    - **Expected Output**: `200 OK`
    - **Actual Output**:
    ```json
    {
      "message": "Access granted"
    }
    ```
  - **Invalid or Missing Token**: 
    - **Expected Output**: `401 Unauthorized`
    - **Actual Output**:
    ```json
    {
      "error": "Invalid token"
    }
    ```

- **Task Use Case Test Example**:
  - **Creating a Task**:
    - **Expected Output**: `201 Created`
    - **Actual Output**:
    ```json
    {
      "id": "1",
      "title": "New Task",
      "description": "Description for the new task",
      "due_date": "2024-08-01T00:00:00Z",
      "status": "pending"
    }
    ```
  - **Retrieving Tasks**:
    - **Expected Output**: `200 OK`
    - **Actual Output**:
    ```json
    [
      {
        "id": "1",
        "title": "Task 1",
        "description": "Description for Task 1",
        "due_date": "2024-07-31T00:00:00Z",
        "status": "pending"
      }
    ]
    ```
    ```
    === RUN   TestAuthMiddleware_ValidToken
    --- PASS: TestAuthMiddleware_ValidToken (0.003s)
        PASS ```


  [API Documentation](https://documenter.getpostman.com/view/36018169/2sA3kdAck4)





























































































































































