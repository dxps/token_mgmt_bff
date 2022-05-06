<template>
	<form :action="sendMessage" @click.prevent="onSubmit">
		<input v-model="message" type="text" />
		<input value="Send" type="submit" @click="sendMessage" />
	</form>
	<div>Server response: {{ this.serverResponse }}</div>
</template>

<script>
export default {
	data() {
		return {
			socket: null,
			message: "",
			serverResponse: "",
		}
	},
	methods: {
		sendMessage() {
			let msg = { greeting: this.message }
			this.socket.send(JSON.stringify(msg))
		},
	},
	mounted() {
		this.socket = new WebSocket("ws://localhost:8084/socket")
		this.socket.onmessage = (msg) => {
			this.serverResponse = msg.data
		}
	},
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
	margin: 40px 0 0;
}
ul {
	list-style-type: none;
	padding: 0;
}
li {
	display: inline-block;
	margin: 0 10px;
}
a {
	color: #42b983;
}
</style>
