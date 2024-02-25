import { Skeleton } from "@/components/ui/skeleton";
import { ColumnDef } from "@tanstack/react-table";
import { DataTableColumnHeader } from "../columnHeader";

export function IconCell<T>({
	id,
	title,
	isLoading,
	accessorKey,
	accessorFn,
}: {
	id: string;
	title: string;
	isLoading: boolean;
	accessorKey: string;
	accessorFn: (edge: T) => JSX.Element;
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
			isLoading ? <Skeleton className="h-8 w-8" /> : accessorFn(row.original),
		enableSorting: true,
		enableHiding: true,
	};
}
