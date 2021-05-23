# grpc_blog_api
A simple implementation grpc CRUD APIs for a blog MongoDB

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


### Generate pb.go from proto file (already done for this repo, can skip)

```
protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.
```

## Run

### Server
```
go run blog/blog_server/server.go
```
### Client
```
go run blog/blog_client/client.go
```

### Evans CLI
This is a gRPC Client that allows you to send gRPC requests to a server via Terminal

Start Evans CLI
```
evans -p 50051 -r
```

Using Evans CLI example commands

```
blog.BlogService@127.0.0.1:50051> show service
+-------------+------------+-------------------+--------------------+
|   SERVICE   |    RPC     |   REQUEST TYPE    |   RESPONSE TYPE    |
+-------------+------------+-------------------+--------------------+
| BlogService | CreateBlog | CreateBlogRequest | CreateBlogResponse |
| BlogService | ReadBLog   | ReadBlogRequest   | ReadBlogResponse   |
| BlogService | UpdateBlog | UpdateBlogRequest | UpdateBlogResponse |
| BlogService | DeleteBlog | DeleteBlogRequest | DeleteBlogResponse |
| BlogService | ListBlog   | ListBlogRequest   | ListBlogResponse   |
+-------------+------------+-------------------+--------------------+

blog.BlogService@127.0.0.1:50051> service BlogService

blog.BlogService@127.0.0.1:50051> call CreateBlog
blog::id (TYPE_STRING) => 
blog::author_id (TYPE_STRING) => "Jaiesh"
blog::title (TYPE_STRING) => "Evans Blog"
blog::content (TYPE_STRING) => "Creating blog using Evans CLI"
{
  "blog": {
    "id": "60a9fc2218eaf130a4d38ce0",
    "authorId": "\"Jaiesh\"",
    "title": "\"Jaiesh\"",
    "content": "\"Creating blog using Evans CLI\""
  }
}


blog.BlogService@127.0.0.1:50051> call ReadBLog
bbblog_id (TYPE_STRING) => 60a9fc2218eaf130a4d38ce0
{
  "blog": {
    "id": "60a9fc2218eaf130a4d38ce0",
    "authorId": "\"Jaiesh\"",
    "title": "\"Evans Blog\"",
    "content": "\"Creating blog using Evans CLI\""
  }
}
```

## Other Resources

- https://grpc.io/
- https://github.com/gogo/protobuf
- https://jbrandhorst.com/post/gogoproto/
