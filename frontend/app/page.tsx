"use client";
import { AuthContext } from "@/components/providers/authClient";
import { buttonVariants } from "@/components/ui/button";
import {
	PageActions,
	PageHeader,
	PageHeaderDescription,
	PageHeaderHeading,
} from "@/components/ui/page-header";
import { cn } from "@/lib/utils";
import { GitHubLogoIcon } from "@radix-ui/react-icons";
import Link from "next/link";
import { useContext } from "react";
const HomePage = () => {
	const { user } = useContext(AuthContext);

	return (
		<PageHeader>
			<div className="flex justify-center mt-10">
				<img
					src="/skeleton.svg"
					alt="skeleton-logo"
					style={{ height: "200px", width: "auto" }}
				/>
			</div>
			<PageHeaderHeading>Build your app better</PageHeaderHeading>
			<PageHeaderDescription>
				Full Stack Skeleton is an opinionated app starter kit designed to
				kickstart your project with the best practices in mind. It features Go
				on the backend serving a GraphQL API, integrated with PostgreSQL and
				Redis for data management, all connected to a Next.js/TypeScript/React
				frontend. Explore how you can leverage this powerful stack to build
				efficient, scalable applications.
			</PageHeaderDescription>
			<PageActions>
				<Link
					href={user ? "/apples" : "/login"}
					className={cn(buttonVariants({ variant: "primary" }))}
				>
					View Demo
				</Link>
				<Link
					target="_blank"
					rel="noreferrer"
					href={"https://github.com/gwkline/full-stack-skeleton"}
					className={cn(buttonVariants({ variant: "outline" }))}
				>
					<GitHubLogoIcon className="mr-2 h-4 w-4" />
					GitHub
				</Link>
			</PageActions>
		</PageHeader>
	);
};

export default HomePage;
