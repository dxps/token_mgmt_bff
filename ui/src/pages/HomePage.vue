<template>
	<div class="flex flex-col h-screen">
		<div class="p-14 max-w-full mx-auto mb-auto">
			<div class="grid gap-1 grid-cols-1 md:grid-cols-4">
				<div class="max-w-xs">
					<span class="primary pb-8">Money Transfer</span><br />
					<button class="btn btn-xs btn-outline mr-5" @click="getToken">Get Token</button>
					<button class="btn btn-xs btn-outline" @click="getAccounts" :disabled="!hasToken">Get Data</button>
				</div>

				<div class="max-w-3xl col-span-3 mx-auto">
					<table class="table table-normal table-border">
						<thead>
							<tr>
								<th class="pr-64">account</th>
							</tr>
						</thead>
						<tbody>
							<tr v-for="account in accounts.data" :key="account.id" class="hover">
								<td class="primary pr-16">{{ account.name }}</td>
							</tr>
						</tbody>
					</table>
				</div>
			</div>
		</div>
		<footer class="footer p-2 flex flex-row items-center bg-base-200 neutral text-gray-400">
			<div class="grid-flow-col flex-row">
				<svg
					width="18"
					height="18"
					viewBox="0 0 24 24"
					xmlns="http://www.w3.org/2000/svg"
					fill-rule="evenodd"
					clip-rule="evenodd"
					class="fill-current"
				>
					<path
						d="M22.672 15.226l-2.432.811.841 2.515c.33 1.019-.209 2.127-1.23 2.456-1.15.325-2.148-.321-2.463-1.226l-.84-2.518-5.013 1.677.84 2.517c.391 1.203-.434 2.542-1.831 2.542-.88 0-1.601-.564-1.86-1.314l-.842-2.516-2.431.809c-1.135.328-2.145-.317-2.463-1.229-.329-1.018.211-2.127 1.231-2.456l2.432-.809-1.621-4.823-2.432.808c-1.355.384-2.558-.59-2.558-1.839 0-.817.509-1.582 1.327-1.846l2.433-.809-.842-2.515c-.33-1.02.211-2.129 1.232-2.458 1.02-.329 2.13.209 2.461 1.229l.842 2.515 5.011-1.677-.839-2.517c-.403-1.238.484-2.553 1.843-2.553.819 0 1.585.509 1.85 1.326l.841 2.517 2.431-.81c1.02-.33 2.131.211 2.461 1.229.332 1.018-.21 2.126-1.23 2.456l-2.433.809 1.622 4.823 2.433-.809c1.242-.401 2.557.484 2.557 1.838 0 .819-.51 1.583-1.328 1.847m-8.992-6.428l-5.01 1.675 1.619 4.828 5.011-1.674-1.62-4.829z"
					></path>
				</svg>
				<p>© 2022 - Security Week @ Signicat</p>
			</div>
			<div class="grid-flow-col flex-row flex-auto"></div>
			<div class="grid-flow-col gap-2 sm:place-self-center sm:justify-self-end">
				<a
					><svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" class="fill-current">
						<path
							d="M24 4.557c-.883.392-1.832.656-2.828.775 1.017-.609 1.798-1.574 2.165-2.724-.951.564-2.005.974-3.127 1.195-.897-.957-2.178-1.555-3.594-1.555-3.179 0-5.515 2.966-4.797 6.045-4.091-.205-7.719-2.165-10.148-5.144-1.29 2.213-.669 5.108 1.523 6.574-.806-.026-1.566-.247-2.229-.616-.054 2.281 1.581 4.415 3.949 4.89-.693.188-1.452.232-2.224.084.626 1.956 2.444 3.379 4.6 3.419-2.07 1.623-4.678 2.348-7.29 2.04 2.179 1.397 4.768 2.212 7.548 2.212 9.142 0 14.307-7.721 13.995-14.646.962-.695 1.797-1.562 2.457-2.549z"
						></path>
					</svg>
				</a>
				<a
					><svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" class="fill-current">
						<path
							d="M19.615 3.184c-3.604-.246-11.631-.245-15.23 0-3.897.266-4.356 2.62-4.385 8.816.029 6.185.484 8.549 4.385 8.816 3.6.245 11.626.246 15.23 0 3.897-.266 4.356-2.62 4.385-8.816-.029-6.185-.484-8.549-4.385-8.816zm-10.615 12.816v-8l8 3.993-8 4.007z"
						></path></svg
				></a>
				<a
					><svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" class="fill-current">
						<path
							d="M9 8h-3v4h3v12h5v-12h3.642l.358-4h-4v-1.667c0-.955.192-1.333 1.115-1.333h2.885v-5h-3.808c-3.596 0-5.192 1.583-5.192 4.615v3.385z"
						></path></svg
				></a>
			</div>
		</footer>
	</div>
