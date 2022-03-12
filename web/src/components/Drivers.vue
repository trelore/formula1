<template>
  <div class="apollo">
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
  name: "Drivers-Component",
  apollo: {
    driversQuery: {
      query: DRIVERS_QUERY,
      variables() {
        return { year: this.searchYear };
      },
    },
  },
  data() {
    return {
      driversQuery: [],
      searchYear: "",
    };
  },
  methods: {
    searchByYear() {
      this.$apollo.queries.driversQuery.refetch(this.searchYear);
    },
  },
};
</script>
