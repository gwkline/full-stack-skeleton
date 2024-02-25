import { DataTableColumnHeader } from "@/components/table/columnHeader";
import { Skeleton } from "@/components/ui/skeleton";
import { ColumnDef } from "@tanstack/react-table";

interface ModelEdge {
	node: {
		id: string;
	};
}

export function IdCell<T extends ModelEdge>(isLoading: boolean): ColumnDef<T> {
	return {
		id: "id",
		accessorKey: "node.id",
		enableHiding: true,
		enableSorting: true,
		header: ({ column }) => (
			<DataTableColumnHeader column={column} title="ID" isLoading={isLoading} />
		),
		cell: ({ row }) =>
			isLoading ? (
				<Skeleton className="h-8 w-full" />
			) : (
				<div className="w-[55px] overflow-hidden overflow-ellipsis  whitespace-nowrap">
					{row.original.node.id}
				</div>
			),
	};
}
