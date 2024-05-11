import { useEffect, useState } from "react";

export default function Tmp() {
    const [users, setUsers] = useState([
        { id: 1, name: "ram" },
        { id: 2, name: "sita" },
    ]);
    const [selectedUser, setSelectedUser] = useState();

    return (
        <>
            {users.map((user, _) => (
                <UserCard
                    key={user.id}
                    user={user}
                    clickHandler={() => setSelectedUser(user)}
                />
            ))}
            {selectedUser && <Messages user={selectedUser} />}
        </>
    );
}

function UserCard({ user, clickHandler }) {
    return <div onClick={clickHandler}>Name: {user.name}</div>;
}

function Messages({ user }) {
    useEffect(() => {
        console.log(`mounting user ${user.name}`);
        return () => {
            console.log(`unmouting user ${user.name}`);
        };
    }, []);

    return `Messages with ${user.name}...`;
}
