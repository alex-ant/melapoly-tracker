<template>
  <div class="container-fluid bgmain">
    <img class="img-fluid" src="https://raw.githubusercontent.com/alex-ant/melapoly-tracker/master/fe/melapoly-tracker/img/logo-no-bg.png" />
    <br/><br/>
    <div v-if="!validToken" class="container-fluid gameinfo">
      <input placeholder="Enter your name" class="name-input" type="text" v-model="regNameModel">
      <br/><br/>
      <button type="button" v-on:click="registerPlayer()" class="action-button">MOVE</button>
    </div>

    <div v-else class="container-fluid gameinfo">
      <!-- Current player info -->
      <div class="row">
        <div class="col current-user-name">Hello, {{currentPlayer.name}}!</div>
        <div class="w-100"></div>
        <div class="col cash">€{{currentPlayer.cashAmount}}</div>
      </div>

      <!-- Players list -->
      <div v-for="player in playersData" class="container-fluid playerinfo" v-bind:key="player.id" v-bind:class="{'current-user-list': player.id ===  currentPlayer.id}">       
        <div class="row rowplayerinfo">
          <div class="col-1 playercolor" v-bind:style="'background-color:'+player.color">
            {{player.name[0]}}
          </div>
          <div class="col-5">
            <span class="playername">{{player.name}}</span>
            <span v-if="player.isAdmin" class="badge playeradmin">Admin</span>
          </div>
          <div class="col-6 cash">
            €{{player.cashAmount}}
          </div>
        </div>

        <!-- Cash transfer -->
        <div v-if="player.id!=currentPlayer.id" class="row rowplayerinfo">
          <div class="col-6">
            <input class="inputcash" v-bind:id="'idTo-'+player.id" placeholder="€0" type="number" min="0" oninput="validity.valid||(value='');">
          </div>
          <div class="col-6">
            <button class="btnsend" type="button" v-on:click="cashTransfer(player.id)">Send</button>
          </div>
        </div>

        <!-- Add salary -->
        <div v-if="currentPlayer.isAdmin" class="row rowplayerinfo">
          <div class="col-6">
           €2000 
          </div>
          <div class="col-6">
            <button type="button" class="btnsend" v-on:click="salaryAdd(player.id)">Add salary</button><br>
          </div>
        </div>

        <!-- Add/deduct cash -->
        <div v-if="currentPlayer.isAdmin" class="row rowplayerinfo">
          <div class="col-6">
              <input v-bind:id="'cashToId-'+player.id" class="inputcash" placeholder="€0" type="number" min="0" oninput="validity.valid||(value='');"> 
          </div>
          <div class="col-6 btn-group">
            <button type="button" class="btnsend" v-on:click="cashAdd(player.id)">+</button>
            <button type="button" class="btnsend" v-on:click="cashDeduct(player.id)">-</button>
          </div>
        </div>
        
      </div>

      <br/>
      <button type="button" v-on:click="removePlayer()" class="action-button">Leave</button>
    </div>

    <br/>
    <br/>
  </div>
</template>

<script>
import axios from 'axios';

import VueCookies from 'vue-cookies';

const beURL = 'http://' + window.location.hostname + ':30303';
const tokenCookie = 'token';

