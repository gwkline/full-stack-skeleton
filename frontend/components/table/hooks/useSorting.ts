import { Updater } from "@tanstack/react-query";
import { SortingState } from "@tanstack/react-table";
import { useCallback, useState } from "react";

export function useSorting(initialState: SortingState = []) {
	const [sorting, setSorting] = useState(initialState);

	const handleSortChange = useCallback(
		(updaterOrValue: Updater<SortingState, SortingState>) => {
			let newSorting: SortingState;
			if (typeof updaterOrValue === "function") {
				newSorting = updaterOrValue(sorting);
			} else {
				newSorting = updaterOrValue;
			}

			if (JSON.stringify(newSorting) === JSON.stringify(sorting)) {
				return;
			}

			setSorting(newSorting);
		},
		[sorting],
	);

	return { sorting, handleSortChange };
}
