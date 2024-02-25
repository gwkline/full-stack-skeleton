import { PaginationState, RowSelectionState } from "@tanstack/react-table";
import { useCallback, useEffect, useState } from "react";

export function useSelection(pagination: PaginationState) {
	const [rowSelection, setRowSelection] = useState<RowSelectionState>({});

	// Reset selection when pagination changes
	// biome-ignore lint/correctness/useExhaustiveDependencies: <explanation>
	useEffect(() => {
		setRowSelection({});
	}, [pagination]);

	const handleSelectionChange = useCallback(
		(
			updaterOrValue:
				| RowSelectionState
				| ((old: RowSelectionState) => RowSelectionState),
		) => {
			setRowSelection((prevSelection) => {
				if (typeof updaterOrValue === "function") {
					return updaterOrValue(prevSelection);
				}
				return updaterOrValue;
			});
		},
		[],
	);

	return { rowSelection, handleSelectionChange };
}
