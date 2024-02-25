"use client";
import { Button } from "@/components/ui/button";
import {
	DropdownMenu,
	DropdownMenuCheckboxItem,
	DropdownMenuContent,
	DropdownMenuLabel,
	DropdownMenuSeparator,
} from "@/components/ui/dropdown-menu";
import { DropdownMenuTrigger } from "@radix-ui/react-dropdown-menu";
import { MixerHorizontalIcon } from "@radix-ui/react-icons";
import { useContext } from "react";
import { TableContext } from "../providers/table";

export function DataTableViewOptions() {
	const { table } = useContext(TableContext);

	const hideableCols = table?.getAllColumns().filter((column) => {
		return (
			column.getCanHide() &&
			column.columnDef.id &&
			column.columnDef.id.trim() !== ""
		);
	});

	return (
		<DropdownMenu>
			<DropdownMenuTrigger asChild>
				<Button variant="outline" size="sm" className="h-8 flex">
					<MixerHorizontalIcon className="mr-2 h-4 w-4" />
					View
				</Button>
			</DropdownMenuTrigger>
			<DropdownMenuContent align="end" className="w-full">
				<DropdownMenuLabel className="text-center">
					Toggle columns
				</DropdownMenuLabel>
				<DropdownMenuSeparator />
				{hideableCols?.map((column) => {
					return (
						<DropdownMenuCheckboxItem
							key={column.id}
							className="capitalize"
							checked={column.getIsVisible()}
							onCheckedChange={(value) => {
								table
									?.getAllColumns()
									.find((col) => col.id === column.id)
									?.toggleVisibility(!!value);
							}}
						>
							{column.columnDef.id
								?.split("_")
								.map((word) => word.charAt(0).toUpperCase() + word.slice(1))
								.join(" ") || ""}
						</DropdownMenuCheckboxItem>
					);
				})}
			</DropdownMenuContent>
		</DropdownMenu>
	);
}
