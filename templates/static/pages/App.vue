<template>
<!--  <v-app >
    <v-app-bar color="surface-variant" app>
      <v-toolbar-title>Social Network without politics</v-toolbar-title>
      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-btn href="/google/login">
        <v-icon>mdi-login</v-icon>
      </v-btn>

    </v-app-bar>

    <v-main>
      <v-container>Hello user</v-container>
      <v-container>
        <messages-list :messages="messages"/>
      </v-container>
    </v-main>


  </v-app>-->

  <v-app>
    <v-toolbar app>
      <v-toolbar-title>Social Network</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon href="/logout">
        <v-icon>exit_to_app</v-icon>
      </v-btn>
    </v-toolbar>
    <v-content>
        You should login
        <a href="/google/login">Google</a>
        <messages-list :messages="messages" />
    </v-content>
  </v-app>


</template>

<script>

  import MessagesList from "../components/messages/MessageList.vue";
  export default {
    components: {
      MessagesList
    },
    data() {
      return {
        messages: []
      }
    }

  }
</script>

<style>


</style>