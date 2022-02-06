<template>
  <div class="container">
    <div class="col-12">
      <div class="form-group">
        <label for="textInput">Type new task</label>
        <input type="text" class="form-control" id="textInput" v-model="text" />
        <p />
        <!-- <small id="textHelp" class="form-text text-muted"> a</small> -->
      </div>
      <button type="submit" class="btn btn-primary" @click="CreateItem">
        Add
      </button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "TodoCreator",
  data: function () {
    return {
      text: "",
    };
  },
  methods: {
    async CreateItem() {
      await axios.post("http://localhost:8081/items", { text: this.text });
      this.$emit("ItemCreated");
    },
  },
};
</script>

<style scoped>
</style>