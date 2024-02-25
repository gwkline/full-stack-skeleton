"use client";
import {
	ResizableHandle,
	ResizablePanel,
	ResizablePanelGroup,
} from "@/components/ui/resizable";
import { Separator } from "@/components/ui/separator";
import { cn } from "@/lib/utils";
import { Apple } from "lucide-react";
import { useCookies } from "next-client-cookies";
import Image from "next/image";
import Link from "next/link";
import { ReactNode, useState } from "react";
import { Nav } from "./nav";

interface SidebarProps {
	defaultLayout: number[] | undefined;
	defaultCollapsed?: boolean;
	children: ReactNode;
}

export default function Sidebar({
	defaultLayout = [265, 1095],
	defaultCollapsed = false,
	children,
}: SidebarProps) {
	const cookies = useCookies();
	const [isCollapsed, setIsCollapsed] = useState(defaultCollapsed);
	return (
		<ResizablePanelGroup
			direction="horizontal"
			onLayout={(sizes: number[]) => {
				cookies.set("sidebar-layout", JSON.stringify(sizes));
			}}
			className="h-screen overflow-hidden"
		>
			<ResizablePanel
				defaultSize={defaultLayout[0]}
				collapsedSize={4}
				collapsible={true}
				minSize={10}
				maxSize={17}
				onCollapse={() => {
					setIsCollapsed(true);
					cookies.set("sidebar-collapsed", "true");
				}}
				onExpand={() => {
					setIsCollapsed(false);
					cookies.set("sidebar-collapsed", "false");
				}}
				className={cn(
					"flex flex-col",
					isCollapsed && "min-w-[50px] transition-all duration-300 ease-in-out",
				)}
			>
				<div className="flex-1 overflow-auto">
					<div
						className={cn(
							"flex h-[52px] items-center justify-center",
							isCollapsed ? "h-[52px]" : "px-2",
						)}
					>
						<Link href="/" prefetch>
							<Image
								src="https://upload.wikimedia.org/wikipedia/commons/3/31/Devil_Skull_Icon.svg"
								alt="Logo"
								width="0"
								height="0"
								sizes="100vw"
								className="w-6 h-auto"
								placeholder="blur"
								blurDataURL={
									"https://upload.wikimedia.org/wikipedia/commons/3/31/Devil_Skull_Icon.svg"
								}
							/>
						</Link>
					</div>
					<Separator />
					<Nav
						isCollapsed={isCollapsed}
						links={[
							{
								title: "Apples",
								link: "/apples",
								role: "user",
								icon: Apple,
							},
						]}
					/>
				</div>
			</ResizablePanel>
			<ResizableHandle withHandle />
			<ResizablePanel
				defaultSize={defaultLayout[1]}
				minSize={30}
				className="flex flex-col"
			>
				<div className="flex-1 overflow-auto">{children}</div>
			</ResizablePanel>
		</ResizablePanelGroup>
	);
}
