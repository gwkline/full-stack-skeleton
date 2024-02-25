import { Updater } from "@tanstack/react-query";
import { PaginationState, SortingState } from "@tanstack/react-table";
import { ColumnFiltersState } from "@tanstack/react-table";
import { MutableRefObject, useCallback, useEffect, useState } from "react";

export function usePagination(
	cursorsRef: MutableRefObject<string[]>,
	sorting: SortingState,
	filters: ColumnFiltersState,
) {
	const [pagination, setPagination] = useState<PaginationState>({
		pageIndex: 0,
		pageSize: 10,
	});
	const [pageInfo, setPageInfo] = useState({
		endCursor: "",
		startCursor: "",
		count: 0,
		hasNextPage: false,
	});

	const pageCount = Math.ceil(pageInfo.count / pagination.pageSize);

	// biome-ignore lint/correctness/useExhaustiveDependencies: Reset pagination if sorting or filters changes
	useEffect(() => {
		setPagination((prev) => ({ ...prev, pageIndex: 0 }));
	}, [sorting, filters, setPagination]);

	const handlePageChange = useCallback(
		(updaterOrValue: Updater<PaginationState, PaginationState>) => {
			let newPageIndex: number;
			let newPageSize: number;
			if (typeof updaterOrValue === "function") {
				newPageIndex = updaterOrValue(pagination).pageIndex;
				newPageSize = updaterOrValue(pagination).pageSize;
			} else {
				newPageIndex = updaterOrValue.pageIndex;
				newPageSize = updaterOrValue.pageSize;
			}

			if (
				JSON.stringify(newPageIndex) === JSON.stringify(pagination.pageIndex) &&
				JSON.stringify(newPageSize) === JSON.stringify(pagination.pageSize)
			) {
				return;
			}

			// If pageSize changes, reset pageIndex and clear cursors
			if (newPageSize !== pagination.pageSize) {
				newPageIndex = 0;
				cursorsRef.current = [""];
			} else if (newPageIndex > pagination.pageIndex && pageInfo.hasNextPage) {
				// If moving to the next page, update the cursors array
				cursorsRef.current[newPageIndex] = pageInfo.endCursor || "";
			}

			setPagination({
				...pagination,
				pageIndex: newPageIndex,
				pageSize: newPageSize,
			});
		},
		[cursorsRef, pageInfo, pagination],
	);

	return {
		pagination,
		setPagination,
		handlePageChange,
		pageCount,
		setPageInfo,
		pageInfo,
	};
}
