# go-websockets
Run the command go --run main.go-- to start the server.

## Usage

use online websocket client tester like,  [link](https://www.piesocket.com/websocket-tester)  https://www.piesocket.com/websocket-tester

OR Use it in --Javascript code--:

Once the server is running, you can connect to it using a WebSocket client. Here is an example WebSocket client in JavaScript:

> const socket = new WebSocket("ws://localhost:8080/ws");

> socket.addEventListener("open", event => {
>  console.log("WebSocket connected");
>  socket.send("Hello, server!");
> });

> socket.addEventListener("message", event => {
>  console.log(`Received message: ${event.data}`);
> });

> socket.addEventListener("close", event => {
>  console.log("WebSocket disconnected");
> });
