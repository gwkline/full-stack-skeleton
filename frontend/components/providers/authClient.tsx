"use client";
import { RefreshToken } from "@/gql/mutations/refreshToken";
import { Viewer, viewerFrag } from "@/gql/queries/viewer";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { ResultOf } from "gql.tada";
import { useCookies } from "next-client-cookies";
import { useRouter } from "next/navigation";
import { ReactNode, createContext, useEffect, useState } from "react";

interface TAuthContext {
	user: ResultOf<typeof viewerFrag> | null;
	setUser: (user: ResultOf<typeof viewerFrag> | null) => void;
}

export const AuthContext = createContext<TAuthContext>({
	user: null,
	setUser: () => {},
});

interface Props {
	children: ReactNode;
}

export const AuthProvider = ({ children }: Props) => {
	const cookies = useCookies();

	const [user, setUser] = useState<ResultOf<typeof viewerFrag> | null>(null);
	const { data, isError, error } = useQuery(Viewer(cookies));
	const queryClient = useQueryClient();
	const router = useRouter();

	const { mutate } = useMutation(RefreshToken(queryClient, router, cookies));

	useEffect(() => {
		if (!user) {
			if (data && !isError) {
				setUser(data);
			}

			if (isError) {
				if (error?.message.includes("user not found in context")) {
					const accessToken = cookies.get("accessToken");
					const refreshToken = cookies.get("refreshToken");

					if (accessToken !== undefined && refreshToken !== undefined) {
						mutate({ accessToken, refreshToken });
					}
				}
			}
		}
	}, [data, isError, error, user, mutate, cookies.get]);

	return (
		<AuthContext.Provider value={{ user, setUser }}>
			{children}
		</AuthContext.Provider>
	);
};
