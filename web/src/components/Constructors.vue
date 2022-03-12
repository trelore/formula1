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
    <p v-if="error">Something went wrong...</p>
    <p v-if="loading">Loading...</p>
    <p
      v-else
      v-for="team in result.ConstructorStandings.teams"
      :key="team.team.id"
    >
      {{ team.points }} <a v-bind:href="team.team.url">{{ team.team.name }}</a>
    </p>
  </div>
</template>

<script>
import gql from "graphql-tag";
import { useQuery } from "@vue/apollo-composable";

const CONSTRUCTORS_QUERY = gql`
  query Constructors($year: String!) {
    ConstructorStandings(filter: { year: $year }) {
      teams {
        points
        team {
          id
          name
          url
        }
      }
    }
  }
`;

export default {
  name: "Constructors-Component",
  data() {
    this.searchYear = "current";
    console.log(this.searchYear);
    const variables = { year: this.searchYear };
    const { result, loading, error } = useQuery(CONSTRUCTORS_QUERY, variables);
    return {
      result,
      loading,
      error,
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
