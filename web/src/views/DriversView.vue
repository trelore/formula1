<template>
  <div class="about">
    <h1>Drivers standings</h1>
  </div>
  <Drivers />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Drivers from "@/components/Drivers.vue";

export default defineComponent({
  name: "ConstructorsView",
  components: {
    Drivers,
  },
});
</script>
