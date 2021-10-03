<template>
  <div>
    <v-btn @click="getUsers" color="primary" elevation="2" x-large x-small
      >Get Users</v-btn
    >
    <v-data-table
      :headers="headers"
      :items="users"
      :items-per-page="10"
      class="elevation-1"
    ></v-data-table>
  </div>
</template>
<script>
export default {
  data() {
    return {
      headers: [
        {
          text: "Name",
          align: "start",
          sortable: false,
          value: "name",
        },
        { text: "Email", value: "email" },
      ],
      users: [],
    };
  },
  mounted() {
    this.getUsers();
  },
  methods: {
    getUsers() {
      this.$axios
        .get("api/v1/user/")
        .then((res) => {
          this.users = res.data;
        })
    },
  },
};
</script>