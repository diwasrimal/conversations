import "./BaseWithNav.css";

// Represents a base layout of the page including nav bar for a logged in user.
// Pages like home, search can extend upon this layout.
export default function BaseWithNav({ children }) {
    return (
        <div className="base">
            <NavBar />
            <main className="base-main-content">{children}</main>
        </div>
    );
}

function NavBar() {
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
