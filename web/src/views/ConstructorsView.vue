<template>
  <div class="about">
    <h1>Constructors Standings</h1>
  </div>
  <Constructors />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Constructors from "@/components/Constructors.vue";

export default defineComponent({
  name: "ConstructorsView",
  components: {
    Constructors,
  },
});
</script>
