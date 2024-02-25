"use client";
import {
	DropdownMenuItem,
	DropdownMenuShortcut,
} from "@/components/ui/dropdown-menu";
import { DeleteApple } from "@/gql/mutations/deleteApple";
import { appleEdge } from "@/gql/queries/appleConnection";
import { DataTableRowActionsProps } from "@/lib/types";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { ResultOf } from "gql.tada";
import { useCookies } from "next-client-cookies";

export function DataTableRowActions({
	row,
}: DataTableRowActionsProps<ResultOf<typeof appleEdge>>) {
	const queryClient = useQueryClient();

	const cookies = useCookies();
	const { mutate } = useMutation(DeleteApple(queryClient, cookies));

	return (
		<>
			<DropdownMenuItem onClick={() => mutate({ ids: [row.original.node.id] })}>
				Delete
				<DropdownMenuShortcut>⌘⌫</DropdownMenuShortcut>
			</DropdownMenuItem>
		</>
	);
}
