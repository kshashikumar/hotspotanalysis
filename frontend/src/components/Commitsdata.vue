<template>
  <div class="hello">
        <Navbar id="help"/>

    <h1 style="font-size:40px;">Total Commits: {{allcommits}}</h1>

<b-container>
  <b-table striped bordered small hover :items="message" :fields="fields"></b-table>
</b-container>
  <div>
  {{info}}
  </div>

  </div>
</template>

<script>
import axios from 'axios';
import Navbar from './Navbar'

export default {
  name: 'Commitsdata',
   components: {
    Navbar
  },
  data(){
    return { 
      allcommits:"",
      message:[],
      fields: [
          {
            key: 'author',
            sortable: true
          },
          {
            key: 'commit',
            sortable: true
          },
          {
            key: 'additions',
            sortable: true
          },
          {
            key: 'deletions',
            sortable: true
          },
          {
            key: 'changes',
            sortable: true
          }
      ],
       info: null,
    }
  },
  methods: {
     
     
  },
   mounted(){
        axios.get('http://127.0.0.1:8090/').then(res => {
           for(var i=0;i<res.data.length;i++){
                this.message = res.data[i].commithistory
               this.allcommits = String(res.data[i].allcommits)
           }

          //   this.message = res.data['commithistory']
          //  this.allcommits = String(res.data['allcommits'])
          //  for(var i=0;i<res.data.length;i++){
          //       console.log(res.data[i].allcommits);
          //  }
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
.help{
    position: fixed;
}
h1{
    top: 5%;
    height: 2;
}
</style>
