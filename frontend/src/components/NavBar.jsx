import "./NavBar.css";

export default function NavBar() {
    return (
        <aside>
            <nav>
                <ul>
                    <li>
                        <a href="/" title="Home">
                            <i className="fa-solid fa-house-chimney"></i>
                        </a>
                    </li>
                    <li>
                        <a href="/search" title="Search">
                            <i className="fa-solid fa-magnifying-glass"></i>
                        </a>
                    </li>
                    <li>
                        <a href="/logout" title="Logout">
                            <i className="fa-solid fa-right-from-bracket"></i>
                        </a>
                    </li>
                </ul>
            </nav>
        </aside>
    );
}
