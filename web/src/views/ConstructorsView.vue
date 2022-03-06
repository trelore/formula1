<template>
  <div class="about">
    <h1>Constructors standings</h1>
  </div>
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
  query Constructors {
    ConstructorStandings(filter: { year: "2018" }) {
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
  name: "App",
  setup() {
    const { result, loading, error } = useQuery(CONSTRUCTORS_QUERY);
    return {
      result,
      loading,
      error,
    };
  },
};
</script>
