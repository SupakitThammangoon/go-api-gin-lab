# Student API with Gin (6609650657)
---
## How to Run
```bash
go run .
```
---
## Database: students.db
- id: "001", name: "test1", major: "CS", gpa: "3.0"
- id: "002", name: "test2", major: "CS", gpa: "3.1"
---
## Test API: GET / GETByID / POST / PUT / DELETE

### GET all students
- GET /students
```bash
[
  {
    "id": "001",
    "name": "test1",
    "major": "CS",
    "gpa": 3
  },
  {
    "id": "002",
    "name": "test2",
    "major": "CS",
    "gpa": 3.1
  }
]
```
- Status code: 200 ok
---
### Get students by ID
- GET /students/001
```bash
  {
    "id": "001",
    "name": "test1",
    "major": "CS",
    "gpa": 3
  }
```
- Status code: 200 ok
---
### POST 
- POST /students
```bash
  {
    "id": "003",
    "name": "test3",
    "major": "CS",
    "gpa": 3.5
  }
```
- Status code: 201 Created

### POST (Validation: ID)
- POST /students
```bash
  {
    "id": "",
    "name": "test4",
    "major": "CS",
    "gpa": 3.5
  }
```
- Status code: 400 Bad Request
```bash
  {
    "error": "id must not be empty"
  }
```
---
### PUT
- PUT /students/003
```bash
  {
    "name": "test3",
    "major": "CS",
    "gpa": 3.3
  }
```
- Status code: 200 ok

### PUT (Validation: Name)
- PUT /students/003
```bash
  {
    "name": "",
    "major": "CS",
    "gpa": 3.3
  }
```
- Status code: 400 Bad Request
```bash
  {
    "error": "name must not be empty"
  }
```

### PUT (Validation: GPA)
- PUT /students/003
```bash
  {
    "name": "test3",
    "major": "CS",
    "gpa": -1
  }
```
- Status code: 400 Bad Request
```bash
  {
    "error": "gpa must be between 0.00 and 4.00"
  }
```

### PUT (Not Found)
- PUT /students/999
```bash
  {
    "name": "test3",
    "major": "CS",
    "gpa": 4.0
  }
```
- Status code: 404 Not Found
```bash
  {
    "error": "Student not found"
  }
```
---
### DELETE (No Content)
- DELETE /students/003
- Status code: 204 No Content

### DELETE (Not Found)
- DELETE /students/999

- Status code: 404 Not Found
```bash
  {
    "error": "Student not found"
  }
```