"use client";
import { Button } from "@/components/ui/button";
import { useCookies } from "next-client-cookies";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";
import { useRouter } from "next/navigation";

export const redirectToLogin = async (router: AppRouterInstance) => {
	router.push("/login");
};

export function LoggedInButton() {
	const router = useRouter();
	const cookies = useCookies();

	const logout = async () => {
		cookies.remove("accessToken");
		cookies.remove("refreshToken");
		router.push("/");
		router.refresh();
	};

	return (
		<Button variant="primary" onClick={logout}>
			Logout
		</Button>
	);
}

export function LoggedOutButton() {
	const router = useRouter();
	return (
		<Button variant="primary" onClick={() => redirectToLogin(router)}>
			Login
		</Button>
	);
}
