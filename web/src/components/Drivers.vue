<template>
  <form>
    Year:
    <input v-model="searchYear" placeholder="current" />
    <button
      class="btn btn-outline-primary"
      v-on:click="searchByYear"
      type="button"
    >
      Search
    </button>
  </form>
  <div class="apollo">
    <p>{{ driversQuery }}</p>
    <p
      v-for="driver in driversQuery.DriverStandings?.drivers"
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
// import { useQuery } from "@vue/apollo-composable";

const DRIVERS_QUERY = gql`
  query Drivers($year: String!) {
    DriverStandings(filter: { year: $year }) {
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
  apollo: {
    driversQuery: {
      query: DRIVERS_QUERY,
      variables() {
        return {
          year: this.year,
        };
      },
    },
  },
  data() {
    return {
      driversQuery: [],
      searchYear: "",
      year: "current",
    };
  },
  methods: {
    searchByYear() {
      this.year = this.searchYear;
    },
  },
};
</script>
