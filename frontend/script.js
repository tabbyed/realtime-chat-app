// Simple chat app for beginners

var selectedChat = 'general';

function changeChatRoom() {
    var chatInput = document.getElementById('chatroom');
    var newChat = chatInput.value;
    
    if (newChat) {
        selectedChat = newChat;
        document.getElementById('chat-header').textContent = newChat;
        chatInput.value = ''; // Clear input
    }
    return false; // Don't submit form
}

function sendMessage() {
    var messageInput = document.getElementById('message');
    var message = messageInput.value;
    
    if (message) {
        var chatArea = document.getElementById('chatmessages');
        chatArea.value += 'You: ' + message + '\n';
        conn.send(message)
        messageInput.value = ''; // Clear input
    }
    return false; // Don't submit form
}

// Set up the page when it loads
window.onload = function() {
    document.getElementById('chat-selection').onsubmit = changeChatRoom;
    document.getElementById('chatroom-message').onsubmit = sendMessage;
    
    if(window.WebSocket){
        console.log('supports websockets');
        conn = new WebSocket('ws://' + document.location.host + "/ws")  // Fixed!

        conn.onmessage = function(event) {
            const chatArea = document.getElementById('chatmessages');
            chatArea.value += 'Them: ' + event.data + '\n';
        };
    }
}