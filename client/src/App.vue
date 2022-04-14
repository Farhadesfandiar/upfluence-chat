<template>
  <form ref="chat" :action="sendMessage" @click.prevent="onSubmit">
    <input v-model="message" type="text">
    <input type="submit" value="Send" @click="sendMessage">
  </form>
  <p>
    Here are the messages:
  </p>
  <p>
    {{ message }}
  </p>
  <div>
    <h3>Received message :</h3>
    <p>
      {{ recievedMessage }}
    </p>
  </div>
</template>

<script>


export default {
  name: 'App',
  data () {
    return {
      message : "",
      socket: null,
      recievedMessage : "",
      err : ""
    }
  },
  mounted() {
    this.socket = new WebSocket("ws://localhost:9100/ws")
    this.socket.onmessage = (msg) => {
      this.recievedMessage = msg.data;
    }
    this.socket.onerror = (error) => {
      this.err = error;
    }
  },
  methods: {
    sendMessage() {

      let msg = this.message;
      this.socket.send(msg);
      this.$refs.chat.reset();
    },

  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
