"use client";
import { TableHead } from "@/components/ui/table";
import { Table, flexRender } from "@tanstack/react-table";
import { useCookies } from "next-client-cookies";
import { usePathname } from "next/navigation";
import { isColumnVisible } from "./helpers";

export default function DataTableHeaders<T>({
	table,
	isLoading,
}: { table: Table<T>; isLoading: boolean }) {
	if (isLoading) {
		return [];
	}

	const cookies = useCookies();
	const pathname = usePathname();
	return table?.getHeaderGroups().map((headerGroup) => {
		return headerGroup.headers
			.filter((header) => isColumnVisible(header.column, pathname, cookies))
			.map((header) => (
				<TableHead
					key={header.column.id}
					className={
						header.column.getIsPinned()
							? "right-0 sticky h-8 flex flex-row"
							: "h-8"
					}
				>
					{flexRender(header.column.columnDef.header, header.getContext())}
				</TableHead>
			));
	});
}
