<template>
  <div class="hello">
        <Navbar1 id="help" :msg="this.path"/>

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
import Navbar1 from './Navbar1'

export default {
  name: 'Commitsdata',
   components: {
    Navbar1
  },
  data(){
    return { 
      allcommits:"",
      message:[],
      path:"",
      path1:"",
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
//      this.path=this.$route.params.id
//  const Http = new XMLHttpRequest();
// const url='http://127.0.0.1:8090/commits/'+this.path;
// Http.open("GET", url);
// Http.send();

// Http.onreadystatechange = () => {
//   console.log(Http.responseText)
// }
        this.path=this.$route.params.id
      this.path1='http://127.0.0.1:8090/commits/'+this.path
        axios.get(this.path1).then(res => {
            console.log(res)
                this.message = res.data.commithistory
               this.allcommits = String(res.data.allcommits)

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