</template>

<script>
import axios from "axios"

let sseClient

export default {
	computed: {
		hasToken() {
			return this.authn.token.length > 0
		},
	},
	data() {
		return {
			authn: {
				token: "",
				url: "http://localhost:9093/authn",
			},
			accounts: {
				url: "http://localhost:9093/accounts",
				data: [],
			},
			sse: {
				url: "http://localhost:9093/sse/stream",
				format: "plain",
				isConnected: false,
				checkConnTimer: null,
			},
			isExternalTxnAvail: true,
		}
	},
	methods: {
		async getToken() {
			const creds = { client_id: "joe", client_secret: "black" }
			const resp = await axios.post(this.authn.url, creds)
			this.authn.token = resp.data.token
			console.log("[getToken] Authn ok, got token", this.authn.token)
			this.sseConnect()
		},

		async getAccounts() {
			if (this.hasToken) {
				const headers = { Authorization: "Bearer " + this.authn.token }
				const resp = await axios.post(this.accounts.url, {}, { headers })
				this.accounts.data = resp.data
			}
		},

		sseOnMessage(msg) {
			console.log("[sseOnMessage]   Got:", msg)
			this.authn.token = msg

			if (msg.startsWith("AccountDebitDisabled")) {
				let accountID = msg.substr(21)
				console.log(`>>> Account with id="${accountID}" has debit as disabled.`)
				this.accounts.data.forEach((account) => {
					if (account.id === accountID) account.is_frozen = true
				})
			}
		},

		sseConnect() {
			sseClient = this.$sse.create({
				url: `${this.sse.url}?token=${this.authn.token}`,
				includeCredentials: this.includeCredentials,
				format: this.sse.format,
			})

			sseClient.on("message", this.sseOnMessage)

			sseClient.on("error", (e) => {
				// eslint-disable-line
				console.log("[sseConnect]   Got error:", e)
				this.sseReconnect()
			})

			sseClient
				.connect() // eslint-disable-line
				.then(() => {
					this.sse.isConnected = true
					console.log("[sseConnect]   SSE Connection established")
				})
				.catch((err) => {
					this.sse.isConnected = false
					console.log("[sseConnect]   SSE Connection failed:", err)
				})

			// Regular checks to keep SSE connection established.
			clearInterval(this.sse.checkConnTimer) // Just for sanity.
			this.sse.checkConnTimer = setInterval(() => {
				this.sseCheckConnection()
			}, 1000)
		},

		sseReconnect() {
			console.log(">>> Reconnecting ...")
			this.sseDisconnect()
			this.sseConnect()
		},

		sseCheckConnection() {
			if (this.sse.isConnected !== true) {
				this.sseReconnect()
			}
		},

		sseDisconnect() {
			if (sseClient) {
				sseClient.disconnect()
				sseClient = null
			}
		},
	},

	async mounted() {
		this.getAccounts()

		// this.sseConnect() // Connect to SSE by default.
		// Regular checks to keep SSE connection established.
		// this.sse.checkConnTimer = setInterval(() => {
		//	this.sseCheckConnection()
		// }, 1000)
	},

	beforeUnmount() {
		clearInterval(this.sse.checkConnTimer)
		this.sseDisconnect()
	},
}
</script>
