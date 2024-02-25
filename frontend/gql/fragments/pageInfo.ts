import { graphql } from "@/gql/graphql";

export const pageInfoFrag = graphql(`
fragment PageInfoFrag on PageInfo @_unmask {
    startCursor
    endCursor
    hasNextPage
    count
  }
`);
