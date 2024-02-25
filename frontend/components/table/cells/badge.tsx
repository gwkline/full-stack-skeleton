import { Badge, badgeVariants } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { ColumnDef } from "@tanstack/react-table";
import { DataTableColumnHeader } from "../columnHeader";

export function BadgeCell<T>({
	id,
	title,
	isLoading,
	accessorKey,
	accessorFn,
	variantAccessorFn,
}: {
	id: string;
	title: string;
	isLoading: boolean;
	accessorKey: string;
	accessorFn: (edge: T) => string;
	variantAccessorFn?: (edge: T) => keyof typeof badgeVariants;
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
				<Badge
					variant={
						variantAccessorFn ? variantAccessorFn(row.original) : "default"
					}
				>
					{accessorFn(row.original)}
				</Badge>
			),
		enableSorting: false,
		enableHiding: true,
	};
}
