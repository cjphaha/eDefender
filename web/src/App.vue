<template>
  <div id="app">
    <router-view v-if="isRouterAlice" />
  </div>
</template>

<script>
export default {
  name: 'App',
  provide () {
    return {
      reload: this.reload
    }
  },
  data () {
    return {
      isRouterAlice: true
    }
  },
  methods: {
    reload () {
      this.isRouterAlice = false
      this.$nextTick(function () {
        this.isRouterAlice = true
      })
    }
  }
}
</script>

<style type="text/css">
  html,body,#app,.el-container{
        /*设置内部填充为0，几个布局元素之间没有间距*/
        padding: 0px;
         /*外部间距也是如此设置*/
        margin: 0px;
        /*统一设置高度为100%*/
        height: 100%;
    }
</style>
