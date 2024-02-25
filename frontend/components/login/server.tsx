import { cookies } from "next/headers";
import { LoggedInButton, LoggedOutButton } from "./client";

export const revalidate = 0;

export default async function LoginButton() {
	const cookieStore = cookies();
	const token = cookieStore.get("accessToken");

	return (
		<>
			{token !== undefined && token?.value !== "" ? (
				<LoggedInButton />
			) : (
				<LoggedOutButton />
			)}
		</>
	);
}
