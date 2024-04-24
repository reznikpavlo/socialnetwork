<template>
<!--  <div>
    <i>{{message.id}}</i>. {{message.text}}
    <span style="position: absolute; right: 0">
                <input type="button" value = "Edit" @click="edit"/>
      &lt;!&ndash;input type="button" value="Delete" @click="delete"/&ndash;&gt;
            </span>
  </div>-->

  <v-card class="my-2">
    <v-card-text primary-title>
      <i>({{ message.id }})</i>
      {{ message.text }}
    </v-card-text>
    <v-card-actions>
      <v-btn value="Edit" @click="edit" small flat round>Edit</v-btn>
    </v-card-actions>
  </v-card>

</template>

<script>
  export default {
    props: ['message', 'editMessage'],
    methods: {
      edit: function () {this.editMessage(this.message)},
      delete: function () {}
    }
  }

</script>

<style>

</style>