"use client";

import { usePathname } from "next/navigation";

export default function Title() {
	const pathname = usePathname();
	const index = pathname.split("/").length - 1;
	return (
		<h1 className="text-xl font-bold mb-5">
			{pathname.split("/")[index].charAt(0).toUpperCase() +
				pathname.split("/")[index].slice(1)}
		</h1>
	);
}
