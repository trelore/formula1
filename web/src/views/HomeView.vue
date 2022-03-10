<template>
  <div class="home">
    <h1>Standings</h1>
    <Drivers />
    <Constructors />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Constructors from "@/components/Constructors.vue";
import Drivers from "@/components/Drivers.vue";

export default defineComponent({
  name: "HomeView",
  components: {
    Drivers,
    Constructors,
  },
});
</script>
