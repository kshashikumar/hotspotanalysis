<template>
<div>
  <b-navbar toggleable="lg" type="dark" variant="info">
    <b-navbar-brand href="/" id="rez"><b-icon icon="bar-chart-fill" ></b-icon>Repo Analysis</b-navbar-brand>
    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
    <b-collapse id="nav-collapse" is-nav>
      <b-navbar-nav>
      <b-nav-item v-bind:href="'/commits/'+urlid">Commits</b-nav-item>
      <b-nav-item v-bind:href="'/charts/'+urlid">Charts</b-nav-item>
      </b-navbar-nav>
       <b-navbar-nav class="ml-auto">
           <b-nav-item href="/">BacK</b-nav-item>
       </b-navbar-nav>
    </b-collapse>
  </b-navbar>
</div>
</template>

<script>
export default {
  name: 'Navbar',
   props: {
    msg: {
      type: String
    }
  },
  data() {
    return {
      urlid:""
    };
  },
  watch: {
    msg: function() {
      this.urlid=this.msg;
    }
  },
  components: {
  }
}
</script>

<style>
#rez{
  color: white;
}
</style>
