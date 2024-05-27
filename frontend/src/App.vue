<script>
export default {
  mounted() {
        document.body.classList.add('bg-light')
    },
    destroyed() {
        document.body.classList.remove('bg-light')
    }
}
</script>

<template>
  <router-view :key="$route.fullPath"></router-view>
</template>

<style>
  /* #app > div {
      border: dashed black 1px;
      display: inline-block;
      margin: 10px;
      padding: 10px;
      background-color: lightgreen;
    } */
</style>
