[ Start server ] 
- php -S localhost:4444 -t ./


[ Hello-world ]
http://localhost:4444/index.php?query={hello}


==> respond
{
    data: {
        hello: "Hello World!"
    }
}


[ GET BOOK ]
http://localhost:4444/index.php?query={book(id:1){title,author}}
http://localhost:4444/index.php?query={book(id:1){title}}


[ API - GET - Book - bY Method Post]

curl -X POST \
-H "Content-Type: application/json" \
-d '{"query": "{book(id:1){title,author}}"}' \
http://localhost:4444/index.php