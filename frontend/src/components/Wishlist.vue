<template>
<div class="container">
    <!-- <a  v-on:click="refresh">Refresh</a> -->
    <div class="card" >
      <md-list :md-expand-single="expandSingle" >
          <div v-for="(item,i) in Data1" :key="i">
        <md-list-item md-expand>
          <span class="md-list-item-text">Author:{{item.Authors}}  -   Repo:{{item.Repo}}</span>

          <md-list slot="md-expand">
            <md-list-item class="md-inset">Total Commits: {{item.Alcommits}}<span><a v-bind:href="'/commits/'+item.Id" class="btn btn-success" id="rock">Analyze</a><a href="/" v-on:click="deleted(item.Id)"  class="btn btn-danger" id="rock">Delete</a></span></md-list-item>
          </md-list>
        </md-list-item>
          </div>
      </md-list>
    </div>
  </div>
  
</template>

<script>
import axios from 'axios';

export default {
  name: 'Wishlist',
  data(){
      return {
        expandSingle: false,
       Data1:[],
       length:0,
       deletepath:"",
      }
  },
  methods: {
    deleted: function(id){
      this.deletepath='http://127.0.0.1:8090/'+id
      axios.delete(this.deletepath).then(res=>{
        console.log(res)
        console.log("deletion successful")
      })
    }
  },
//   methods: {
//      refresh: function () {
//          axios.get('http://127.0.0.1:8090/').then(res => {
//              console.log(res)
//              console.log(res.data.length)
//             for(var i=0;i<res.data.length;i++){
//               this.Data1.push({Authors:res.data[i].authorname, Repo:res.data[i].reponame, Alcommits:res.data[i].allcommits});
//             }
//             console.log(this.Data1)
//           })
//      }
//   },
mounted(){
    setInterval(() => {
        axios.get('http://127.0.0.1:8090/').then(res => {
            this.Data1=[]
            console.log(res)
            if(length!=res.data.length){
                      for(var i=0;i<res.data.length;i++){
              this.Data1.push({Authors:res.data[i].authorname, Repo:res.data[i].reponame, Alcommits:res.data[i].allcommits, Id:res.data[i]._id});
            }
            length==res.data.length
            }
        })
    }, 2000);
    // axios.get('http://127.0.0.1:8090/').then(res => {
    //          console.log(res)
    //          console.log(res.data.length)
    //         for(var i=0;i<res.data.length;i++){
    //           this.Data1.push({Authors:res.data[i].authorname, Repo:res.data[i].reponame, Alcommits:res.data[i].allcommits});
    //         }
    //         console.log(this.Data1)
    //       })
    //  }
}
}
</script>

<style lang="scss" scoped>
  $list-width: 320px;
#rock{
  color: white;
  margin-left: 1em;
}
  .full-control {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap-reverse;
  }

  .list {
    width: $list-width;
  }

  .full-control > .md-list {
    width: $list-width;
    max-width: 100%;
    height: 400px;
    display: inline-block;
    overflow: auto;
    border: 1px solid rgba(#000, .12);
    vertical-align: top;
  }

  .control {
    min-width: 250px;
    display: flex;
    flex-direction: column;
    padding: 16px;
  }
</style>