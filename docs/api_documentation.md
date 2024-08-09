# Task Management API Documentation

## Overview
The Task Management API allows you to manage tasks through a set of RESTful endpoints. You can create, read, update, and delete tasks using this API. The API is built using Go and the Gin framework, and it uses an in-memory database to store tasks.

## Base URL
http://localhost:8000


## Endpoints

### 1. Get All Tasks
**Endpoint:** `GET /tasks`

**Description:** Retrieves a list of all tasks.

**Response:**
- **200 OK**: Returns a JSON array of tasks.
  
**Example Request:**
curl -X GET http://localhost:8000/tasks


**Example Response:**
[
  {
    "id": "1",
    "title": "Task Manager Project",
    "description": "Add/View/Delete Tasks",
    "due_date": "2024-08-08T14:15:22Z",
    "status": "In Progress"
  },
  {
    "id": "2",
    "title": "Books Management Project",
    "description": "Add/View/Delete Books",
    "due_date": "2024-08-07T14:15:22Z",
    "status": "Completed"
  }
]


### 2. Get Task by ID
**Endpoint:** `GET /tasks/:id`

**Description:** Retrieves the details of a specific task by its ID.

**Parameters:**
- `id` (string) - The ID of the task.

**Response:**
- **200 OK**: Returns the details of the task.
- **404 Not Found**: Task not found.

**Example Request:**
curl -X GET http://localhost:8000/tasks/1

**Example Response:**
{
  "id": "1",
  "title": "Task Manager Project",
  "description": "Add/View/Delete Tasks",
  "due_date": "2024-08-08T14:15:22Z",
  "status": "In Progress"
}

### 3. Create a New Task
**Endpoint:** `POST /tasks`

**Description:** Creates a new task with the provided details.

**Request Body:**
- `id` (string) - The ID of the task.
- `title` (string) - The title of the task.
- `description` (string) - The description of the task.
- `due_date` (string) - The due date of the task in ISO8601 format.
- `status` (string) - The status of the task.

**Response:**
- **201 Created**: Task created successfully.
- **400 Bad Request**: Invalid JSON or missing required fields.

**Example Request:**
curl -X POST http://localhost:8000/tasks \
-H "Content-Type: application/json" \
-d '{
  "id": "3",
  "title": "New Task",
  "description": "This is a new task.",
  "due_date": "2024-08-09T14:15:22Z",
  "status": "Pending"
}'

**Example Response:**
{
  "message": "Task created"
}

### 4. Update a Task
**Endpoint:** `PUT /tasks/:id`

**Description:** Updates the details of a specific task by its ID.

**Parameters:**
- `id` (string) - The ID of the task.

**Request Body:**
- Any of the fields: `title`, `description`, `due_date`, `status`.

**Response:**
- **200 OK**: Task updated successfully.
- **400 Bad Request**: Invalid JSON or empty request body.
- **404 Not Found**: Task not found.

**Example Request:**
curl -X PUT http://localhost:8000/tasks/1 \
-H "Content-Type: application/json" \
-d '{
  "title": "Updated Task Title"
}'

**Example Response:**
{
  "message": "Task updated"
}


### 5. Delete a Task
**Endpoint:** `DELETE /tasks/:id`

**Description:** Deletes a specific task by its ID.

**Parameters:**
- `id` (string) - The ID of the task.

**Response:**
- **200 OK**: Task deleted successfully.
- **404 Not Found**: Task not found.

**Example Request:**
curl -X DELETE http://localhost:8000/tasks/1


**Example Response:**
{
  "message": "Task removed"
}

## Error Handling
The API returns appropriate HTTP status codes and error messages in the response body for different scenarios:

- **400 Bad Request:** Invalid JSON or empty request body.
- **404 Not Found:** The requested task was not found.
- **500 Internal Server Error:** An unexpected error occurred.

## Testing
You can use tools like Postman or `curl` to test the API endpoints. Make sure the server is running by executing `go run .` in the project directory.

## Notes
- This API uses an in-memory database, so all data will be lost when the server restarts.
- The API is built for demonstration purposes and does not include persistent storage.
