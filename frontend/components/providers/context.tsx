import { ReactNode } from "react";
import { AuthProvider } from "./authClient";
import { TableProvider } from "./table";

const providers = [AuthProvider, TableProvider];

export const ContextProvider = ({ children }: { children: ReactNode }) => {
	return providers.reduceRight(
		(kids, Parent) => <Parent>{kids}</Parent>,
		children,
	);
};
