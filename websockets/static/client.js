let ws = new WebSocket("ws://127.0.0.1:3000/socket");

const textbox = document.querySelector(".text-box");
const result = document.querySelector(".result");

textbox.addEventListener("keydown", function(event) {
	if (event.key === "Enter") {
		event.preventDefault();
		let text  = textbox.value; 
		ws.send(text);
		textbox.value = "";
	}
});

ws.addEventListener("open", function(event) {
	console.log("ws : new connection established");
})

ws.addEventListener("close", function(event) {
	console.log("connection closed");
});

ws.addEventListener("error", function(event) {
	console.log("connection error : "+ event.data);
});

ws.addEventListener("message", function(event) {
	let el = document.createElement("p");
	el.textContent = "server : " + event.data;
	result.appendChild(el);
});


