import { useEffect, useState } from "react";
import Button from "../components/Button";
import { InputField } from "../components/InputFields";

export default function Tmp() {
    return (
        <>
            <Button>Hello</Button>
            <InputField placeholder={"Enter you input"} />
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
