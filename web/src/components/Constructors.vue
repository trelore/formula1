<template>
  Year:
  <input v-model.lazy="variables.year" placeholder="current" />
  <div id="constructors">
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
  setup() {
    const { result, loading, error, refetch, variables } = useQuery(
      CONSTRUCTORS_QUERY,
      { year: "current" }
    );
    return {
      result,
      loading,
      error,
      refetch,
      variables,
    };
  },
};
</script>
