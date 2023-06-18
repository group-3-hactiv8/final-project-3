# final-project-3
Kanban Board

## Endpoints
Users
1. POST - http://localhost:8080/users/register - Register a user
2. POST - http://localhost:8080/users/login - Login
3. PUT - http://localhost:8080/users/update-account - Update a user
4. DELETE - http://localhost:8080/users/delete-account - Delete a user

Categories
1. POST - http://localhost:8080/categories - Create a category
2. GET - http://localhost:8080/categories - Get all categories
3. PATCH - http://localhost:8080/categories/{categoryId} - Update a category
4. DELETE - http://localhost:8080/categories/{categoryId} - Delete a category

Tasks
1. POST - http://localhost:8080/tasks - Create a task
2. GET - http://localhost:8080/tasks - Get all tasks
3. PUT - http://localhost:8080/tasks/{taskId} - Update task's title and description
4. PATCH - http://localhost:8080/tasks/update-status/{taskId} - Update status of a task
5. PATCH - http://localhost:8080/tasks/update-category/{taskId} - Update category of a task
6. DELETE - http://localhost:8080/tasks/{taskId} - Delete a task

## Group 3
1. Prinata Rakha Santoso - GLNG-KS06-005
2. Iqbal Hasanu Hamdani - GLNG-KS06-001
3. Angga Anugerah Saputro - GLNG-KS06-019

## Pembagian Tugas
Prinata Rakha Santoso
- Initialized App
- Users (Register, Login, Update, Delete)
- Tasks (Create, Update status of a task)

Iqbal Hasanu Hamdani
- Tasks (Get all) 
- Tasks (Update task's title and description)
- Tasks (Delete)

Angga Anugerah Saputro
- Categories (Create, Get all, Update, Delete)
- Tasks (Update category of a task)