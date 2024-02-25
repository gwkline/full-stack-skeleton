"use client";
import { ContextProvider } from "@/components/providers/context";
import { ReactQueryProvider } from "@/components/providers/reactQuery";
import { Toaster } from "@/components/ui/toaster";
import { TooltipProvider } from "@/components/ui/tooltip";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { ThemeProvider } from "next-themes";
import { ReactNode } from "react";

export default function Providers({ children }: { children: ReactNode }) {
	return (
		<ReactQueryProvider>
			<ContextProvider>
				<TooltipProvider>
					<ThemeProvider
						attribute="class"
						defaultTheme="system"
						enableSystem
						disableTransitionOnChange
					>
						{children}
						<Toaster />
					</ThemeProvider>
				</TooltipProvider>
				<ReactQueryDevtools initialIsOpen={false} />
			</ContextProvider>
		</ReactQueryProvider>
	);
}
