import { useEffect, useState } from "react";
import { Navigate } from "react-router-dom";
import { logoutSession } from "../api/functions";
import Spinner from "../components/Spinner";

export default function Logout() {
	if (!localStorage.getItem("loggedInUserId"))
		return <Navigate to="/login" />;

	const [ok, setOk] = useState(false);

	useEffect(() => {
		logoutSession().then((payload) => {
			setOk(payload.ok);
		});
	}, []);

	if (ok) return <Navigate to="/login" />;

	return <Spinner />;
}
