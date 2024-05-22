import { useEffect, useState } from "react";
import { Navigate } from "react-router-dom";
import FriendshipManagerButton from "../components/FriendshipManagerButton";
import Spinner from "../components/Spinner";
import UserInfo from "../components/UserInfo";
import BaseWithNav from "../layouts/BaseWithNav";
import { getFriendRequestors, getFriends } from "../api/functions";
import "./People.css";

const loggedInUserId = Number(localStorage.getItem("loggedInUserId"));

export default function People() {
	const [loading, setLoading] = useState(true);
	const [friends, setFriends] = useState([]);
	const [friendRequestors, setFriendRequestors] = useState([]);
	const [errMsg, setErrMsg] = useState("");
	const [unauthorized, setUnauthorized] = useState(false);

	useEffect(() => {
		getFriends().then((payload) => {
			console.log("getFriends payload:", payload);
			setLoading(false);
			if (payload.ok) {
				setFriends(payload.friends);
			} else {
				setUnauthorized(payload.statusCode === 401);
				setErrMsg(payload.message);
			}
		});
		getFriendRequestors().then((payload) => {
			console.log("getFriendRequestors payload:", payload);
			setLoading(false);
			if (payload.ok) {
				setFriendRequestors(payload.friendRequestors);
			} else {
				setUnauthorized(payload.statusCode === 401);
				setErrMsg(payload.message);
			}
		});
	}, []);

	if (loading) return <Spinner />;

	if (unauthorized) return <Navigate to="/login" />;

	if (errMsg) return <p className="red-text">{errMsg}</p>;

	return (
		<BaseWithNav>
			<div className="people-page-content">
				<div className="requests-list-container">
					<h2>Requests</h2>
					{friendRequestors ? (
						<UsersList users={friendRequestors} />
					) : (
						<p>No friend requests</p>
					)}
				</div>
				<div className="friends-list-container">
					<h2>Friends</h2>
					{friends ? (
						<UsersList users={friends} />
					) : (
						<p>No friends to show</p>
					)}
				</div>
			</div>
		</BaseWithNav>
	);
}

function UsersList({ users }) {
	return (
		<ul className="users-list">
			{users.map((user) => (
				<li key={user.id}>
					<UserInfo user={user} />
					{user.id !== loggedInUserId && (
						<FriendshipManagerButton otherUserId={user.id} />
					)}
				</li>
			))}
		</ul>
	);
}
