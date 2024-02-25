"use client";
import { Skeleton } from "@/components/ui/skeleton";
import { TableCell, TableRow } from "@/components/ui/table";
import { EyeNoneIcon } from "@radix-ui/react-icons";
import { Row, Table, flexRender } from "@tanstack/react-table";
import { useCookies } from "next-client-cookies";
import { usePathname } from "next/navigation";
import { MouseEvent, useState } from "react";
import { getRowRange, isColumnVisible } from "./helpers";

export default function DataTableBody<T>({
	table,
	isLoading,
}: { table: Table<T>; isLoading: boolean }) {
	if (isLoading) {
		return Array.from({ length: 10 }).map((_, index) => (
			<TableRow key={`${index * 2}a`} className="h-8">
				{Array.from({ length: 7 }).map((_, index2) => (
					<TableCell key={`${index2 * 2}a`}>
						<Skeleton className="h-8 w-full" />
					</TableCell>
				))}
			</TableRow>
		));
	}

	if (table?.getRowModel().rows.length === 0) {
		return (
			<TableRow>
				<TableCell colSpan={100} className="text-center h-[490px]">
					<div className="flex shrink-0 items-center justify-center">
						<div className="mx-auto flex max-w-[420px] flex-col items-center justify-center text-center">
							<EyeNoneIcon className="h-10 w-10 text-muted-foreground" />
							<h3 className="mt-4 text-lg font-semibold">No results found</h3>
							<p className="mb-4 mt-2 text-sm text-muted-foreground">
								If you're expecting results, try adjusting your filters.
							</p>
						</div>
					</div>
				</TableCell>
			</TableRow>
		);
	}

	const handleRowClick = (e: MouseEvent, row: Row<T>) => {
		if (e.shiftKey) {
			const { rows, rowsById } = table.getRowModel();
			const rowsToToggle = getRowRange(
				rows,
				Number(row.id),
				Number(lastSelectedID),
			);
			const isCellSelected = rowsById[row.id].getIsSelected();
			for (const _row of rowsToToggle) {
				_row.toggleSelected(!isCellSelected);
			}
		} else {
			row.toggleSelected();
		}

		setLastSelectedID(row.id);
	};

	const pathname = usePathname();
	const cookies = useCookies();
	const [lastSelectedID, setLastSelectedID] = useState("");
	return table.getRowModel().rows.map((row) => (
		<TableRow
			key={row.id}
			onClick={(e) => handleRowClick(e, row)}
			data-state={row.getIsSelected() && "selected"}
			className="h-8"
		>
			{row
				.getVisibleCells()
				.filter((cell) => isColumnVisible(cell.column, pathname, cookies))
				.map((cell) => (
					<TableCell
						key={cell.id}
						className={
							cell.column.getIsPinned()
								? "right-0 sticky justify-center flex flex-row"
								: ""
						}
					>
						{flexRender(cell.column.columnDef.cell, cell.getContext())}
					</TableCell>
				))}
		</TableRow>
	));
}
