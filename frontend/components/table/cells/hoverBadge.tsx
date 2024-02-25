import { Badge, badgeVariants } from "@/components/ui/badge";
import {
	HoverCard,
	HoverCardContent,
	HoverCardTrigger,
} from "@/components/ui/hover-card";
import { Skeleton } from "@/components/ui/skeleton";
import { ColumnDef } from "@tanstack/react-table";
import { DataTableColumnHeader } from "../columnHeader";

export function HoverBadgeCell<T>({
	id,
	title,
	isLoading,
	accessorKey,
	accessorFn,
	innerAccessorFn,
	variantAccessorFn,
}: {
	id: string;
	title: string;
	isLoading: boolean;
	accessorKey: string;
	accessorFn: (edge: T) => string;
	innerAccessorFn: (edge: T) => string;
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
				<HoverCard>
					<HoverCardTrigger>
						<Badge
							variant={
								variantAccessorFn ? variantAccessorFn(row.original) : "default"
							}
						>
							{accessorFn(row.original)}
						</Badge>
					</HoverCardTrigger>
					{innerAccessorFn(row.original) && (
						<HoverCardContent
							align="start"
							side="bottom"
							className="overflow-auto min-h-[50px] max-h-[200px]"
						>
							{innerAccessorFn(row.original)}
						</HoverCardContent>
					)}
				</HoverCard>
			),
		enableSorting: false,
		enableHiding: true,
	};
}
