<template>
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
<div class="container-sm">
<div class="row">
<div class="col-12">
<div class="card bg-light" >
  <div class="card-body bg-dark text-white">
<p>{{ recievedMessage.body}}</p>

</div>
    <div class="card-footer">

    <form ref="chat" :action="sendMessage" @click.prevent="onSubmit">
<div class="input-group mb-3">
  <input v-model="message" type="text" class="form-control" placeholder="Type here...">
  <button class="btn btn-secondary" @click="sendMessage" type="submit">Send</button>
</div>
  </form>
  </div>
</div>
</div>
</div>


</div>
</template>

<script>


export default {
  name: 'App',
  data () {
    return {
      message : "",
      socket: null,
      recievedMessage : {
        type : "",
        body : ""
      }
    }
  },
  mounted() {
    this.socket = new WebSocket("ws://localhost:9100/ws")
    this.socket.onmessage = (msg) => {
      this.recievedMessage = JSON.parse(msg.data);
    }
  },
  methods: {
    sendMessage() {

      let msg = this.message;
      this.socket.send(msg);
      this.$refs.chat.reset();
      this.message = '';
    },

  }
}
</script>

