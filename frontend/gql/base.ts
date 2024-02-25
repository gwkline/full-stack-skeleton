import { isServer } from "@tanstack/react-query";
import { TadaDocumentNode } from "gql.tada";
import {
	GraphQLClient,
	RequestExtendedOptions,
	Variables,
} from "graphql-request";
import { Cookies } from "next-client-cookies";

export async function makeGqlQuery<TQuery, TResult>({
	document,
	accessorFn,
	cookies,
}: {
	document: TadaDocumentNode<TQuery, { [k: string]: never }, void>;
	accessorFn: (response: TQuery) => TResult;
	cookies: Cookies;
}) {
	let url = process.env.NEXT_PUBLIC_GRAPHQL_URL;
	if (!url) {
		throw new Error("NEXT_PUBLIC_GRAPHQL_URL is not set");
	}

	if (isServer && process.env.NODE_ENV !== "production") {
		url = "http://backend:8888/graphql";
	}

	const client = new GraphQLClient(url, {
		fetch: fetch,
		headers: {
			Authorization: `Bearer ${cookies.get("accessToken")}`,
		},
	});

	const response = await client.request(document, {});
	return accessorFn(response);
}

export async function makeGqlWithInput<
	TQuery,
	TInput extends Variables,
	TResult,
>({
	document,
	queryParams,
	accessorFn,
	cookies,
}: {
	document: TadaDocumentNode<TQuery, { input: TInput }, void>;
	queryParams: TInput;
	accessorFn: (response: TQuery) => TResult;
	cookies: Cookies;
}) {
	let url = process.env.NEXT_PUBLIC_GRAPHQL_URL;
	if (!url) {
		throw new Error("NEXT_PUBLIC_GRAPHQL_URL is not set");
	}
	if (isServer && process.env.NODE_ENV !== "production") {
		url = "http://backend:8888/graphql";
	}

	const variables: Variables = { input: queryParams };
	const options: RequestExtendedOptions = {
		url,
		document,
		variables,
	};

	const client = new GraphQLClient(url, {
		headers: {
			Authorization: `Bearer ${cookies.get("accessToken")}`,
		},
		fetch,
	});

	const response = await client.request(options.document, options.variables);
	return accessorFn(response as Awaited<TQuery>);
}
