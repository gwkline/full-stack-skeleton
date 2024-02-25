import { initGraphQLTada } from "gql.tada";
import type { introspection } from "./graphql-env.js";

export const graphql = initGraphQLTada<{
	introspection: introspection;
	scalars: {
		ID: string;
		Time: Date;
	};
}>();

export type { FragmentOf, ResultOf, VariablesOf } from "gql.tada";
export { readFragment } from "gql.tada";
