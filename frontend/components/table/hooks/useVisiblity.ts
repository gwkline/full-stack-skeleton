import { Updater } from "@tanstack/react-query";
import { VisibilityState } from "@tanstack/react-table";
import { useCookies } from "next-client-cookies";
import { useCallback, useState } from "react";
import { initVisibilityCookieStore } from "../helpers";

export function useVisibility(
	pathname: string,
	initialState = initVisibilityCookieStore(pathname, useCookies()),
) {
	const [visibility, setVisibility] = useState(initialState);
	const cookies = useCookies();

	const handleVisibilityChange = useCallback(
		(updaterOrValue: Updater<VisibilityState, VisibilityState>) => {
			let newVisibility: VisibilityState;
			if (typeof updaterOrValue === "function") {
				newVisibility = updaterOrValue(visibility);
			} else {
				newVisibility = updaterOrValue;
			}

			if (JSON.stringify(newVisibility) === JSON.stringify(visibility)) {
				return;
			}
			cookies.set(
				`columnVisibility-${pathname}`,
				JSON.stringify(newVisibility),
			);

			setVisibility(newVisibility);
		},
		[pathname, visibility, cookies.set],
	);

	return { visibility, handleVisibilityChange };
}
