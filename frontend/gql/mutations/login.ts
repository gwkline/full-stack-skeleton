import { toast } from "@/components/ui/use-toast";
import { ResultOf, graphql } from "@/gql/graphql";
import { QueryClient } from "@tanstack/react-query";
import { Cookies } from "next-client-cookies";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";
import { makeGqlWithInput } from "../base";

interface QueryParams {
	email: string;
	password: string;
}

const loginMutation = graphql(/* GraphQL */ `
  mutation login($input: LoginInput!) {
    login(input: $input) {
	accessToken
	refreshToken
	}
  }
`);

export const LoginMutation = (
	queryClient: QueryClient,
	router: AppRouterInstance,
	cookies: Cookies,
) => {
	return {
		mutationKey: ["Login"],
		mutationFn: (queryParams: QueryParams) =>
			makeGqlWithInput({
				document: loginMutation,
				queryParams,
				accessorFn: (response) => response,
				cookies,
			}),
		onSuccess: (data: ResultOf<typeof loginMutation>) => {
			if (data) {
				cookies.set("accessToken", data.login.accessToken);
				cookies.set("login", data.login.refreshToken);
			}
			toast({
				variant: "default",
				title: "Success",
				description: "Successfully logged in",
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
			cookies.remove("login");
			queryClient.invalidateQueries({ queryKey: ["Viewer"] });
			router.push("/");
		},
	};
};
