# assignment3-4go
need to add postgresql database
## description
simple project with registration system and manipulating with users db using postgresql queries through gorm
## Full name of creator:
Damir Baigenzhin Bolatuly
## instruction:
* use
  ```bash
  go run main.go
  ```
  to launch server
* server is on 8080 port
* use query below to create db
## query for postgresql:
```postgresql
create table users (
	id INT primary key,
	name VARCHAR(255),
	email VARCHAR(255),
	address VARCHAR(255),
	age INT,
	created_at timestamp(6),
	updated_at timestamp(6),
	deleted_at timestamp(6)
)
```
## screenshots
![image](https://github.com/DamirBai/goproject/assets/123295047/854f5db6-4724-4d3e-b68d-b5431bec7d5d)
![image](https://github.com/DamirBai/goproject/assets/123295047/7b02f1de-9836-4fe3-aea4-b7f06a08e8f7)
![image](https://github.com/DamirBai/goproject/assets/123295047/8f9b7a75-a1c8-40a9-b0c2-c621d790aa13)
