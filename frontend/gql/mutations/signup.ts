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

const signupMutation = graphql(/* GraphQL */ `
  mutation signup($input: UserInput!) {
    signup(input: $input) {
	accessToken
	refreshToken
	}
  }
`);

export const SignupMutation = (
	queryClient: QueryClient,
	router: AppRouterInstance,
	cookies: Cookies,
) => {
	return {
		mutationKey: ["Signup"],
		mutationFn: (queryParams: QueryParams) =>
			makeGqlWithInput({
				document: signupMutation,
				queryParams,
				accessorFn: (response) => response,
				cookies,
			}),
		onSuccess: (data: ResultOf<typeof signupMutation>) => {
			if (data) {
				cookies.set("accessToken", data.signup.accessToken);
				cookies.set("signup", data.signup.refreshToken);
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
			cookies.remove("signup");
			queryClient.invalidateQueries({ queryKey: ["Viewer"] });
			router.push("/");
		},
	};
};
