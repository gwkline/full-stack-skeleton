import { QueryParams } from "@/components/table/dataTable";
import { graphql } from "@/gql/graphql";
import { keepPreviousData, queryOptions } from "@tanstack/react-query";
import { Cookies } from "next-client-cookies";
import { makeGqlWithInput } from "../base";
import { pageInfoFrag } from "../fragments/pageInfo";

export const appleEdge = graphql(
	/* GraphQL */ `
fragment AppleEdge on AppleEdge @_unmask {
  cursor
  node {
    id
    variety
  }
}`,
);

const appleConnectionQuery = graphql(
	/* GraphQL */ `
  query appleConnection($input: ConnectionInput!) {
    viewer {
      appleConnection(input: $input) {
        pageInfo {
          ...PageInfoFrag
        }
        edges {
          ...AppleEdge
        }
      }
    }
  }
`,
	[pageInfoFrag, appleEdge],
);

export const AppleConnection = (queryParams: QueryParams, cookies: Cookies) => {
	return queryOptions({
		queryKey: ["AppleConnection", queryParams],
		queryFn: () =>
			makeGqlWithInput({
				document: appleConnectionQuery,
				queryParams,
				accessorFn: (response) => {
					return response.viewer.appleConnection;
				},
				cookies,
			}),
		retry: 1,
		placeholderData: keepPreviousData,
	});
};
