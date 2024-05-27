import { useContext, useEffect, useState } from "react";
import { Navigate } from "react-router-dom";
import { logoutSession } from "../api/functions";
import Spinner from "../components/Spinner";
import { LoginContext } from "../contexts/LoginProvider";

export default function Logout() {
	const {loginInfo, setLoginInfo} = useContext(LoginContext)
	if (!loginInfo.loggedIn) return;

	useEffect(() => {
		logoutSession().then((payload) => {
			if (payload.ok) {
				setLoginInfo({loggedIn: false, userId: undefined})
			}
		});
	}, []);
}