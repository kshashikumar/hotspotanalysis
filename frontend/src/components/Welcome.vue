<template>
  <div >
<div class="container">

    <b-button href="/commits" variant="primary" id="help">Analyze</b-button>
 
</div>

  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Home',
  data(){
    return {
      // name1:"",reponame1:"", 
      allcommits:"",
      message:[],
    }
  },
  methods: {
  },
  mounted() {
      axios.get('http://127.0.0.1:8090/').then(
          res => {
           this.message = res.data['commithistory']
           this.allcommits = String(res.data['allcommits'])
          });
  },

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
#primary{
    text-decoration-color: aliceblue;
}
#help{
  color: white;
}
</style>
