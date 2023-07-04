# final-project-3
Kanban Board

## Admin Credentials
- email: admin@gmail.com
- password: 123456

## Endpoints
Users
1. POST - https://final-project-3-production-1fec.up.railway.app/users/register - Register a user
2. POST - https://final-project-3-production-1fec.up.railway.app/users/login - Login
3. PUT - https://final-project-3-production-1fec.up.railway.app/users/update-account - Update a user
4. DELETE - https://final-project-3-production-1fec.up.railway.app/users/delete-account - Delete a user

Categories
1. POST - https://final-project-3-production-1fec.up.railway.app/category - Create a category
2. GET - https://final-project-3-production-1fec.up.railway.app/category - Get all categories
3. PATCH - https://final-project-3-production-1fec.up.railway.app/category/{categoryId} - Update a category
4. DELETE - https://final-project-3-production-1fec.up.railway.app/category/{categoryId} - Delete a category

Tasks
1. POST - https://final-project-3-production-1fec.up.railway.app/tasks - Create a task
2. GET - https://final-project-3-production-1fec.up.railway.app/tasks - Get all tasks
3. PUT - https://final-project-3-production-1fec.up.railway.app/tasks/{taskId} - Update task's title and description
4. PATCH - https://final-project-3-production-1fec.up.railway.app/tasks/update-status/{taskId} - Update status of a task
5. PATCH - https://final-project-3-production-1fec.up.railway.app/tasks/update-category/{taskId} - Update category of a task
6. DELETE - https://final-project-3-production-1fec.up.railway.app/tasks/{taskId} - Delete a task

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
