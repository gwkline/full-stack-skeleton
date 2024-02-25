import { Updater } from "@tanstack/react-query";
import { ColumnFiltersState } from "@tanstack/react-table";
import { useCallback, useState } from "react";

export function useFilters() {
	const [filters, setFilters] = useState<ColumnFiltersState>([]);

	const handleFilterChange = useCallback(
		(updaterOrValue: Updater<ColumnFiltersState, ColumnFiltersState>) => {
			let newFilters: ColumnFiltersState;
			if (typeof updaterOrValue === "function") {
				newFilters = updaterOrValue(filters);
			} else {
				newFilters = updaterOrValue;
			}

			// Transform the filters into the expected format
			newFilters = newFilters.map((filter) => ({
				id: filter.id,
				field: filter.id,
				value: filter.value,
			}));

			if (JSON.stringify(newFilters) === JSON.stringify(filters)) {
				return;
			}

			setFilters(newFilters);
		},
		[filters],
	);

	return { filters, handleFilterChange };
}
