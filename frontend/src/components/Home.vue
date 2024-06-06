<template>
 
  <div >
    <Navbar position="fixed" />
    <div class="row ">
    <div class="col-6" id="rog">
  <div class="card border-light bg-light mb-3" id="hell">
     <div class="card-body">
  <h5 class="card-title" id="details">Enter Details</h5>
  <form >
  <div class="form-group" style="display:flex; flex-direction: row; justify-content: center; align-items: center">
    <label for="exampleInputEmail1">Author name:</label>
    <input type="text" v-model="name" class="form-control" id="exampleInputEmail1"  placeholder="Enter Author name" style="margin-left:1.5em;">
  </div>
  <div class="form-group"  style="display:flex; flex-direction: row; justify-content: center; align-items: center">
    <label for="exampleInputPassword1">Repository name:</label>
    <input type="text" v-model="reponame" class="form-control" id="exampleInputPassword1" placeholder="Enter Repositry name">
  </div>
  <a href="#" v-on:click="postdata" class="btn btn-light">Submit</a>
</form>
      <b-card
    :title=reponame
    tag="article"
    style="max-height: 20rem;"
    class="mb-2"
  >
    <b-card-text>
     {{alcommits}}
    </b-card-text>
     <component :is="component"></component>
  </b-card>
     </div>
</div>
  </div>
  <div class="col-6" id="rog1">
  <Wishlist></Wishlist>
  </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import Welcome from './Welcome'
import Free from './Free'
import Wishlist from './Wishlist'
import Navbar from './Navbar'

export default {
  name: 'Home',
  components: {
   Welcome,
   Free,
   Wishlist,
   Navbar
  },
  data(){
    return {
      name: "",
      reponame: "",
      component: "Free",
      // name1:"",reponame1:"", 
      alcommits:"",
      message:[],
    }
  },
  methods: {
      getcom(){
        if(this.component!== Welcome){
          this.component = Welcome
        }
        
      },
      postdata: function(){
        var data = {"author":this.name , "reponame":this.reponame}
         console.log(data)
        axios({ method:"POST" , url:"http://127.0.0.1:8090/" , data: data, headers: {"content-type": "text/plain" } }).then(
          res => {
          //   for(var i=0;i<res.length;i++){
          //  this.message = res.data[i].commithistory
          //  this.allcommits = String(res.data[i].allcommits)
          //   }
           console.log(res)
           console.log("data posted")
           this.alcommits ="data posted"
          })
          setTimeout(function(){
            axios.get('http://127.0.0.1:8090/').then(res => {
              console.log(res.data)
           })
          },4000);
          this.getcom()
      },
  },
  
  // updated() {
  //   console.log("in updated")
  //     axios.get('http://127.0.0.1:8090/').then(res => {
  //           this.allcommits = res.data.allcommits
  //         })
  // },

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#main1{
  overflow: hidden;
}
#details{
  font-size: 30px;
  font-family: Arial;
}
h3 {
  margin: 40px 0 0;
}
#rog{
  /* background-image:url("https://www.lifewire.com/thmb/jJkYDdgVRPBtpQzhxn_t4RNDu8Q=/768x0/filters:no_upscale():max_bytes(150000):strip_icc()/GettyImages-486840926-424eecdec9d94ca2a12401816cb2ab12.jpg"); */
   background-color: #00203FFF;
   
}
#rog1{
  /* background-image: url("https://previews.123rf.com/images/iulias/iulias1611/iulias161101033/65667829-retro-vintage-colored-background-with-noise-effect-grunge-texture-with-different-color-patterns-yell.jpg");
  background-size: 500px;
  width: 100%;
  height: 1500px; */
    background-color: #ADEFD1FF;
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
  color: lightblue;
}
html {
  height: 100%;
}
body {
  margin:0;
  padding:0;
  font-family: sans-serif;
  background: linear-gradient(#141e30, #243b55);
}

#hell{
  margin-top: 10em;
  margin-left: 8em;
  margin-right: 8em;
  padding-left: 1em;
  padding-right: 1em;
  border-radius: 8px;
  
}

</style>


<!-- <div class="container">
    <b-container >
   <b-row class="my-1">
    <b-col sm="2">
      <label for="input1">Author:</label>
    </b-col>
    <b-col sm="10">
      <b-form-input id="input1" placeholder="Enter Author name" v-model="name"></b-form-input>
    </b-col>
  </b-row>

  <b-row class="my-1">
    <b-col sm="2">
      <label for="input2">Repository Name:</label>
    </b-col>
    <b-col sm="10">
      <b-form-input id="input2" placeholder="Enter Repo name" v-model="reponame"></b-form-input>
    </b-col>
  </b-row>
   <b-button variant="outline-primary" v-on:click="postdata">Submit</b-button>
</b-container>
  </div> -->