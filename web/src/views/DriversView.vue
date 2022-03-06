<template>
  <div class="about">
    <h1>Drivers standings</h1>
  </div>
  <div class="apollo">
    <p v-if="error">Something went wrong...</p>
    <p v-if="loading">Loading...</p>
    <p
      v-else
      v-for="driver in result.DriverStandings.drivers"
      :key="driver.Driver.code"
    >
      {{ driver.points }}
      <a v-bind:href="driver.Driver.url"
        >{{ driver.Driver.familyName.toUpperCase() }},
        {{ driver.Driver.givenName }}</a
      >
    </p>
  </div>
</template>

<script>
import gql from "graphql-tag";
import { useQuery } from "@vue/apollo-composable";

const DRIVERS_QUERY = gql`
  query Drivers {
    DriverStandings(filter: { year: "2018", top: 5 }) {
      drivers {
        points
        Driver {
          code
          givenName
          familyName
          url
        }
      }
    }
  }
`;

export default {
  name: "App",
  setup() {
    const { result, loading, error } = useQuery(DRIVERS_QUERY);
    return {
      result,
      loading,
      error,
    };
  },
};
</script>
