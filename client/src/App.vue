<template>
    <div>
        <Navbar></Navbar>
        <Feed></Feed>
    </div>
</template>

<script>
import Navbar from "./components/Navbar";
import Feed from "./components/Feed";

export default {
  name: 'app',
  components: {
      Navbar,
      Feed
  }
}
</script>

<style lang="scss">
 html {
     font-size: 16px;
 }
</style>
