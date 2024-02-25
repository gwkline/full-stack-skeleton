"use client";
import { Table } from "@tanstack/react-table";
import { ReactNode, createContext, useState } from "react";
interface TableContext {
	table: Table<unknown> | null;
	setTable: (table: Table<unknown> | null) => void;
}

export const TableContext = createContext<TableContext>({
	table: null,
	setTable: () => {},
});
interface SearchTermContext {
	searchTerm: string;
	setSearchTerm: (searchTerm: string) => void;
}

export const SearchContext = createContext<SearchTermContext>({
	searchTerm: "",
	setSearchTerm: () => {},
});

interface Props {
	children: ReactNode;
}

export const TableProvider = ({ children }: Props) => {
	const [table, setTable] = useState<Table<unknown> | null>(null);
	const [searchTerm, setSearchTerm] = useState("");

	return (
		<TableContext.Provider value={{ table, setTable }}>
			<SearchContext.Provider value={{ searchTerm, setSearchTerm }}>
				{children}
			</SearchContext.Provider>
		</TableContext.Provider>
	);
};
