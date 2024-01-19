# assignment2go
#need to add postgresql database
##query:
```postgres
create table users (
	id INT primary key,
	name VARCHAR(255),
	email VARCHAR(255),
	address VARCHAR(255),
	age INT,
	created_at timestamp(6),
	updated_at timestamp(6),
	deleted_at timestamp(6)
)```
