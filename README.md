# api-in-go

1. initialize the project

```terminal
go mod init _name_of_project_
```

<i>Note</i>

---

- if you plan to work locally, then you can choose an name
- if online, then you need to chooser a corresponding path to where the module will be hosted, such `github.com/<username>/name_of_project_on_your_github

This command will create `go.mod` where we can manage dependencies. You can think of it like `package.json` in NodeJs or `Nuget` in Asp.net.

- create a folder called `cmd` stand for command, it is a common convention.
- create a file called inside the `cmd` called `main.go`

2. Configuration
   First of all, configuration is a nightmare if not set up correctly, it is a blocker. So, we need to get it out of the way at the beginning.
   And since this project is for learning purposes, so we don't have a lot to worry about. The gaol of this step is to menage the configuration settings like the db configuration in centralized and secure manner.

- create a config folder
  create a file inside the config folder `config/config.go`

- Install a library called `viper`
  run in the terminal

```bash
go get github.com/spf13/viper
```

<i>Note</i>

---

Now, after searching the viper, I found it's an overkill for a small project

- create a file named `env` in the root directory of the project as the same level of `go.mod`.
  Copy and paste the following into the `.env`

```code
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdb
```

see [config](config/config.go)

you need to figure out what the values in the environment variables are and then use to connect to the database/server

3. Connect to the database
   Im gonna use PostgreSQL, for nothing but I like pain

- install postgresql driver

```bash
go get github.com/lib/pq
```

- create a folder called `internal` in the root directory of the project
- inside internal, create another folder called `db` and inside the folder create file named `db.go`
- setup the database connection. see [connection](internal/db/db.go)
