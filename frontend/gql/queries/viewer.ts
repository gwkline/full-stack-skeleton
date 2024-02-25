import { makeGqlQuery } from "@/gql/base";
import { graphql } from "@/gql/graphql";
import { keepPreviousData, queryOptions } from "@tanstack/react-query";
import { Cookies } from "next-client-cookies";

export const viewerFrag = graphql(`
fragment ViewerFrag on Viewer @_unmask {
	id
	email
	role
}`);
const viewerQuery = graphql(
	/* GraphQL */ `query viewer {
	viewer {
		...ViewerFrag
	}
  }
`,
	[viewerFrag],
);

export const Viewer = (cookies: Cookies) => {
	return queryOptions({
		queryKey: ["Viewer"],
		queryFn: () =>
			makeGqlQuery({
				document: viewerQuery,
				accessorFn: (response) => response.viewer,
				cookies,
			}),
		retry: 1,
		placeholderData: keepPreviousData,
	});
};
