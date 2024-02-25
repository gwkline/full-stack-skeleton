"use client";
import { Button } from "@/components/ui/button";
import { Search } from "@/components/ui/search";
import { Cross2Icon } from "@radix-ui/react-icons";
import { useCookies } from "next-client-cookies";
import { ReactNode, useContext } from "react";
import { AuthContext } from "../providers/authClient";
import { SearchContext, TableContext } from "../providers/table";
import { DataTableViewOptions } from "./viewOptions";

export interface ToolbarProps {
	children: ReactNode;
}

export interface FilterProps {
	children?: ReactNode;
}

export interface ActionsProps {
	children: ReactNode;
}

const Filter = ({ children }: FilterProps) => {
	const { table } = useContext(TableContext);
	const { user } = useContext(AuthContext);
	const { searchTerm, setSearchTerm } = useContext(SearchContext);
	const cookies = useCookies();

	const tableState = table?.getState();

	const isFiltered =
		searchTerm !== "" || (tableState?.columnFilters?.length || 0) > 0;

	return (
		<div className="col-span-1 flex flex-row justify-start space-x-2">
			{children}
			{isFiltered ? (
				<Button
					variant="ghost"
					onClick={() => {
						setSearchTerm("");
						table?.resetColumnFilters();
					}}
					className="h-8 px-2 lg:px-3"
				>
					Reset
					<Cross2Icon className="ml-2 h-4 w-4" />
				</Button>
			) : null}
		</div>
	);
};

const Actions = ({ children }: ActionsProps) => {
	return (
		<div className="col-span-1 flex flex-row justify-end space-x-2">
			{children}
			<DataTableViewOptions />
		</div>
	);
};

const Toolbar: React.FC<ToolbarProps> & {
	Filter: React.FC<FilterProps>;
	Actions: React.FC<ActionsProps>;
} = ({ children }) => {
	const { searchTerm, setSearchTerm } = useContext(SearchContext);

	return (
		<div className="flex mt-7 gap-4">
			<Search
				placeholder="Search..."
				className="max-w-[200px]"
				value={`${searchTerm}`}
				onChange={(event) => setSearchTerm(event.target.value)}
			/>
			<div className="flex flex-row justify-between w-full">{children}</div>
		</div>
	);
};

Toolbar.Filter = Filter;
Toolbar.Actions = Actions;

export default Toolbar;
