

const websocket = new WebSocket("ws://localhost:8080/ws")

websocket.onopen = function() {
    console.log("Websocket connection established")
}

websocket.onmessage = function(event) {
    console.log("Received message:", event.data)
}

websocket.onclose = function() {
    console.log("Websocket connection closed")
}

websocket.send("hello")

