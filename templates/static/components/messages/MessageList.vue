<template>
  <v-layout align-space-around justify-start column>
    <message-form :messages="messages" :message="message"/>
    <message-row v-for="message in messages"
                 :key="message.id"
                 :message="message"
                 :editMessage="editMessage"/>
  </v-layout>

</template>

<script>
import MessageRow from "./MessageRow.vue"
import MessageForm from "./MessageForm.vue"
import axios from 'axios'
import Cookes from "js-cookie"

export default {
  components: {
    MessageRow, MessageForm
  },
  props: ['messages'],
  data: function () {
    return {
      message: null
    }
  },
  methods: {
    async getMessages() {
      let token = Cookes.get('sessionid')
      if (token) {
        const {data} = await axios.get("/message/")
        data.forEach(message => this.messages.push(message))
      }
    },

    editMessage(message) {
      this.message = message

    }
  },
  created: function () {
    this.getMessages();
  },
}


</script>

<style>

</style>