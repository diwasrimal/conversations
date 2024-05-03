import "../styles/NavBar.css";

export default function NavBar() {
    return (
        <nav>
            <ul>
                <li>
                    <a href="/">Home</a>
                </li>
                <li>
                    <a href="/profile">Profile</a>
                </li>
                <li>
                    <a href="/search">Search</a>
                </li>
                <li>
                    <a href="/logout">Logout</a>
                </li>
            </ul>
        </nav>
    );
}
