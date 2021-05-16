# grpc_blog_api
A simple implementation grpc CRUD APIs for a blog with backend as MongoDB

## Setup

### Install Homebrew

Homebrew is an installation manager for OSX. To install homebrew and it's services management tool run this command

```
$ /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
$ brew tap homebrew/services
```

More details about brew installation [here](https://brew.sh/)

### Install Golang

Install golang version `1.16`

```
$ brew install go
```

Follow the installation instruction at [golang official site](https://golang.org/doc/install) for detailed instruction

### Install MongoDB

Install MongoDB version `4.4`

```
brew tap mongodb/brew

brew install mongodb-community@4.4
```

Start MongoDB

```
brew services start mongodb-community@4.4
```

Stop MongoDB

```
brew services stop mongodb-community@4.4
```