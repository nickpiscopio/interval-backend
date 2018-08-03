# This is the README for the backend setup for Interval.

## Download and install Golang binaries
https://golang.org/doc/install?download=go1.10.3.darwin-amd64.pkg

### Linux
https://golang.org/doc/install    
https://www.tecmint.com/install-go-in-linux/

Download and push to /usr/local

    wget -c https://storage.googleapis.com/golang/go1.10.3.linux-amd64.tar.gz
    shasum -a 256 go1.10.3.linux-amd64.tar.gz
    sudo tar -C /usr/local -xvzf go1.10.3.linux-amd64.tar.gz
    go get -v -u github.com/gorilla/mux
    go get -v -u github.com/mattn/go-sqlite3
    sudo apt-get update
    sudo apt install gcc
    
Set the path

    export PATH=$PATH:/usr/local/go/bin

## Check the GOPATH
https://golang.org/doc/code.html#Workspaces

## SQLite3

Most Linux and Mac distributions already come with SQLite3. To check if it exists, type the following command in the terminal:
    
    sqlite3
    
If it doesn't exist, install it

Linux

    sudo apt install sqlite3
    
After installing, check again by typing:

    sqlite3

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