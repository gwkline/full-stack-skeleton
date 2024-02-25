import Title from "@/components/layout/title";
import { ReactNode } from "react";

export default function CompsLayout({
	children,
	create,
}: {
	children: ReactNode;
	create: ReactNode;
}) {
	return (
		<>
			{create}
			<div className="container mx-auto py-7">
				<Title />
				{children}
			</div>
		</>
	);
}
