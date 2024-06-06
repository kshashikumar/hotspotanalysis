<template>
  <div class="hello">
        <Navbar1 id="help" :msg="this.path" />
<b-container>
 <b-col sm="10">
      <apexchart type="bar" height="350" :options="options" :series="series"></apexchart>
      <apexchart type="bubble" height="350" :options="chartOptions" :series="series1"></apexchart>
    </b-col>
</b-container>
  </div>
</template>

<script>
import axios from 'axios';
import Navbar1 from './Navbar1'
export default {
  name: 'Chartsdata',
   components: {
    Navbar1
  },
  data(){
    return {
      name: "",
      reponame: "",
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
      categories1:[],
      data2:[],
      changs:[],
      adds:[],
      dels:[],
     options : {
          chart: {
          id:'vuechart-example'
        },
        xaxis: {
          categories: [],
        }
        },
        series: [{
          name: 'commits',
          data: []
        }],
    series1: [{
            name: 'commits',
            data: []
          },
          {
            name: 'commits',
            data: []
          },
          {
            name: 'commits',
            data: []
          }],
          chartOptions: {
            chart: {
               id:'vuechart-example'
            },
            dataLabels: {
                enabled: false
            },
            fill: {
                opacity: 0.8
            },
            title: {
                text: 'Simple Bubble Chart'
            },
            xaxis: {
                 type: 'category',
                labels: {
               formatter: function (value) {
                     return value;
                                }
                        }
            },
           
          },
    }
  },
  methods: {
     
  },
   mounted() {
      this.path=this.$route.params.id
      this.path1='http://127.0.0.1:8090/commits/'+this.path
      axios.get(this.path1).then(res => {
           // for(var i1=0;i1<res.length;i1++){
           // this.message = res[i1].data['commithistory']
           // this.allcommits = String(res[i1].data['allcommits'])
           //  }
                       console.log(res)

           
            if(res.data.commithistory != undefined){
              this.categories1.push(res.data.commithistory) 
              for(var j=0;j<res.data.commithistory.length;j++){
                this.changs.push({x:res.data.commithistory[j].author,y:res.data.commithistory[j].commit,z:res.data.commithistory[j].changes })
                this.adds.push({x:res.data.commithistory[j].author,y:res.data.commithistory[j].commit,z:res.data.commithistory[j].additions })
                this.dels.push({x:res.data.commithistory[j].author,y:res.data.commithistory[j].commit,z:res.data.commithistory[j].deletions })
                this.data2.push(res.data.commithistory[j].commit)
                //console.log(res.data[i].commithistory[j].commit)
              }
            }
           
            this.series =[{data:this.data2}]
            this.series1 =[{data:this.changs}]
            //console.log(this.data1)
          //  this.xaxis= [{categories: this.categories1}]
          // this.series = [{data:this.data1}]
          //  console.log(this.categories1)
          //  console.log(this.data1)
          // console.log(res)
       })
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
</style>
