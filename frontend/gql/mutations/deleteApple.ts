import { toast } from "@/components/ui/use-toast";
import { graphql } from "@/gql/graphql";
import { QueryClient } from "@tanstack/react-query";
import { Cookies } from "next-client-cookies";
import { makeGqlWithInput } from "../base";

interface QueryParams {
	ids: string[];
}

const deleteAppleMutation = graphql(/* GraphQL */ `
  mutation deleteApple($input: DeleteInput!) {
    deleteApple(input: $input)
  }
`);

export const DeleteApple = (queryClient: QueryClient, cookies: Cookies) => {
	return {
		mutationKey: ["RefreshToken"],
		mutationFn: (queryParams: QueryParams) =>
			makeGqlWithInput({
				document: deleteAppleMutation,
				queryParams,
				accessorFn: (response) => response,
				cookies,
			}),
		onSuccess: () => {
			toast({
				variant: "default",
				title: "Success",
				description: "Apple Deleted",
			});
			queryClient.invalidateQueries({ queryKey: ["AppleConnection"] });
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
