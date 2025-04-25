# js-go-benchmark
A simple benchmark for js server and go, No sql involved.

For js, just go into the js directory, all you need is one command:
`bun server.ts`
and then you can open the url: http://localhost:3000

For go, just go into the go directory, all you need is one command:
`go build`
and then you can run the output file, on windows it's stress.exe, 
then you can open the url: http://localhost:8080/api/v1/tasks/

Actually for js, I tested more framework, express, fastify, and hono, and situation on different runtime,  such as node and bun, The result is in benchmark-result.md
The code I give is just about hono, if interested you can add another code. The pressure tool used is go-stress-testing, you can use another too.
