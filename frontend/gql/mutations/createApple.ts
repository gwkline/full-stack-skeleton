import { toast } from "@/components/ui/use-toast";
import { graphql } from "@/gql/graphql";
import type { AppleVariety } from "@/lib/types";
import { QueryClient } from "@tanstack/react-query";
import { Cookies } from "next-client-cookies";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";
import { makeGqlWithInput } from "../base";

interface QueryParams {
	variety: AppleVariety;
	userId: string;
	quantity: number;
}

const createAppleMutation = graphql(/* GraphQL */ `
  mutation createApple($input: AppleCreateInput!) {
    createApple(input: $input) {
        id
        variety
	}
  }
`);

export const CreateApple = (
	queryClient: QueryClient,
	cookies: Cookies,
	router: AppRouterInstance,
) => {
	return {
		mutationKey: ["RefreshToken"],
		mutationFn: (queryParams: QueryParams) =>
			makeGqlWithInput({
				document: createAppleMutation,
				queryParams,
				accessorFn: (response) => response,
				cookies,
			}),
		onSuccess: () => {
			toast({
				variant: "default",
				title: "Success",
				description: "Apple Created",
			});
			queryClient.invalidateQueries({ queryKey: ["AppleConnection"] });
			router.push("/apples");
		},
		onError: (error: Error) => {
			console.log(error.message);
			toast({
				variant: "destructive",
				title: "Error",
				description: error.message.split(`: {"response":`)[0],
			});
			queryClient.invalidateQueries({ queryKey: ["AppleConnection"] });
		},
	};
};
