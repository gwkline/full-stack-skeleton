import {
	HoverCard,
	HoverCardContent,
	HoverCardTrigger,
} from "@/components/ui/hover-card";
import { Skeleton } from "@/components/ui/skeleton";
import { ColumnDef } from "@tanstack/react-table";
import { ReactNode } from "react";
import { DataTableColumnHeader } from "../columnHeader";

export function HoverCell<T>({
	id,
	title,
	isLoading,
	accessorKey,
	accessorFn,
	innerAccessorFn,
}: {
	id: string;
	title: string;
	isLoading: boolean;
	accessorKey: string;
	accessorFn: (edge: T) => string;
	innerAccessorFn: (edge: T) => ReactNode;
}): ColumnDef<T> {
	return {
		id: id,
		accessorKey: accessorKey,
		header: ({ column }) => (
			<DataTableColumnHeader
				column={column}
				title={title}
				isLoading={isLoading}
			/>
		),
		cell: ({ row }) =>
			isLoading ? (
				<Skeleton className="h-8 w-full" />
			) : (
				<div className="w-[150px] overflow-hidden overflow-ellipsis  whitespace-nowrap">
					<HoverCard>
						<HoverCardTrigger>{accessorFn(row.original)}</HoverCardTrigger>
						<HoverCardContent
							align="start"
							side="bottom"
							className="max-w-[300px] max-h-[600px] overflow-y-auto whitespace-normal break-words"
						>
							{innerAccessorFn(row.original)}
						</HoverCardContent>
					</HoverCard>
				</div>
			),
		enableSorting: true,
		enableHiding: true,
	};
}
