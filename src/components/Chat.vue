<template>
  <div>dsdas
    <input v-model="msg" type="text" value="">
<!--    <input type="button" @click="send" value="发送">-->
    <button v-on:click="websocketsend">Send Message</button>
    <ul>
      <li v-for="m in msgs" v-bind:key="m">{{ m }}</li>
    </ul>
  </div>

</template>

<script>
// import Vue from "vue";
// Vue.forceUpdate();
export default {
  name: "chat",
  data(){
    return{
      msgs: [],
      msg: "",
      connection: null,
      key: 1,
    }
  },
  change:function (){
    this.msgs = [];
    this.$forceUpdate()
  },
  methods: {
    initWebSocket() {
      try {
        if ('WebSocket' in window) {
          this.socket = new WebSocket("ws://localhost:8989/chatws")
        } else {
          console.log('您的浏览器不支持websocket')
        }
        this.socket.onopen = this.websocketonopen
        this.socket.onmessage = this.websocketonmessage
        this.socket.onclose = this.websocketclose
        this.socket.onerror = this.websocketonerror
      } catch (e) {
        console.error('e: ', e)
      }
    },
    websocketonopen() {
      console.log('WebSocket连接成功', this.socket.readyState)
      this.websocketsend()
    },
    websocketsend() {
      // this.socket.send(JSON.stringify(this.msg))
      this.socket.send(this.msg)
      // this.msgs.push(this.msg);
      // this.$nextTick(()=>{
      //   this.msgs.push(this.msg);
      // })
      // this.$forceUpdate()
    },
    websocketonmessage(e) {
      console.log('WebSocket获取数据', e)
        // const data = JSON.parse(e.data);
      this.msgs.push(e.data);
        // [].push.apply(this.cache, [data])
      // Vue.set(this.msgs, 1, "asdf");
      // this.msgs.splice(0, 1,"sd")
      console.log(this.msgs)
      this.$forceUpdate()
    },
    websocketclose(e) {
      console.log('connection closed (' + e.code + ')')
      this.reconnect()
    },
    websocketonerror(e) {
      console.log('WebSocket连接发生错误', e)
      this.reconnect()
    },
  },
  mounted() {
    this.initWebSocket()
  },
  watch:{
    msgs(val){
      this.msgs = val;
      console.log("kandao: " + val)
    }
  },
}
</script>

<style scoped>

</style>