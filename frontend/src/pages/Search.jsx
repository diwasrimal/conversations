import { useRef, useState } from "react";
import { Navigate } from "react-router-dom";
import Button from "../components/Button";
import FriendshipManagerButton from "../components/FriendshipManagerButton";
import { InputField } from "../components/InputFields";
import Spinner from "../components/Spinner";
import UserInfo from "../components/UserInfo";
import BaseWithNav from "../layouts/BaseWithNav";
import "./Search.css";
import { searchUser } from "../api/functions";

const loggedInUserId = Number(localStorage.getItem("loggedInUserId"));

export default function Search() {
    const [loading, setLoading] = useState(false);
    const [results, setResults] = useState(undefined);
    const [errMsg, setErrMsg] = useState("");
    const [unauthorized, setUnauthorized] = useState(false);
    const searchTypeRef = useRef();
    const searchQueryRef = useRef();

    function handleSearch(e) {
        e.preventDefault();
        const type = searchTypeRef.current.value;
        const query = searchQueryRef.current.value.trim();
        if (!type || !query) {
            setErrMsg("Invalid data for performing search");
            return;
        }
        setLoading(true);
        searchUser(type, query).then((payload) => {
            setLoading(false);
            if (payload.ok) {
                setResults(payload.matches || []);
                setErrMsg("");
            } else {
                setErrMsg(payload.message);
                setUnauthorized(payload.statusCode === 401);
            }
        });
    }

    if (unauthorized) return <Navigate to="/login" />;

    return (
        <BaseWithNav>
            <div className="search-page-content">
                <div className="search-area">
                    <h2>Search users</h2>
                    {errMsg && <p className="red-text">{errMsg}</p>}
                    <form className="search-form" onSubmit={handleSearch}>
                        <label htmlFor="search-type">Search By</label>
                        <select id="search-type" ref={searchTypeRef}>
                            <option value="normal">Name</option>
                            <option value="by-username">Username</option>
                        </select>
                        <InputField
                            ref={searchQueryRef}
                            placeholder="Search Query"
                            autoFocus
                        />
                        <Button>Search</Button>
                    </form>
                </div>
                {loading ? (
                    <Spinner />
                ) : (
                    results !== undefined && <SearchResults results={results} />
                )}
            </div>
        </BaseWithNav>
    );
}

function SearchResults({ results }) {
    return (
        <div className="search-results-container">
            {results.length === 0 ? (
                <p>No matches found!</p>
            ) : (
                <ul>
                    {results.map(
                        (user) =>
                            user.id !== loggedInUserId && (
                                <li key={user.id}>
                                    <SearchResult user={user} />
                                </li>
                            ),
                    )}
                </ul>
            )}
        </div>
    );
}

function SearchResult({ user }) {
    return (
        <>
            <UserInfo user={user} />
            <FriendshipManagerButton
                otherUserId={user.id}
                style={{ width: "70px" }}
            />
        </>
    );
}
