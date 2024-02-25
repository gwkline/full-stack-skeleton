import { useFilters } from "@/components/table/hooks/useFilters";
import { usePagination } from "@/components/table/hooks/usePagination";
import { useSelection } from "@/components/table/hooks/useSelection";
import { useSorting } from "@/components/table/hooks/useSorting";
import { useVisibility } from "@/components/table/hooks/useVisiblity";
import { Table, TableBody, TableHeader, TableRow } from "@/components/ui/table";
import { toast } from "@/components/ui/use-toast";
import { UseQueryOptions, useQuery } from "@tanstack/react-query";
import { ColumnDef } from "@tanstack/react-table";
import { getCoreRowModel, useReactTable } from "@tanstack/react-table";
import { Table as TableType } from "@tanstack/react-table";
import { graphql } from "gql.tada";
import { Cookies, useCookies } from "next-client-cookies";
import { usePathname } from "next/navigation";
import { ReactNode, useContext, useEffect, useRef } from "react";
import { useDebounce } from "use-debounce";
import { SearchContext, TableContext } from "../providers/table";
import DataTableBody from "./body";
import DataTableHeaders from "./headers";
import { DataTablePagination } from "./pagination";
import Toolbar from "./toolbar";

type Filter = {
	field: string;
	value: string[];
};

export type QueryParams = {
	after: string | undefined;
	query: string | undefined;
	filters: Filter[];
	limit: number | undefined;
	sortBy: string | undefined;
	direction: ReturnType<typeof graphql.scalar<"AscOrDesc">> | undefined;
	accessToken?: string;
};

export type ConnectionResult<T> = {
	pageInfo: {
		startCursor: string;
		endCursor: string;
		hasNextPage: boolean;
		count: number;
	};
	edges: T[];
};

interface DataTableProps<T> {
	children: ReactNode;
	columns: ColumnDef<T, unknown>[];
	queryFn: (
		params: QueryParams,
		cookies: Cookies,
	) => UseQueryOptions<
		ConnectionResult<T>,
		Error,
		ConnectionResult<T>,
		(string | QueryParams)[]
	>;
}

export const DataTable = <T,>({
	children,
	columns,
	queryFn,
}: DataTableProps<T>) => {
	const { setTable } = useContext(TableContext);
	const cookies = useCookies();
	const pathname = usePathname();
	const cursorsRef = useRef<string[]>([""]);

	const { searchTerm } = useContext(SearchContext);
	const { visibility, handleVisibilityChange } = useVisibility(pathname);
	const { sorting, handleSortChange } = useSorting([]);
	const { filters, handleFilterChange } = useFilters();
	const { pagination, handlePageChange, pageCount, setPageInfo } =
		usePagination(cursorsRef, sorting, filters);
	const { rowSelection, handleSelectionChange } = useSelection(pagination);

	const { data, isError, error, isLoading } = useQuery(
		queryFn(
			{
				sortBy: sorting[0]?.id,
				direction: sorting[0] && !sorting[0]?.desc ? "ASC" : "DESC",
				query: useDebounce(searchTerm, 600)[0],
				filters: filters.map((filter) => ({
					field: filter.id,
					value: Array.isArray(filter.value) ? filter.value : [filter.value],
				})),
				limit: pagination.pageSize,
				after: cursorsRef.current[pagination.pageIndex],
			},
			cookies,
		),
	);

	const table = useReactTable<T>({
		data: data?.edges || [],
		columns: columns,
		manualSorting: true,
		manualFiltering: true,
		manualPagination: true,
		pageCount: pageCount,
		getCoreRowModel: getCoreRowModel(),
		onSortingChange: handleSortChange,
		onPaginationChange: handlePageChange,
		onRowSelectionChange: handleSelectionChange,
		onColumnFiltersChange: handleFilterChange,
		onColumnVisibilityChange: handleVisibilityChange,
		state: {
			columnPinning: {
				right: ["actions"],
			},
			sorting,
			rowSelection,
			pagination: pagination,
			columnFilters: filters,
			columnVisibility: visibility,
		},
	});

	useEffect(() => {
		if (isError) {
			toast({
				variant: "destructive",
				title: "Query Error",
				description: error?.message.split(`: {"response":`)[0],
			});
		}
	}, [isError, error]);

	// biome-ignore lint/correctness/useExhaustiveDependencies: for some reason visibility needs this
	useEffect(() => {
		if (data) {
			setPageInfo(data.pageInfo);
			setTable({ ...(table as TableType<unknown>) });
		}
	}, [table, setTable, data, setPageInfo, visibility]);

	return (
		<>
			<Toolbar>{children}</Toolbar>
			<div className="rounded-md border mt-5 mb-3">
				<Table>
					<TableHeader className="h-8">
						<TableRow>
							<DataTableHeaders<T> table={table} isLoading={isLoading} />
						</TableRow>
					</TableHeader>
					<TableBody>
						<DataTableBody<T> table={table} isLoading={isLoading} />
					</TableBody>
				</Table>
			</div>
			<DataTablePagination<T>
				table={table}
				totalCount={data?.pageInfo.count || 0}
			/>
		</>
	);
};

DataTable.Actions = Toolbar.Actions;
DataTable.Filters = Toolbar.Filter;
