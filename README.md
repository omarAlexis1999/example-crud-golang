The project is simple and consists of a database with a table of employees, and a CRUD with which the information is managed

Making use of the Golang language

The database is called system1 and has one table "employee" with 3 attributes
id, name, email

Additional packages used are:

github.com/go-sql-driver/mysql v1.6.0
github.com/joho/godotenv v1.4.0

Requirements to run
- Golang installed
- Create a database called system1 in mysql with the table employee(id, name, email)
- The table can be imported with the project's ./employee.sql file
- Configure the project's .env file with the database credentials

To execute

go run main.go

run this command in project path