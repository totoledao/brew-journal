const socket = new WebSocket("ws://localhost:3000/ping");

socket.onclose = function (event) {
  console.log("Connection closed. Reloading the page...");
  location.reload();
};

// Send a ping message every 500ms
setInterval(function () {
  if (socket.readyState === WebSocket.OPEN) {
    socket.send("ping");
  }
}, 500);
