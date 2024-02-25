import { cn } from "@/lib/utils";
import { MagnifyingGlassIcon } from "@radix-ui/react-icons";
import React from "react";

export type SearchProps = React.InputHTMLAttributes<HTMLInputElement>;

const Search = React.forwardRef<HTMLInputElement, SearchProps>(
	({ className, type, ...props }, ref) => {
		return (
			<div
				className={cn(
					"flex items-center h-8 w-full rounded-md border border-stone-200 bg-transparent px-3 py-1 text-sm shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-stone-500 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-stone-950 disabled:cursor-not-allowed disabled:opacity-50 dark:border-stone-800 dark:placeholder:text-stone-400 dark:focus-visible:ring-stone-300",
					className,
				)}
			>
				<MagnifyingGlassIcon className="h-[16px] w-[16px] mr-2" />
				<input
					{...props}
					ref={ref}
					type={type}
					className="w-full placeholder:text-muted-foreground bg-transparent focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
				/>
			</div>
		);
	},
);

Search.displayName = "Search";

export { Search };
