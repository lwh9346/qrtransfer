<template>
  <div>
    <h2>receiver</h2>
    <vue-qr :text="data"></vue-qr>
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