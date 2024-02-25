"use client";
import { DataTableRowActions as AppleActions } from "@/app/apples/actions";
import { AuthContext } from "@/components/providers/authClient";
import { Button } from "@/components/ui/button";
import {
	DropdownMenu,
	DropdownMenuContent,
	DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Skeleton } from "@/components/ui/skeleton";
import { appleEdge } from "@/gql/queries/appleConnection";
import { DotsHorizontalIcon } from "@radix-ui/react-icons";
import { ColumnDef, Row } from "@tanstack/react-table";
import { ResultOf } from "gql.tada";
import { usePathname } from "next/navigation";
import { useContext } from "react";

export function ActionsCell<T>({
	isLoading,
}: { isLoading: boolean }): ColumnDef<T> {
	const path = usePathname();
	const { user } = useContext(AuthContext);

	return {
		id: "actions",
		enableSorting: false,
		enableColumnFilter: false,
		enableHiding: false,
		header: () => (
			<Button
				variant="ghost"
				size="sm"
				className="h-8 w-full flex flex-row justify-center"
			>
				<span>Actions</span>
			</Button>
		),
		cell: ({ row }) => {
			if (isLoading) return <Skeleton className="h-8 w-full" />;
			return (
				<DropdownMenu>
					<DropdownMenuTrigger asChild>
						<Button
							variant="outline"
							className="flex h-8 w-8 p-0 bg-background"
						>
							<DotsHorizontalIcon className="h-4 w-4" />
							<span className="sr-only">Open menu</span>
						</Button>
					</DropdownMenuTrigger>
					<DropdownMenuContent align="end" className="w-[160px]">
						{getActionsCell(row, path)}
					</DropdownMenuContent>
				</DropdownMenu>
			);
		},
	};
}

const getActionsCell = <T,>(row: Row<T>, path: string) => {
	if (path.includes("/apples")) {
		return (
			<AppleActions row={row as unknown as Row<ResultOf<typeof appleEdge>>} />
		);
	}
};
