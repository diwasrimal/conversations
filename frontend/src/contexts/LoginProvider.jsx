import { useState, createContext, useEffect } from "react";
import { getLoginStatus } from "../api/functions";

const defaultValue = {
	loginInfo: {
		loggedIn: false,
		userId: undefined,
	},
	setLoginInfo: () => {},
};
export const LoginContext = createContext(defaultValue);

export  function LoginProvider({ children }) {
	const [info, setInfo] = useState(defaultValue);

	useEffect(() => {
		getLoginStatus().then((payload) => {
			if (payload.ok) {
				setInfo({
					loggedIn: true,
					userId: payload.userId,
				});
			}
		});
	}, []);

	return (
		<LoginContext.Provider
			value={{ loginInfo: info, setLoginInfo: setInfo }}
		>
			{children}
		</LoginContext.Provider>
	);
}
