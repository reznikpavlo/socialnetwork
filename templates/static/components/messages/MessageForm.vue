<template>
  <v-layout row>
    <v-text-field
        label="New message"
        placeholder="Write something"
        v-model="text"
    />
    <v-btn @click="save">
      Save
    </v-btn>
  </v-layout>

</template>

<script>
  import axios from "axios";

  function getIndex (id, list) {
    for (let i = 0; i < list.length; i++) {
      if (id === list[i].id) {
        return i
      }
    }
    return -1
  }

  export default {

    props: ['messages', 'message'],
    watch: {
      message:function (n, o) {
        this.text = n.text
        this.id = n.id
      }
    },
    data () {
      return {
        text: '',
        id : ''
      }
    },
    methods: {
      save: function () {

        if (this.id) {
          var message = {id: this.id, text: this.text};
          axios.put("/message/"+ this.id, message).then(response => {
            var updatedMessage = response.data
            let index = getIndex(updatedMessage.id, this.messages);
            this.messages.splice(index, 1, updatedMessage)
          })
        }
        else {
          let message = {text: this.text};
          axios.post("/message/", message).then(response => {
            this.messages.push(response.data)

          })
        }
        this.id = ''
        this.text = ''


      }
    }
  }

</script>

<style>


</style>