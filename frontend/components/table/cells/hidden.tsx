"use client";
import { ColumnDef } from "@tanstack/react-table";

export function HiddenCell<T>({
	isLoading,
	id,
}: { isLoading: boolean; id: string }): ColumnDef<T> {
	return {
		id: id,
		enableSorting: false,
		enableColumnFilter: false,
		enableHiding: false,
		header: () => <div aria-label={id} />,
		cell: ({ row }) => {
			null;
		},
	};
}
