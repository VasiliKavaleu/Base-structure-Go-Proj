## Example structure of project

### Technologies
+ Gin Framework
+ GORM

Clone project, and install dependency
Create .env file in root directory in accordance with .env.example

```bash
# Install dependency
go mod download
```

Run database in container
```bash
docker-compose -f docker-compose.db.yml up --build
```

Run server
```bash
go run main.go
```

[After running, project available on this link](http://0.0.0.0:8080/)
