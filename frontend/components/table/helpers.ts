import { Column, Row, VisibilityState } from "@tanstack/react-table";
import { Cookies } from "next-client-cookies";

export function isColumnVisible<T>(
	column: Column<T, unknown>,
	pageName: string,
	cookies: Cookies,
): boolean {
	const cookie =
		cookies.get(`columnVisibility-${pageName}`) ||
		cookies.get(`columnVisibility-${pageName.replaceAll("/", "%2F")}`) ||
		"{}";

	let parsedCookie = {} as VisibilityState;
	try {
		const jsonCookie = JSON.parse(cookie);
		if (isVisibilityState(jsonCookie)) {
			parsedCookie = jsonCookie;
		}
	} catch (error) {
		console.error("Error parsing visibility cookie", error);
	}

	const isVisibleFromTable = column.getIsVisible();
	const isInCookie = column.id in parsedCookie;
	const isVisibleFromCookie = !isInCookie || parsedCookie[column.id] === true;
	const isColumnDefPresent = !!column.columnDef;

	const isHeaderVisible =
		isVisibleFromTable && isVisibleFromCookie && isColumnDefPresent;

	return isHeaderVisible;
}

export function initVisibilityCookieStore(
	pageName: string,
	cookies: Cookies,
): VisibilityState {
	const savedCookie =
		cookies.get(`columnVisibility-${pageName}`) ||
		cookies.get(`columnVisibility-${pageName.replaceAll("/", "%2F")}`) ||
		"{}";

	if (savedCookie) {
		try {
			const vis = JSON.parse(savedCookie);
			if (isVisibilityState(vis)) {
				return vis;
			}
		} catch (error) {
			console.error("Error parsing visibility cookie", error);
		}
	}

	return { id: false } as VisibilityState;
}

function isVisibilityState(obj: unknown): obj is VisibilityState {
	if (typeof obj !== "object" || obj === null) {
		return false;
	}

	for (const key in obj as Record<string, unknown>) {
		if (typeof (obj as Record<string, unknown>)[key] !== "boolean") {
			return false;
		}
	}

	return true;
}

export function getRowRange<T>(
	rows: Row<T>[],
	currentID: number,
	selectedID: number,
): Row<T>[] {
	const rangeStart = selectedID > currentID ? currentID : selectedID;
	const rangeEnd = rangeStart === currentID ? selectedID : currentID;
	return rows.slice(rangeStart, rangeEnd + 1);
}
