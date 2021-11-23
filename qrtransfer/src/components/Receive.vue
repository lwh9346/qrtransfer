<template>
  <div>
    <a :href="data" v-if="data != 'N'"
      >如果使用手机访问，可以尝试直接点击改链接，无需扫码</a
    >
    <vue-qr :text="data"></vue-qr>
    <p>如果二维码没有变动说明分享者还未分享二维码或所分享二维码是静态的</p>
    <p>本站无法分享静态二维码</p>
  </div>
</template>

<script>
import VueQr from "vue-qr";

function url(s) {
  var l = window.location;
  return (
    (l.protocol === "https:" ? "wss://" : "ws://") +
    l.hostname +
    (l.port != 80 && l.port != 443 ? ":" + l.port : "") +
    s
  );
}

export default {
  components: {
    VueQr,
  },
  data: function () {
    return {
      data: "N",
      ws: new WebSocket(url("/get")),
    };
  },
  mounted: function () {
    this.ws.onmessage = this.onMessage;
    this.ws.onclose = this.reconnectWS;
  },
  methods: {
    onMessage: function (event) {
      console.log(event.data);
      this.data = event.data;
    },
    reconnectWS: function (event) {
      console.log("reconnecting websocket:", event);
      this.ws = new WebSocket(url("/get"));
      this.ws.onmessage = this.onMessage;
      this.ws.onclose = this.reconnectWS;
    },
  },
};
</script>

<style>
</style>