export default {
  name: 'Tracker',
  data () {
    return {
      validToken: false,
      regName: "",
      playersData: [],
      currentPlayer: {}
    }
  },
  methods: {
    checkToken: function() {
      axios.get(beURL + '/player/'+VueCookies.get(tokenCookie))
      .then(response => {
        this.validToken = response.data.auth.authenticated;
        if (this.validToken === true) {
          this.getPlayers();
        }
      })
      .catch(error => {
        console.log(error);
      })
    },
    playersDataUpdateLP: function() {
      let a = axios.get(beURL + '/lp?timeout=30&category=update-players');
      a.then(response => {
          if (("events" in response.data) && (this.validToken === true)) {
            this.getPlayers();
          }

          this.playersDataUpdateLP();
        })
        .catch(error => {
          console.log(error);
        })
    },
    registerPlayer: function() {
      axios.post(beURL + '/player', {
        name: this.regName
      })
      .then(response => {
        VueCookies.set(tokenCookie, response.data.player.token);
        this.validToken = true;
        this.regName = '';
        this.getPlayers();
      })
      .catch(error => {
        console.log(error);
      })
    },
    removePlayer: function() {
      this.validToken = false;

      axios.delete(beURL + '/player', {
        headers: {
          'X-Token': VueCookies.get(tokenCookie)
        }
      })
      .then(response => {
        VueCookies.remove(tokenCookie)
      })
      .catch(error => {
        console.log(error);
      })
    },
    getPlayers: function() {
      axios.get(beURL + '/players', {
        headers: {
          'X-Token': VueCookies.get(tokenCookie)
        }
      })
      .then(response => {
        this.playersData = response.data.players;
        this.currentPlayer = this.playersData.filter(function(player) {
          return player.you;
        })[0];
      })
      .catch(error => {
        console.log(error);
      })
    },
    cashTransfer: function(idTo) {
      let amount = document.getElementById('idTo-'+idTo).value;
      axios.post(beURL + '/cash/transfer', {
        token: VueCookies.get(tokenCookie),
        idTo: idTo,
        idFrom: this.currentPlayer.id,
        amount: Number(amount)
      })
      .then(response => {
        document.getElementById('idTo-'+idTo).value = '';
      })
      .catch(error => {
        console.log(error);
      })
    },
    salaryAdd: function(idTo) {
      axios.post(beURL + '/salary/add', {
        token: VueCookies.get(tokenCookie),
        id: idTo
      })
      .then(response => {
        // ???
      })
      .catch(error => {
        console.log(error);
      }) 
    },
    cashAdd: function(idTo) {
      let amount = document.getElementById('cashToId-'+idTo).value;
      axios.post(beURL + '/cash/add', {
        token: VueCookies.get(tokenCookie),
        id: idTo,
        amount: Number(amount)
      })
      .then(response => {
        document.getElementById('cashToId-'+idTo).value = '';
      })
      .catch(error => {
        console.log(error);
      }) 
    },
    cashDeduct: function(idTo) {
      let amount = document.getElementById('cashToId-'+idTo).value;
      axios.post(beURL + '/cash/deduct', {
        token: VueCookies.get(tokenCookie),
        id: idTo,
        amount: Number(amount)
      })
      .then(response => {
        document.getElementById('cashToId-'+idTo).value = '';
      })
      .catch(error => {
        console.log(error);
      })   
    }
  },
  created: function(){
    this.checkToken();
    this.playersDataUpdateLP();
  },
  computed: {
    regNameModel: {
        get(){},
        set(newValue){
            this.regName = newValue;
        }
    } 
  }
}
</script>

<style scoped>
.gameinfo {
  background-color: #ffffff;
  border-radius: 12px;
  margin-top: 20px;
  text-align: left;
  padding-top: calc(100vw * 0.01);
  padding-bottom: calc(100vw * 0.03);
}
.bgmain {
  background-color: #008066;
  padding-top: 60px;
}
.row.no-gutter [class*='col-']:not(:first-child) {
  padding-left: 0;
  padding-left: 0;
}

.current-user-name {
  font-weight: bold;  
  font-size: calc(100vw * 0.07);
}
.cash {
  font-size: calc(100vw * 0.05);
  color:#008066;
}
.playername{
  font-weight: bold;
}
.playerinfo {
  background-color:#e6e6e6;
  border-radius: 12px;
  margin-top: 10px;
  margin-bottom: 20px;
  padding-bottom: 10px;
  padding-top: 10px;
  text-align: left;
  font-size: calc(100vw * 0.05);
}
.rowplayerinfo {
  margin-bottom: 10px;
}
.current-user-list{
  border: 3px solid #008066;    
}
.inputcash {
  width: 100%;
  border-color: #008066;
  border-radius: 6px;
}
input:focus{
  background-color: #eed7f4;
}
.btnsend {
  width: 100%;
  height: 100%;
  background-color: #008066;
  color: #ffffff;
  font-weight: bold;
  border-radius: 6px;
}
.btnsend:active{
  transform: translateY(4px);
}
.playeradmin{
  color: #008066;
}
.playerpoint{
  font-size: calc(100vw * 0.05);
  font-weight: bold;
}
.playercolor{
  border-radius: 50%;
  color: #e6e6e6;
  text-align: center;
  font-weight: bold;
  margin-left: calc(100vw * 0.02)
}
.action-button {
  background-color: #008066;
  color: #FFF;
  text-align: center;
  font-size: calc(100vw * 0.08);
  font-weight: bold;
  border-radius: 12px;
  width: 100%;
  height: calc(100vw * 0.2);
}
.action-button:active{
  transform: translateY(4px); 
}
.name-input {
  width: 100%;
  border: 1px solid #008066;
  border-radius: 12px;
  font-size: calc(100vw * 0.08);
  height: calc(100vw * 0.2);
}
</style>
