# This is the README for the backend setup for Interval.

## Download and install Golang binaries
https://golang.org/doc/install?download=go1.10.3.darwin-amd64.pkg

## Check the GOPATH
https://golang.org/doc/code.html#Workspaces

## SQLite3

Most Linux and Mac distributions already come with SQLite3. To check if it exists, type the following command in the terminal:
    
    sqlite3
    
If it doesn't exist, install it. Here is a reference:
http://www.codebind.com/sqlite/how-to-install-sqlite-on/

### Create the database

    sqlite3 Interval.db
    
#### Helpful commands
https://www.tutorialspoint.com/sqlite/sqlite_create_database.htm

    .databases              # Shows a list of databases.
    .tables                 # Shows a list of tables.
    .schema [Table Name]    # Shows the schema for a specifed table.

## Run the application

### Build and run the application
    
    cd $HOME/{path to main.go directory}
    go build && ./app