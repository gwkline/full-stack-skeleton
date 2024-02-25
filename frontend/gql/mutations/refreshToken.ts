import { toast } from "@/components/ui/use-toast";
import { ResultOf, graphql } from "@/gql/graphql";
import { QueryClient } from "@tanstack/react-query";
import { Cookies } from "next-client-cookies";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";
import { makeGqlWithInput } from "../base";

interface QueryParams {
	refreshToken: string;
	accessToken: string;
}

const refreshTokenMutation = graphql(/* GraphQL */ `
  mutation refreshToken($input: RefreshInput!) {
    refreshToken(input: $input) {
	accessToken
	refreshToken
	}

  }
`);

export const RefreshToken = (
	queryClient: QueryClient,
	router: AppRouterInstance,
	cookies: Cookies,
) => {
	return {
		mutationKey: ["RefreshToken"],
		mutationFn: (queryParams: QueryParams) =>
			makeGqlWithInput({
				document: refreshTokenMutation,
				queryParams,
				accessorFn: (response) => response,
				cookies,
			}),
		onSuccess: (data: ResultOf<typeof refreshTokenMutation>) => {
			if (data) {
				cookies.set("accessToken", data.refreshToken.accessToken);
				cookies.set("refreshToken", data.refreshToken.refreshToken);
			}
			toast({
				variant: "default",
				title: "Success",
				description: "Auth Token Refreshed",
			});
		},
		onError: (error: Error) => {
			console.log(error.message);
			toast({
				variant: "destructive",
				title: "Refresh Token Error",
				description: error.message.split(`: {"response":`)[0],
			});
			cookies.remove("accessToken");
			cookies.remove("refreshToken");
			queryClient.invalidateQueries({ queryKey: ["Viewer"] });
			router.push("/");
		},
	};
};
