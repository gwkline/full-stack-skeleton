import { Skeleton } from "@/components/ui/skeleton";
import { ColumnDef } from "@tanstack/react-table";
import { DataTableColumnHeader } from "../columnHeader";

export function BasicCell<T>({
	id,
	title,
	width,
	isLoading,
	accessorKey,
	accessorFn,
}: {
	id: string;
	title: string;
	width: string;
	isLoading: boolean;
	accessorKey: string;
	accessorFn: (edge: T) => string;
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
				<div
					className={`w-${width} overflow-hidden overflow-ellipsis whitespace-nowrap`}
				>
					{accessorFn(row.original)}
				</div>
			),
		enableSorting: true,
		enableHiding: true,
	};
}
