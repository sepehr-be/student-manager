
# Student Management API

A simple student management system built with Go and Echo framework.

## Features

- Add a new student
- Add a course to an existing student
- Get all students
- Calculate average grade
- Find the highest grade student
- Find the lowest grade student

## Endpoints

### 1. Get All Students
- **Method:** GET
- **Route:** `/students`
- **Description:** Returns a list of all students.

### 2. Create New Student
- **Method:** POST
- **Route:** `/students`
- **Body:**
```json
{
  "name": "John Doe",
  "age": 20,
  "grade": 85.5,
  "course": ["Math", "Science"]
}
```
- **Response:**
```json
{
  "id": 1
}
```

### 3. Add Course to Student
- **Method:** POST
- **Route:** `/students/course/:id`
- **Body:**
```json
{
  "course": "History"
}
```
- **Description:** Adds a new course to the student with the given ID.

### 4. Calculate Average Grade
- **Method:** GET
- **Route:** `/students/average`
- **Description:** Returns the average grade of all students.

### 5. Get Highest Grade Student
- **Method:** GET
- **Route:** `/students/highest`
- **Description:** Returns the student with the highest grade.

### 6. Get Lowest Grade Student
- **Method:** GET
- **Route:** `/students/lowest`
- **Description:** Returns the student with the lowest grade.

## Project Structure

```
student-app/
├── dto/
├── entity/
├── handler/
├── repository/
├── studentservice/
└── main.go
```

## Getting Started

```bash
go run main.go
```

## Requirements
- Go 1.20 or higher

## Author
[Sepehr]

---
Feel free to contribute or raise issues!
