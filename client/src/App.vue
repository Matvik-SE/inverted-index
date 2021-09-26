<template>
  <div>foo-bar: threads</div>
</template>

<script>
import { ref } from 'vue';

export default {
  name: 'App',
  setup() {
    const threads = ref(1);

    return {
      threads,
    };
  },
};
</script>

<style scoped></style>
