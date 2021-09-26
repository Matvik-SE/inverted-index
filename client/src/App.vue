<template>
  <div class="container mt-5">
    <div class="d-flex justify-content-between">
      <div class="btn-group">
        <div class="input-group mx-3">
          <div class="input-group-prepend">
            <span class="input-group-text">Threads:</span>
          </div>
          <input v-model="inputThreads" type="text" class="form-control">
        </div>
        <button @click="threads = inputThreads" class="form-control bg-primary text-white">
          Set
        </button>
      </div>

      <div>
        <input :value="`Results: ${results.length}`" type="text" class="form-control" disabled>
      </div>

      <div class="btn-group">
        <div class="input-group">
          <div class="input-group-prepend">
            <span class="input-group-text">Search:</span>
          </div>
          <input v-model="inputSearch" type="text" class="form-control">
        </div>
        <button @click="searchRequest" class="form-control bg-primary text-white mx-3">
          Build Index
        </button>
      </div>
    </div>

    <div class="text-center mt-3">
      <table class="table w-75 m-auto">
        <thead>
        <tr>
          <th>#</th>
          <th>File Path</th>
          <th>Occurrences</th>
        </tr>
        </thead>
        <tbody>
        <tr v-if="results.length === 0">
          <td colspan="4" class="text-center">No results</td>
        </tr>
        <tr v-for="(result, idx) of results" :key="idx">
          <td>{{ idx + 1 }}</td>
          <td>{{ result.split('-')[0] }}</td>
          <td>{{ result.split('-')[1] }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { ref, watch, onMounted } from 'vue';
import axios from 'axios';

function apiSetThreads(value) {
  axios.get(`/threads?number=${value}`);
}

export default {
  name: 'App',
  setup() {
    const threads = ref(1);
    const results = ref([]);
    const inputThreads = ref(1);
    const inputSearch = ref('');

    const connection = new WebSocket('ws://localhost:3000/socket');
    connection.onmessage = (event) => {
      results.value.push(event.data);
    };

    onMounted(() => {
      apiSetThreads(threads.value);
    });

    watch(threads, apiSetThreads);

    function searchRequest() {
      results.value = [];
      connection.send(inputSearch.value);
    }

    return {
      threads,
      results,
      inputThreads,
      inputSearch,
      searchRequest,
    };
  },
};
</script>

<style scoped></style>
