import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import {
	ArrowDownIcon,
	ArrowUpIcon,
	CaretSortIcon,
} from "@radix-ui/react-icons";
import { Column } from "@tanstack/react-table";

interface DataTableColumnHeaderProps<TData, TValue>
	extends React.HTMLAttributes<HTMLDivElement> {
	column: Column<TData, TValue>;
	title: string;
	isLoading: boolean;
}

export function DataTableColumnHeader<TData, TValue>({
	column,
	title,
	className,
	isLoading,
}: DataTableColumnHeaderProps<TData, TValue>) {
	return (
		<div
			className={cn("flex items-center space-x-2", className)}
			onClick={() => handleSort(column)}
			onKeyUp={() => handleSort(column)}
		>
			<Button variant="ghost" size="sm" className="-ml-3 h-8">
				<span>{isLoading ? "" : title}</span>

				{column.getCanSort() &&
					(column.getIsSorted() === "desc" ? (
						<ArrowDownIcon className="ml-2 h-4 w-4" />
					) : column.getIsSorted() === "asc" ? (
						<ArrowUpIcon className="ml-2 h-4 w-4" />
					) : (
						<CaretSortIcon className="ml-2 h-4 w-4" />
					))}
			</Button>
		</div>
	);
}

function handleSort<TData, TValue>(column: Column<TData, TValue>) {
	if (!column.getCanSort()) return;

	if (column.getIsSorted() === "asc") {
		column.toggleSorting(true); // Set to descending
	} else if (column.getIsSorted() === "desc") {
		column.clearSorting(); // Set to default
	} else {
		column.toggleSorting(false); // Set to ascending
	}
}
