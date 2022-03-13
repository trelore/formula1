<template>
  <div class="apollo">
    Year:
    <input
      v-model="searchYear"
      @keyup.enter="searchByYear"
      placeholder="current"
    />
    <button @click="searchByYear" :disabled="!searchYear" type="button">
      Search
    </button>
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
      console.log(this);
      console.log(this.searchYear);
      this.$apollo.queries.driversQuery.refresh();
    },
  },
};
</script>
