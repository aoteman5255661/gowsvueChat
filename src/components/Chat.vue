<template>
  <div>dsdas
    <input v-model="msg" type="text" value="">
<!--    <input type="button" @click="send" value="发送">-->
    <button v-on:click="websocketsend">Send Message</button>
<!--    <button v-on:click="websocketsendimg">发送图片</button>-->
    <input type='file' class="uploadphoto" @change="uploadphoto($event)" ref="inputer" multiple accept="image/png,image/jpeg,image/gif,image/jpg">

    <img id="img1" src=""/>
    <!--    <ul>-->
      <div v-for="(m, index) in msgs" :key="index">
<!--        <div>index</div>-->
        <div v-if="m['src']"><img :src="m['src']"  style="width:100px;height:130px" /></div>
        <div v-else>{{m['msg']}}</div>
      </div>
<!--    </ul>-->
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
      
      this.socket.send(JSON.stringify({"msg": this.msg}))
      // this.msgs.push(this.msg);
      // this.$nextTick(()=>{
      //   this.msgs.push(this.msg);
      // })
      // this.$forceUpdate()
    },
    uploadphoto(e){
      var file = e.target.files[0];
      var filesize = file.size;
      // var filename = file.name;
      // 2,621,440   2M
      if (filesize > 2101440) {
        // 图片大于2MB

      }
      var reader = new FileReader();
      reader.readAsDataURL(file);
      console.log(this.socket)
      let _this = this
      reader.onload = function (e) {
        console.log(this.socket)
        // 读取到的图片base64 数据编码 将此编码字符串传给后台即可
        var imgcode = e.target.result;
        console.log(imgcode);
        console.log(JSON.stringify({"src": imgcode}))
        // document.getElementById("img1").src=imgcode;
        _this.socket.send(JSON.stringify({"src": imgcode}))
        // this.socket.send({"src": imgcode})

        console.log(JSON.stringify({"src": imgcode}))

      };

    },
    websocketonmessage(e) {
      console.log('WebSocket获取数据', e)
        // const data = JSON.parse(e.data);
      console.log(e.data)
      console.log(JSON.parse(e.data)['msg'])
      this.msgs.push(JSON.parse(e.data));
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