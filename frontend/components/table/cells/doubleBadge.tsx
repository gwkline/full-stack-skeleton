import { Badge, badgeVariants } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { ColumnDef } from "@tanstack/react-table";
import { DataTableColumnHeader } from "../columnHeader";

export function DoubleBadgeCell<T>({
	id,
	title,
	isLoading,
	accessorKey,
	accessorFnA,
	accessorFnB,
	variantAccessorFnA,
	variantAccessorFnB,
}: {
	id: string;
	title: string;
	isLoading: boolean;
	accessorKey: string;
	accessorFnA: (edge: T) => string;
	accessorFnB: (edge: T) => string;
	variantAccessorFnA?: (edge: T) => keyof typeof badgeVariants;
	variantAccessorFnB?: (edge: T) => keyof typeof badgeVariants;
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
				<div className="flex flex-row gap-1">
					<Badge
						variant={
							variantAccessorFnA ? variantAccessorFnA(row.original) : "default"
						}
					>
						{accessorFnA(row.original)}
					</Badge>
					<Badge
						variant={
							variantAccessorFnB ? variantAccessorFnB(row.original) : "default"
						}
					>
						{accessorFnB(row.original)}
					</Badge>
				</div>
			),
		enableSorting: false,
		enableHiding: true,
	};
}
