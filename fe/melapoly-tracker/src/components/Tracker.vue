<template>
  <div class="container-fluid">
    <img class="img-fluid" src="https://raw.githubusercontent.com/alex-ant/melapoly-tracker/master/fe/melapoly-tracker/img/logo.png" />

    <br/><br/>

    <div v-if="!validToken">
      <input placeholder="Enter your name" class="name-input" type="text" v-model="regNameModel">
      <br/><br/>
      <button type="button" v-on:click="registerPlayer()" class="action-button">MOVE</button>
    </div>

    <div v-else>
      <div> 
        Hello, {{currentPlayer.name}}!
        €{{currentPlayer.cashAmount}}
      </div>

      <table>
        <tr>
          <td>Name</td>
          <td>Cash</td>
          <td>Admin</td>
        </tr>
        <tr v-for="player in playersData" v-bind:key="player.id" v-bind:class="{'current-user': player.id ===  currentPlayer.id}">
          <td>{{player.name}}</td>
          <td>{{player.cashAmount}}</td>
          <td>{{player.isAdmin}}</td>
        </tr>
      </table>

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
.current-user {
  font-weight: bold;
}
.action-button {
  background-color: #217867;
  border: 1px solid #000;
  color: #FFF;
  text-align: center;
  display: inline-block;
  font-size: calc(100vw * 0.08);
  border-radius: 12px;
  text-shadow: -1px 0 black, 0 1px black, 1px 0 black, 0 -1px black;
  width: 100%;
  height: calc(100vw * 0.2);
}
.name-input {
  width: 100%;
  border: 1px solid #000;
  border-radius: 12px;
  font-size: calc(100vw * 0.08);
  height: calc(100vw * 0.2);
}
</style>
