const ws = new WebSocket('ws://**.**.**.**:8080/ws');

ws.onopen = () => {
    console.log('Connected to the WebSocket server');
};

ws.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    const messagesDiv = document.getElementById('messages');
    const messageElement = document.createElement('div');
    messageElement.textContent = `${msg.username}: ${msg.message}`;
    messagesDiv.appendChild(messageElement);
};

function sendMessage() {
    const username = document.getElementById('username').value;
    const message = document.getElementById('message').value;
    ws.send(JSON.stringify({ username, message }));
}