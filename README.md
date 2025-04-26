# js-go-benchmark
A simple benchmark for js server and go, No sql involved.

For js, go into the js directory, all you need is just one command:
`bun server.ts`,
and then you can open the url: http://localhost:3000

For go, go into the go directory, all you need is one command:
`go build`,
and then you can run the output file, namely stress.exe if on windows, 
then you can open the url: http://localhost:8080/api/v1/tasks/

Actually for js, I tested more framework, express, fastify, and hono, and the situation on different runtime,  such as node and bun, The result is in benchmark-result.md. 
The code I give is just about hono, if interested you can add another code. 
The pressure tool used is go-stress-testing: https://github.com/link1st/go-stress-testing, you can use another too.
