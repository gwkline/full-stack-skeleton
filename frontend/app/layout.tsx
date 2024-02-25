import Navbar from "@/components/layout/navbar";
import Sidebar from "@/components/layout/sidebar";
import { cn } from "@/lib/utils";
import { SpeedInsights } from "@vercel/speed-insights/next";
import type { Metadata } from "next";
import { CookiesProvider } from "next-client-cookies/server";
import { Inter } from "next/font/google";
import Head from "next/head";
import { cookies } from "next/headers";
import { ReactNode } from "react";
import "./globals.css";
import Providers from "./providers";
export const runtime = "edge"; // 'nodejs' (default) | 'edge'

const fontSans = Inter({
	subsets: ["latin"],
	variable: "--font-sans",
});

export const metadata: Metadata = {
	title: "Full Stack Skeleton",
};

export default function RootLayout({
	children,
	login,
}: {
	login: ReactNode;
	children: ReactNode;
}) {
	const layout = cookies().get("sidebar-layout");
	const collapsed = cookies().get("sidebar-collapsed");

	const defaultLayout = layout ? JSON.parse(layout.value) : undefined;
	const defaultCollapsed = collapsed ? JSON.parse(collapsed.value) : undefined;

	return (
		<html lang="en" suppressHydrationWarning>
			<Head>
				<meta charSet="utf-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
				<title>Full Stack Skeleton</title>
			</Head>

			<body
				className={cn(
					"min-h-screen bg-background font-sans antialiased overflow-y-hidden",
					fontSans.variable,
				)}
			>
				<SpeedInsights />
				<CookiesProvider>
					<Providers>
						<div className="h-screen flex w-screen">
							<Sidebar
								defaultLayout={defaultLayout}
								defaultCollapsed={defaultCollapsed}
							>
								<Navbar />
								{children}
							</Sidebar>
						</div>
						{login}
					</Providers>
				</CookiesProvider>
			</body>
		</html>
	);
}
