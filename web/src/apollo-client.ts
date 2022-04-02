import { ApolloClient, InMemoryCache, HttpLink } from "@apollo/client/core";
import { onError } from "@apollo/client/link/error";
import { logErrorMessages } from "@vue/apollo-util";

// Create an http link:
const httpLink = new HttpLink({
  uri: "http://localhost:8080/query",
  fetch: (uri: RequestInfo, options: RequestInit) => {
    return fetch(uri, options);
  },
});

const errorLink = onError((error: any) => {
  if (process.env.NODE_ENV !== "production") {
    logErrorMessages(error);
  }
});

// Create the apollo client
export const apolloClient = new ApolloClient({
  cache: new InMemoryCache(),
  link: errorLink.concat(httpLink),
});
