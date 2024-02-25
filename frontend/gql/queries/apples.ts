import { graphql } from "@/gql/graphql";
import { keepPreviousData, queryOptions } from "@tanstack/react-query";
import { Cookies } from "next-client-cookies";
import { makeGqlWithInput } from "../base";

type QueryParams = {
	userId: string | undefined;
};

export const applesFrag = graphql(`
fragment ApplesFrag on Apple @_unmask {
	id
	keyword
}`);

const applesQuery = graphql(
	/* GraphQL */ `
  query apples($input: AppleQueryInput!) {
	viewer {
		apples(input: $input) {
            ...ApplesFrag
		}
	}
  }
`,
	[applesFrag],
);

export const Apples = (queryParams: QueryParams, cookies: Cookies) => {
	return queryOptions({
		queryKey: ["Apples", queryParams],
		queryFn: () =>
			makeGqlWithInput({
				document: applesQuery,
				queryParams,
				accessorFn: (response) => response.viewer.apples,
				cookies,
			}),
		retry: 1,
		placeholderData: keepPreviousData,
	});
};
