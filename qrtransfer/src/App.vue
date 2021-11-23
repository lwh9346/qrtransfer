<template>
  <div id="app">
    <h1>实时二维码转发器</h1>
    <select v-model="model">
      <option disabled value="">选择作为分享者或接收者以开始</option>
      <option value="streamer">分享者</option>
      <option value="receiver">接收者</option>
    </select>
    <stream v-if="model == 'streamer'"></stream>
    <receive v-if="model == 'receiver'"></receive>
  </div>
</template>

<script>
import Stream from "./components/Stream.vue";
import Receive from "./components/Receive.vue";

export default {
  name: "App",
  components: {
    Stream,
    Receive,
  },
  data: function () {
    return {
      model: "",
    };
  },
};
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
